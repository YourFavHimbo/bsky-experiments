package events

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	appbsky "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/ericvolp12/bsky-experiments/pkg/graph"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type Worker struct {
	WorkerID  int
	Client    *xrpc.Client
	ClientMux sync.RWMutex
}

func (bsky *BSky) worker(workerID int) {
	ctx := context.Background()

	prefix := fmt.Sprintf("[w:%d]", workerID)
	log.SetPrefix(prefix)
	log.Printf("starting worker %d\n", workerID)

	// Run a routine that refreshes the auth token every 10 minutes
	authTicker := time.NewTicker(10 * time.Minute)
	quit := make(chan struct{})
	go func() {
		log.Println("starting auth refresh routine...")
		for {
			select {
			case <-authTicker.C:
				log.Println("refreshing auth token...")
				err := bsky.RefreshAuthToken(ctx, workerID)
				if err != nil {
					log.Printf("error refreshing auth token: %s\n", err)
				} else {
					log.Println("successfully refreshed auth token")
				}
			case <-quit:
				authTicker.Stop()
				return
			}
		}
	}()

	// Pull from the work queue and process records as they come in
	for {
		record := <-bsky.RepoRecordQueue
		err := bsky.ProcessRepoRecord(
			record.ctx,
			record.pst,
			record.opPath,
			record.repoName,
			record.eventTime,
			workerID,
		)
		if err != nil {
			log.Printf("failed to process record: %v\n", err)
		}
	}
}

func (bsky *BSky) ProcessRepoRecord(
	ctx context.Context,
	pst appbsky.FeedPost,
	opPath string,
	repoName string,
	eventTime string,
	workerID int,
) error {
	prefix := fmt.Sprintf("[w:%d]", workerID)
	log.SetPrefix(prefix)
	start := time.Now()

	if pst.LexiconTypeID != "app.bsky.feed.post" {
		return nil
	}

	tracer := otel.Tracer("graph-builder")
	ctx, span := tracer.Start(ctx, "HandleRepoCommit")
	defer span.End()

	span.AddEvent("HandleRepoCommit:ResolveProfile")
	authorProfile, err := bsky.ResolveProfile(ctx, repoName, workerID)
	if err != nil {
		log.Printf("error getting profile for %s: %+v\n", repoName, err)
		return nil
	}

	span.SetAttributes(attribute.String("author_did", authorProfile.Did))
	span.SetAttributes(attribute.String("author_handle", authorProfile.Handle))

	span.AddEvent("HandleRepoCommit:DecodeFacets")
	mentions, links, err := bsky.DecodeFacets(ctx, authorProfile.Did, authorProfile.Handle, pst.Facets, workerID)
	if err != nil {
		log.Printf("error decoding post facets: %+v\n", err)
	}

	// Parse time from the event time string
	t, err := time.Parse(time.RFC3339, eventTime)
	if err != nil {
		log.Printf("error parsing time: %+v\n", err)
		return nil
	}

	postBody := strings.ReplaceAll(pst.Text, "\n", "\n\t")

	replyingTo := ""
	replyingToDID := ""
	if pst.Reply != nil && pst.Reply.Parent != nil {
		post, err := bsky.ResolvePost(ctx, pst.Reply.Parent.Uri, workerID)
		if err != nil {
			log.Printf("error resolving post (%s): %+v\n", pst.Reply.Parent.Uri, err)
			return nil
		} else if post == nil {
			log.Printf("post (%s) not found\n", pst.Reply.Parent.Uri)
			return nil
		}

		replyingTo = post.Author.Handle
		replyingToDID = post.Author.Did
		span.SetAttributes(attribute.String("replying_to", replyingTo))
		span.SetAttributes(attribute.String("replying_to_did", replyingToDID))
	}

	// Track replies in the social graph
	if replyingTo != "" && replyingToDID != "" {
		from := graph.Node{
			DID:    graph.NodeID(authorProfile.Did),
			Handle: authorProfile.Handle,
		}

		to := graph.Node{
			DID:    graph.NodeID(replyingToDID),
			Handle: replyingTo,
		}
		span.AddEvent("HandleRepoCommit:AcquireGraphLock")
		bsky.SocialGraphMux.Lock()
		bsky.SocialGraph.IncrementEdge(from, to, 1)
		span.AddEvent("HandleRepoCommit:ReleaseGraphLock")
		bsky.SocialGraphMux.Unlock()

		// Increment the reply count metric
		replyCounter.Inc()
	}

	// Grab Post ID from the Path
	pathParts := strings.Split(opPath, "/")
	postID := pathParts[len(pathParts)-1]

	postLink := fmt.Sprintf("https://staging.bsky.app/profile/%s/post/%s", authorProfile.Handle, postID)

	logMsg := fmt.Sprintf("%s", prefix)

	// Print the content of the post and any mentions or links

	// Add a Timestamp with a post link in it if we want one
	if bsky.IncludeLinks {
		logMsg += fmt.Sprintf("\u001b[90m[\x1b]8;;%s\x07%s\x1b]8;;\x07]\u001b[0m", postLink, t.Local().Format("02.01.06 15:04:05"))
	} else {
		logMsg += fmt.Sprintf("\u001b[90m%s\u001b[0m", t.Local().Format("02.01.06 15:04:05"))
	}

	// Add the user and who they are replying to if they are
	logMsg += fmt.Sprintf(" %s", authorProfile.Handle)
	if replyingTo != "" {
		logMsg += fmt.Sprintf(" \u001b[90m->\u001b[0m %s", replyingTo)
	}

	// Add the Post Body
	logMsg += fmt.Sprintf(": \n\t%s\n", postBody)

	// Add any Mentions or Links
	if len(mentions) > 0 {
		logMsg += fmt.Sprintf("\tMentions: %s\n", mentions)
	}
	if len(links) > 0 {
		logMsg += fmt.Sprintf("\tLinks: %s\n", links)
	}

	// Print the log message
	fmt.Printf("%s", logMsg)

	// Record the time to process and the count
	postsProcessedCounter.Inc()
	postProcessingDurationHistogram.Observe(time.Since(start).Seconds())

	return nil
}