// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: indexes.sql

package repository

import (
	"context"

	"github.com/lib/pq"
)

const createIndex = `-- name: CreateIndex :exec
INSERT INTO
    INDEXES (table_id, index_name, index_type, COLUMNS)
VALUES
    (
        (
            SELECT
                id
            FROM
                tables
            WHERE
                table_name = $1
        ),
        $2,
        $3,
        $4
    )
`

type CreateIndexParams struct {
	TableName string
	IndexName string
	IndexType int32
	Columns   []string
}

func (q *Queries) CreateIndex(ctx context.Context, arg *CreateIndexParams) error {
	_, err := q.exec(ctx, q.createIndexStmt, createIndex,
		arg.TableName,
		arg.IndexName,
		arg.IndexType,
		pq.Array(arg.Columns),
	)
	return err
}
