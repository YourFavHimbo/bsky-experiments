// Code generated by sqlc. DO NOT EDIT.
// source: get_label_members.sql

package search_queries

import (
	"context"
)

const getMembersOfAuthorLabel = `-- name: GetMembersOfAuthorLabel :many
SELECT authors.did, authors.handle
FROM authors
JOIN author_labels ON authors.did = author_labels.author_did
JOIN labels ON author_labels.label_id = labels.id
WHERE labels.id = $1
`

func (q *Queries) GetMembersOfAuthorLabel(ctx context.Context, labelID int64) ([]Author, error) {
	rows, err := q.query(ctx, q.getMembersOfAuthorLabelStmt, getMembersOfAuthorLabel, labelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.Did, &i.Handle); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}