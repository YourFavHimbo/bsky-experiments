package feedgenerator

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/bluesky-social/indigo/xrpc"
	"github.com/ericvolp12/bsky-experiments/pkg/search"
	"github.com/ericvolp12/bsky-experiments/pkg/search/clusters"
	"github.com/gin-gonic/gin"
)

type FeedGenerator struct {
	PostRegistry   *search.PostRegistry
	Client         *xrpc.Client
	ClusterManager *clusters.ClusterManager
}

type FeedPostItem struct {
	Post string `json:"post"`
}

type FeedSkeleton struct {
	Feed   []FeedPostItem `json:"feed"`
	Cursor *string        `json:"cursor,omitempty"`
}

// NewFeedGenerator returns a new FeedGenerator
func NewFeedGenerator(
	ctx context.Context,
	postRegistry *search.PostRegistry,
	client *xrpc.Client,
	graphJSONUrl string,
) (*FeedGenerator, error) {

	clusterManager, err := clusters.NewClusterManager(graphJSONUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create cluster manager: %w", err)
	}

	return &FeedGenerator{
		PostRegistry:   postRegistry,
		Client:         client,
		ClusterManager: clusterManager,
	}, nil
}

func (fg *FeedGenerator) GetFeedSkeleton(c *gin.Context) {
	// Incoming requests should have a query parameter "feed" that looks like: at://did:web:feedsky.jazco.io/app.bsky.feed.generator/feed-name
	// Also a query parameter "limit" that looks like: 50
	// Also a query parameter "cursor" that contains the last post ID from the previous page of results
	feedQuery := c.Query("feed")
	if feedQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "feed query parameter is required"})
		return
	}

	// Make sure the feed query is for our feed generator service
	if !strings.HasPrefix(feedQuery, "at://did:web:feedsky.jazco.io/app.bsky.feed.generator/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "feed requested is not provided by this generator"})
		return
	}

	// Get the feed name from the query
	feedName := strings.TrimPrefix(feedQuery, "at://did:web:feedsky.jazco.io/app.bsky.feed.generator/")
	if feedName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "feed name is required"})
		return
	}

	// Get the limit from the query, default to 50, maximum of 250
	limit := int32(50)
	limitQuery := c.Query("limit")
	if limitQuery != "" {
		limit = int32(c.GetInt("limit"))
		if limit > 250 {
			limit = 250
		}
	}

	// Get the cursor from the query
	cursor := c.Query("cursor")

	posts, err := fg.PostRegistry.GetPostsPageForLabel(c.Request.Context(), feedName, limit, cursor)
	if err != nil {
		if errors.As(err, &search.NotFoundError{}) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	feedItems := make([]FeedPostItem, len(posts))
	for i, post := range posts {
		postAtURL := fmt.Sprintf("at://%s/app.bsky.feed.post/%s", post.AuthorDID, post.ID)
		feedItems[i] = FeedPostItem{Post: postAtURL}
	}

	feedSkeleton := FeedSkeleton{
		Feed: feedItems,
	}

	if len(posts) > 0 {
		feedSkeleton.Cursor = &posts[len(posts)-1].ID
	}

	c.JSON(http.StatusOK, feedSkeleton)
}