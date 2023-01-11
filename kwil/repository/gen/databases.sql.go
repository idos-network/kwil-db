// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: databases.sql

package gen

import (
	"context"
)

const createDatabase = `-- name: CreateDatabase :exec
INSERT INTO
    databases (db_name, db_owner)
VALUES
    ($1, (SELECT id FROM accounts WHERE account_address = $2))
`

type CreateDatabaseParams struct {
	DbName         string
	AccountAddress string
}

func (q *Queries) CreateDatabase(ctx context.Context, arg *CreateDatabaseParams) error {
	_, err := q.exec(ctx, q.createDatabaseStmt, createDatabase, arg.DbName, arg.AccountAddress)
	return err
}

const dropDatabase = `-- name: DropDatabase :exec
DELETE FROM
    databases
WHERE
    db_name = $1
    AND db_owner = (SELECT id FROM accounts WHERE account_address = $2)
`

type DropDatabaseParams struct {
	DbName         string
	AccountAddress string
}

func (q *Queries) DropDatabase(ctx context.Context, arg *DropDatabaseParams) error {
	_, err := q.exec(ctx, q.dropDatabaseStmt, dropDatabase, arg.DbName, arg.AccountAddress)
	return err
}

const getDatabaseId = `-- name: GetDatabaseId :one
SELECT
    id
FROM
    databases
WHERE
    db_name = $1
    AND db_owner = (SELECT id FROM accounts WHERE account_address = $2)
`

type GetDatabaseIdParams struct {
	DbName         string
	AccountAddress string
}

func (q *Queries) GetDatabaseId(ctx context.Context, arg *GetDatabaseIdParams) (int32, error) {
	row := q.queryRow(ctx, q.getDatabaseIdStmt, getDatabaseId, arg.DbName, arg.AccountAddress)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const listDatabases = `-- name: ListDatabases :many
SELECT
    db_name,
    account_address
FROM
    databases
    JOIN accounts ON db_owner = accounts.id
`

type ListDatabasesRow struct {
	DbName         string
	AccountAddress string
}

func (q *Queries) ListDatabases(ctx context.Context) ([]*ListDatabasesRow, error) {
	rows, err := q.query(ctx, q.listDatabasesStmt, listDatabases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*ListDatabasesRow
	for rows.Next() {
		var i ListDatabasesRow
		if err := rows.Scan(&i.DbName, &i.AccountAddress); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listDatabasesByOwner = `-- name: ListDatabasesByOwner :many
SELECT
    db_name
FROM
    databases
WHERE
    db_owner = (SELECT id FROM accounts WHERE account_address = $1)
`

func (q *Queries) ListDatabasesByOwner(ctx context.Context, accountAddress string) ([]string, error) {
	rows, err := q.query(ctx, q.listDatabasesByOwnerStmt, listDatabasesByOwner, accountAddress)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var db_name string
		if err := rows.Scan(&db_name); err != nil {
			return nil, err
		}
		items = append(items, db_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
