// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: read.sql

package db

import (
	"context"
	"database/sql"
)

const getAdminByEmail = `-- name: GetAdminByEmail :one
SELECT
	id, first_name, last_name, email, superuser, created_at, updated_at
FROM
	admin
WHERE
	email = ?
`

func (q *Queries) GetAdminByEmail(ctx context.Context, email string) (Admin, error) {
	row := q.db.QueryRowContext(ctx, getAdminByEmail, email)
	var i Admin
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Superuser,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAdminLogin = `-- name: GetAdminLogin :one
SELECT
	id, first_name, last_name, email, superuser, a.created_at, a.updated_at, admin_id, password_hash, c.created_at, c.updated_at
FROM
	admin AS a
	LEFT JOIN credential AS c ON a.id = c.admin_id
WHERE
	a.email = ?
`

type GetAdminLoginRow struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	Superuser    int64
	CreatedAt    int64
	UpdatedAt    int64
	AdminID      sql.NullString
	PasswordHash []byte
	CreatedAt_2  sql.NullInt64
	UpdatedAt_2  sql.NullInt64
}

func (q *Queries) GetAdminLogin(ctx context.Context, email string) (GetAdminLoginRow, error) {
	row := q.db.QueryRowContext(ctx, getAdminLogin, email)
	var i GetAdminLoginRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Superuser,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AdminID,
		&i.PasswordHash,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
	)
	return i, err
}

const getSession = `-- name: GetSession :one
SELECT
	id, admin_id, expiry
FROM
	session
WHERE
	id = ?
`

func (q *Queries) GetSession(ctx context.Context, id string) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSession, id)
	var i Session
	err := row.Scan(&i.ID, &i.AdminID, &i.Expiry)
	return i, err
}
