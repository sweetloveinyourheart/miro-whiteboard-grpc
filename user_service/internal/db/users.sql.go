// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, first_name, last_name)
VALUES ($1, $2, $3)
RETURNING user_id
`

type CreateUserParams struct {
	Email     string
	FirstName sql.NullString
	LastName  sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.FirstName, arg.LastName)
	var user_id int32
	err := row.Scan(&user_id)
	return user_id, err
}

const createUserCredential = `-- name: CreateUserCredential :exec
INSERT INTO user_credentials (user_id, password_hash)
VALUES ($1, $2)
`

type CreateUserCredentialParams struct {
	UserID       int32
	PasswordHash string
}

func (q *Queries) CreateUserCredential(ctx context.Context, arg CreateUserCredentialParams) error {
	_, err := q.db.ExecContext(ctx, createUserCredential, arg.UserID, arg.PasswordHash)
	return err
}

const getUser = `-- name: GetUser :one
SELECT user_id, email, first_name, last_name, created_at, updated_at FROM users
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userID int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT user_id, email, first_name, last_name, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserInfoWithCredentials = `-- name: GetUserInfoWithCredentials :one
SELECT users.email as email, user_credentials.password_hash as pwd FROM users
INNER JOIN user_credentials ON users.user_id = user_credentials.user_id
WHERE users.email = $1 LIMIT 1
`

type GetUserInfoWithCredentialsRow struct {
	Email string
	Pwd   string
}

func (q *Queries) GetUserInfoWithCredentials(ctx context.Context, email string) (GetUserInfoWithCredentialsRow, error) {
	row := q.db.QueryRowContext(ctx, getUserInfoWithCredentials, email)
	var i GetUserInfoWithCredentialsRow
	err := row.Scan(&i.Email, &i.Pwd)
	return i, err
}
