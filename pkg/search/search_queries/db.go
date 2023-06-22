// Code generated by sqlc. DO NOT EDIT.

package search_queries

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addAuthorStmt, err = db.PrepareContext(ctx, addAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query AddAuthor: %w", err)
	}
	if q.addAuthorBlockStmt, err = db.PrepareContext(ctx, addAuthorBlock); err != nil {
		return nil, fmt.Errorf("error preparing query AddAuthorBlock: %w", err)
	}
	if q.addAuthorToClusterStmt, err = db.PrepareContext(ctx, addAuthorToCluster); err != nil {
		return nil, fmt.Errorf("error preparing query AddAuthorToCluster: %w", err)
	}
	if q.addClusterStmt, err = db.PrepareContext(ctx, addCluster); err != nil {
		return nil, fmt.Errorf("error preparing query AddCluster: %w", err)
	}
	if q.addImageStmt, err = db.PrepareContext(ctx, addImage); err != nil {
		return nil, fmt.Errorf("error preparing query AddImage: %w", err)
	}
	if q.addLabelStmt, err = db.PrepareContext(ctx, addLabel); err != nil {
		return nil, fmt.Errorf("error preparing query AddLabel: %w", err)
	}
	if q.addLikeToPostStmt, err = db.PrepareContext(ctx, addLikeToPost); err != nil {
		return nil, fmt.Errorf("error preparing query AddLikeToPost: %w", err)
	}
	if q.addPostStmt, err = db.PrepareContext(ctx, addPost); err != nil {
		return nil, fmt.Errorf("error preparing query AddPost: %w", err)
	}
	if q.addPostLabelStmt, err = db.PrepareContext(ctx, addPostLabel); err != nil {
		return nil, fmt.Errorf("error preparing query AddPostLabel: %w", err)
	}
	if q.assignLabelToAuthorStmt, err = db.PrepareContext(ctx, assignLabelToAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query AssignLabelToAuthor: %w", err)
	}
	if q.getAllLabelsStmt, err = db.PrepareContext(ctx, getAllLabels); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllLabels: %w", err)
	}
	if q.getAllUniquePostLabelsStmt, err = db.PrepareContext(ctx, getAllUniquePostLabels); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllUniquePostLabels: %w", err)
	}
	if q.getAuthorStmt, err = db.PrepareContext(ctx, getAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query GetAuthor: %w", err)
	}
	if q.getAuthorBlockStmt, err = db.PrepareContext(ctx, getAuthorBlock); err != nil {
		return nil, fmt.Errorf("error preparing query GetAuthorBlock: %w", err)
	}
	if q.getAuthorStatsStmt, err = db.PrepareContext(ctx, getAuthorStats); err != nil {
		return nil, fmt.Errorf("error preparing query GetAuthorStats: %w", err)
	}
	if q.getAuthorsByHandleStmt, err = db.PrepareContext(ctx, getAuthorsByHandle); err != nil {
		return nil, fmt.Errorf("error preparing query GetAuthorsByHandle: %w", err)
	}
	if q.getBlockedByCountForTargetStmt, err = db.PrepareContext(ctx, getBlockedByCountForTarget); err != nil {
		return nil, fmt.Errorf("error preparing query GetBlockedByCountForTarget: %w", err)
	}
	if q.getBlocksForTargetStmt, err = db.PrepareContext(ctx, getBlocksForTarget); err != nil {
		return nil, fmt.Errorf("error preparing query GetBlocksForTarget: %w", err)
	}
	if q.getClustersStmt, err = db.PrepareContext(ctx, getClusters); err != nil {
		return nil, fmt.Errorf("error preparing query GetClusters: %w", err)
	}
	if q.getImageStmt, err = db.PrepareContext(ctx, getImage); err != nil {
		return nil, fmt.Errorf("error preparing query GetImage: %w", err)
	}
	if q.getImagesForAuthorDIDStmt, err = db.PrepareContext(ctx, getImagesForAuthorDID); err != nil {
		return nil, fmt.Errorf("error preparing query GetImagesForAuthorDID: %w", err)
	}
	if q.getImagesForPostStmt, err = db.PrepareContext(ctx, getImagesForPost); err != nil {
		return nil, fmt.Errorf("error preparing query GetImagesForPost: %w", err)
	}
	if q.getLabelByAliasStmt, err = db.PrepareContext(ctx, getLabelByAlias); err != nil {
		return nil, fmt.Errorf("error preparing query GetLabelByAlias: %w", err)
	}
	if q.getLabelsStmt, err = db.PrepareContext(ctx, getLabels); err != nil {
		return nil, fmt.Errorf("error preparing query GetLabels: %w", err)
	}
	if q.getLabelsForAuthorStmt, err = db.PrepareContext(ctx, getLabelsForAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query GetLabelsForAuthor: %w", err)
	}
	if q.getMembersOfAuthorLabelStmt, err = db.PrepareContext(ctx, getMembersOfAuthorLabel); err != nil {
		return nil, fmt.Errorf("error preparing query GetMembersOfAuthorLabel: %w", err)
	}
	if q.getMembersOfClusterStmt, err = db.PrepareContext(ctx, getMembersOfCluster); err != nil {
		return nil, fmt.Errorf("error preparing query GetMembersOfCluster: %w", err)
	}
	if q.getOldestPresentParentStmt, err = db.PrepareContext(ctx, getOldestPresentParent); err != nil {
		return nil, fmt.Errorf("error preparing query GetOldestPresentParent: %w", err)
	}
	if q.getPostStmt, err = db.PrepareContext(ctx, getPost); err != nil {
		return nil, fmt.Errorf("error preparing query GetPost: %w", err)
	}
	if q.getPostPageStmt, err = db.PrepareContext(ctx, getPostPage); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostPage: %w", err)
	}
	if q.getPostWithAuthorHandleStmt, err = db.PrepareContext(ctx, getPostWithAuthorHandle); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostWithAuthorHandle: %w", err)
	}
	if q.getPostsPageByAuthorLabelAliasStmt, err = db.PrepareContext(ctx, getPostsPageByAuthorLabelAlias); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsPageByAuthorLabelAlias: %w", err)
	}
	if q.getPostsPageByAuthorLabelAliasFromViewStmt, err = db.PrepareContext(ctx, getPostsPageByAuthorLabelAliasFromView); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsPageByAuthorLabelAliasFromView: %w", err)
	}
	if q.getPostsPageByClusterAliasStmt, err = db.PrepareContext(ctx, getPostsPageByClusterAlias); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsPageByClusterAlias: %w", err)
	}
	if q.getPostsPageByClusterAliasFromViewStmt, err = db.PrepareContext(ctx, getPostsPageByClusterAliasFromView); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsPageByClusterAliasFromView: %w", err)
	}
	if q.getPostsPageWithAnyPostLabelStmt, err = db.PrepareContext(ctx, getPostsPageWithAnyPostLabel); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsPageWithAnyPostLabel: %w", err)
	}
	if q.getPostsPageWithAnyPostLabelSortedByHotnessStmt, err = db.PrepareContext(ctx, getPostsPageWithAnyPostLabelSortedByHotness); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsPageWithAnyPostLabelSortedByHotness: %w", err)
	}
	if q.getPostsPageWithPostLabelStmt, err = db.PrepareContext(ctx, getPostsPageWithPostLabel); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsPageWithPostLabel: %w", err)
	}
	if q.getPostsPageWithPostLabelChronologicalStmt, err = db.PrepareContext(ctx, getPostsPageWithPostLabelChronological); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsPageWithPostLabelChronological: %w", err)
	}
	if q.getPostsPageWithPostLabelSortedByHotnessStmt, err = db.PrepareContext(ctx, getPostsPageWithPostLabelSortedByHotness); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsPageWithPostLabelSortedByHotness: %w", err)
	}
	if q.getThreadViewStmt, err = db.PrepareContext(ctx, getThreadView); err != nil {
		return nil, fmt.Errorf("error preparing query GetThreadView: %w", err)
	}
	if q.getTopPostersStmt, err = db.PrepareContext(ctx, getTopPosters); err != nil {
		return nil, fmt.Errorf("error preparing query GetTopPosters: %w", err)
	}
	if q.getUnprocessedImagesStmt, err = db.PrepareContext(ctx, getUnprocessedImages); err != nil {
		return nil, fmt.Errorf("error preparing query GetUnprocessedImages: %w", err)
	}
	if q.removeAuthorBlockStmt, err = db.PrepareContext(ctx, removeAuthorBlock); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveAuthorBlock: %w", err)
	}
	if q.removeLikeFromPostStmt, err = db.PrepareContext(ctx, removeLikeFromPost); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveLikeFromPost: %w", err)
	}
	if q.unassignLabelFromAuthorStmt, err = db.PrepareContext(ctx, unassignLabelFromAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query UnassignLabelFromAuthor: %w", err)
	}
	if q.updateImageStmt, err = db.PrepareContext(ctx, updateImage); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateImage: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addAuthorStmt != nil {
		if cerr := q.addAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addAuthorStmt: %w", cerr)
		}
	}
	if q.addAuthorBlockStmt != nil {
		if cerr := q.addAuthorBlockStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addAuthorBlockStmt: %w", cerr)
		}
	}
	if q.addAuthorToClusterStmt != nil {
		if cerr := q.addAuthorToClusterStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addAuthorToClusterStmt: %w", cerr)
		}
	}
	if q.addClusterStmt != nil {
		if cerr := q.addClusterStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addClusterStmt: %w", cerr)
		}
	}
	if q.addImageStmt != nil {
		if cerr := q.addImageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addImageStmt: %w", cerr)
		}
	}
	if q.addLabelStmt != nil {
		if cerr := q.addLabelStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addLabelStmt: %w", cerr)
		}
	}
	if q.addLikeToPostStmt != nil {
		if cerr := q.addLikeToPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addLikeToPostStmt: %w", cerr)
		}
	}
	if q.addPostStmt != nil {
		if cerr := q.addPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addPostStmt: %w", cerr)
		}
	}
	if q.addPostLabelStmt != nil {
		if cerr := q.addPostLabelStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addPostLabelStmt: %w", cerr)
		}
	}
	if q.assignLabelToAuthorStmt != nil {
		if cerr := q.assignLabelToAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing assignLabelToAuthorStmt: %w", cerr)
		}
	}
	if q.getAllLabelsStmt != nil {
		if cerr := q.getAllLabelsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllLabelsStmt: %w", cerr)
		}
	}
	if q.getAllUniquePostLabelsStmt != nil {
		if cerr := q.getAllUniquePostLabelsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllUniquePostLabelsStmt: %w", cerr)
		}
	}
	if q.getAuthorStmt != nil {
		if cerr := q.getAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAuthorStmt: %w", cerr)
		}
	}
	if q.getAuthorBlockStmt != nil {
		if cerr := q.getAuthorBlockStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAuthorBlockStmt: %w", cerr)
		}
	}
	if q.getAuthorStatsStmt != nil {
		if cerr := q.getAuthorStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAuthorStatsStmt: %w", cerr)
		}
	}
	if q.getAuthorsByHandleStmt != nil {
		if cerr := q.getAuthorsByHandleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAuthorsByHandleStmt: %w", cerr)
		}
	}
	if q.getBlockedByCountForTargetStmt != nil {
		if cerr := q.getBlockedByCountForTargetStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getBlockedByCountForTargetStmt: %w", cerr)
		}
	}
	if q.getBlocksForTargetStmt != nil {
		if cerr := q.getBlocksForTargetStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getBlocksForTargetStmt: %w", cerr)
		}
	}
	if q.getClustersStmt != nil {
		if cerr := q.getClustersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getClustersStmt: %w", cerr)
		}
	}
	if q.getImageStmt != nil {
		if cerr := q.getImageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getImageStmt: %w", cerr)
		}
	}
	if q.getImagesForAuthorDIDStmt != nil {
		if cerr := q.getImagesForAuthorDIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getImagesForAuthorDIDStmt: %w", cerr)
		}
	}
	if q.getImagesForPostStmt != nil {
		if cerr := q.getImagesForPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getImagesForPostStmt: %w", cerr)
		}
	}
	if q.getLabelByAliasStmt != nil {
		if cerr := q.getLabelByAliasStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLabelByAliasStmt: %w", cerr)
		}
	}
	if q.getLabelsStmt != nil {
		if cerr := q.getLabelsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLabelsStmt: %w", cerr)
		}
	}
	if q.getLabelsForAuthorStmt != nil {
		if cerr := q.getLabelsForAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLabelsForAuthorStmt: %w", cerr)
		}
	}
	if q.getMembersOfAuthorLabelStmt != nil {
		if cerr := q.getMembersOfAuthorLabelStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMembersOfAuthorLabelStmt: %w", cerr)
		}
	}
	if q.getMembersOfClusterStmt != nil {
		if cerr := q.getMembersOfClusterStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMembersOfClusterStmt: %w", cerr)
		}
	}
	if q.getOldestPresentParentStmt != nil {
		if cerr := q.getOldestPresentParentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOldestPresentParentStmt: %w", cerr)
		}
	}
	if q.getPostStmt != nil {
		if cerr := q.getPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostStmt: %w", cerr)
		}
	}
	if q.getPostPageStmt != nil {
		if cerr := q.getPostPageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostPageStmt: %w", cerr)
		}
	}
	if q.getPostWithAuthorHandleStmt != nil {
		if cerr := q.getPostWithAuthorHandleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostWithAuthorHandleStmt: %w", cerr)
		}
	}
	if q.getPostsPageByAuthorLabelAliasStmt != nil {
		if cerr := q.getPostsPageByAuthorLabelAliasStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsPageByAuthorLabelAliasStmt: %w", cerr)
		}
	}
	if q.getPostsPageByAuthorLabelAliasFromViewStmt != nil {
		if cerr := q.getPostsPageByAuthorLabelAliasFromViewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsPageByAuthorLabelAliasFromViewStmt: %w", cerr)
		}
	}
	if q.getPostsPageByClusterAliasStmt != nil {
		if cerr := q.getPostsPageByClusterAliasStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsPageByClusterAliasStmt: %w", cerr)
		}
	}
	if q.getPostsPageByClusterAliasFromViewStmt != nil {
		if cerr := q.getPostsPageByClusterAliasFromViewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsPageByClusterAliasFromViewStmt: %w", cerr)
		}
	}
	if q.getPostsPageWithAnyPostLabelStmt != nil {
		if cerr := q.getPostsPageWithAnyPostLabelStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsPageWithAnyPostLabelStmt: %w", cerr)
		}
	}
	if q.getPostsPageWithAnyPostLabelSortedByHotnessStmt != nil {
		if cerr := q.getPostsPageWithAnyPostLabelSortedByHotnessStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsPageWithAnyPostLabelSortedByHotnessStmt: %w", cerr)
		}
	}
	if q.getPostsPageWithPostLabelStmt != nil {
		if cerr := q.getPostsPageWithPostLabelStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsPageWithPostLabelStmt: %w", cerr)
		}
	}
	if q.getPostsPageWithPostLabelChronologicalStmt != nil {
		if cerr := q.getPostsPageWithPostLabelChronologicalStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsPageWithPostLabelChronologicalStmt: %w", cerr)
		}
	}
	if q.getPostsPageWithPostLabelSortedByHotnessStmt != nil {
		if cerr := q.getPostsPageWithPostLabelSortedByHotnessStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsPageWithPostLabelSortedByHotnessStmt: %w", cerr)
		}
	}
	if q.getThreadViewStmt != nil {
		if cerr := q.getThreadViewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getThreadViewStmt: %w", cerr)
		}
	}
	if q.getTopPostersStmt != nil {
		if cerr := q.getTopPostersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTopPostersStmt: %w", cerr)
		}
	}
	if q.getUnprocessedImagesStmt != nil {
		if cerr := q.getUnprocessedImagesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUnprocessedImagesStmt: %w", cerr)
		}
	}
	if q.removeAuthorBlockStmt != nil {
		if cerr := q.removeAuthorBlockStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeAuthorBlockStmt: %w", cerr)
		}
	}
	if q.removeLikeFromPostStmt != nil {
		if cerr := q.removeLikeFromPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeLikeFromPostStmt: %w", cerr)
		}
	}
	if q.unassignLabelFromAuthorStmt != nil {
		if cerr := q.unassignLabelFromAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing unassignLabelFromAuthorStmt: %w", cerr)
		}
	}
	if q.updateImageStmt != nil {
		if cerr := q.updateImageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateImageStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                              DBTX
	tx                                              *sql.Tx
	addAuthorStmt                                   *sql.Stmt
	addAuthorBlockStmt                              *sql.Stmt
	addAuthorToClusterStmt                          *sql.Stmt
	addClusterStmt                                  *sql.Stmt
	addImageStmt                                    *sql.Stmt
	addLabelStmt                                    *sql.Stmt
	addLikeToPostStmt                               *sql.Stmt
	addPostStmt                                     *sql.Stmt
	addPostLabelStmt                                *sql.Stmt
	assignLabelToAuthorStmt                         *sql.Stmt
	getAllLabelsStmt                                *sql.Stmt
	getAllUniquePostLabelsStmt                      *sql.Stmt
	getAuthorStmt                                   *sql.Stmt
	getAuthorBlockStmt                              *sql.Stmt
	getAuthorStatsStmt                              *sql.Stmt
	getAuthorsByHandleStmt                          *sql.Stmt
	getBlockedByCountForTargetStmt                  *sql.Stmt
	getBlocksForTargetStmt                          *sql.Stmt
	getClustersStmt                                 *sql.Stmt
	getImageStmt                                    *sql.Stmt
	getImagesForAuthorDIDStmt                       *sql.Stmt
	getImagesForPostStmt                            *sql.Stmt
	getLabelByAliasStmt                             *sql.Stmt
	getLabelsStmt                                   *sql.Stmt
	getLabelsForAuthorStmt                          *sql.Stmt
	getMembersOfAuthorLabelStmt                     *sql.Stmt
	getMembersOfClusterStmt                         *sql.Stmt
	getOldestPresentParentStmt                      *sql.Stmt
	getPostStmt                                     *sql.Stmt
	getPostPageStmt                                 *sql.Stmt
	getPostWithAuthorHandleStmt                     *sql.Stmt
	getPostsPageByAuthorLabelAliasStmt              *sql.Stmt
	getPostsPageByAuthorLabelAliasFromViewStmt      *sql.Stmt
	getPostsPageByClusterAliasStmt                  *sql.Stmt
	getPostsPageByClusterAliasFromViewStmt          *sql.Stmt
	getPostsPageWithAnyPostLabelStmt                *sql.Stmt
	getPostsPageWithAnyPostLabelSortedByHotnessStmt *sql.Stmt
	getPostsPageWithPostLabelStmt                   *sql.Stmt
	getPostsPageWithPostLabelChronologicalStmt      *sql.Stmt
	getPostsPageWithPostLabelSortedByHotnessStmt    *sql.Stmt
	getThreadViewStmt                               *sql.Stmt
	getTopPostersStmt                               *sql.Stmt
	getUnprocessedImagesStmt                        *sql.Stmt
	removeAuthorBlockStmt                           *sql.Stmt
	removeLikeFromPostStmt                          *sql.Stmt
	unassignLabelFromAuthorStmt                     *sql.Stmt
	updateImageStmt                                 *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                 tx,
		tx:                                 tx,
		addAuthorStmt:                      q.addAuthorStmt,
		addAuthorBlockStmt:                 q.addAuthorBlockStmt,
		addAuthorToClusterStmt:             q.addAuthorToClusterStmt,
		addClusterStmt:                     q.addClusterStmt,
		addImageStmt:                       q.addImageStmt,
		addLabelStmt:                       q.addLabelStmt,
		addLikeToPostStmt:                  q.addLikeToPostStmt,
		addPostStmt:                        q.addPostStmt,
		addPostLabelStmt:                   q.addPostLabelStmt,
		assignLabelToAuthorStmt:            q.assignLabelToAuthorStmt,
		getAllLabelsStmt:                   q.getAllLabelsStmt,
		getAllUniquePostLabelsStmt:         q.getAllUniquePostLabelsStmt,
		getAuthorStmt:                      q.getAuthorStmt,
		getAuthorBlockStmt:                 q.getAuthorBlockStmt,
		getAuthorStatsStmt:                 q.getAuthorStatsStmt,
		getAuthorsByHandleStmt:             q.getAuthorsByHandleStmt,
		getBlockedByCountForTargetStmt:     q.getBlockedByCountForTargetStmt,
		getBlocksForTargetStmt:             q.getBlocksForTargetStmt,
		getClustersStmt:                    q.getClustersStmt,
		getImageStmt:                       q.getImageStmt,
		getImagesForAuthorDIDStmt:          q.getImagesForAuthorDIDStmt,
		getImagesForPostStmt:               q.getImagesForPostStmt,
		getLabelByAliasStmt:                q.getLabelByAliasStmt,
		getLabelsStmt:                      q.getLabelsStmt,
		getLabelsForAuthorStmt:             q.getLabelsForAuthorStmt,
		getMembersOfAuthorLabelStmt:        q.getMembersOfAuthorLabelStmt,
		getMembersOfClusterStmt:            q.getMembersOfClusterStmt,
		getOldestPresentParentStmt:         q.getOldestPresentParentStmt,
		getPostStmt:                        q.getPostStmt,
		getPostPageStmt:                    q.getPostPageStmt,
		getPostWithAuthorHandleStmt:        q.getPostWithAuthorHandleStmt,
		getPostsPageByAuthorLabelAliasStmt: q.getPostsPageByAuthorLabelAliasStmt,
		getPostsPageByAuthorLabelAliasFromViewStmt:      q.getPostsPageByAuthorLabelAliasFromViewStmt,
		getPostsPageByClusterAliasStmt:                  q.getPostsPageByClusterAliasStmt,
		getPostsPageByClusterAliasFromViewStmt:          q.getPostsPageByClusterAliasFromViewStmt,
		getPostsPageWithAnyPostLabelStmt:                q.getPostsPageWithAnyPostLabelStmt,
		getPostsPageWithAnyPostLabelSortedByHotnessStmt: q.getPostsPageWithAnyPostLabelSortedByHotnessStmt,
		getPostsPageWithPostLabelStmt:                   q.getPostsPageWithPostLabelStmt,
		getPostsPageWithPostLabelChronologicalStmt:      q.getPostsPageWithPostLabelChronologicalStmt,
		getPostsPageWithPostLabelSortedByHotnessStmt:    q.getPostsPageWithPostLabelSortedByHotnessStmt,
		getThreadViewStmt:                               q.getThreadViewStmt,
		getTopPostersStmt:                               q.getTopPostersStmt,
		getUnprocessedImagesStmt:                        q.getUnprocessedImagesStmt,
		removeAuthorBlockStmt:                           q.removeAuthorBlockStmt,
		removeLikeFromPostStmt:                          q.removeLikeFromPostStmt,
		unassignLabelFromAuthorStmt:                     q.unassignLabelFromAuthorStmt,
		updateImageStmt:                                 q.updateImageStmt,
	}
}
