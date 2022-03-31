package queries

import (
	"context"
	"database/sql"
)

type DBAccessor interface {
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// IsUser returns whether the provided user exists in the database
func IsUser(ctx context.Context, db DBAccessor, username string) (bool, error) {
	var (
		count int64
		query = `SELECT COUNT(*) FROM ( SELECT DISTINCT id FROM users WHERE username = $1 ) AS check_user`
	)
	if err := db.QueryRowContext(ctx, query, username).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

// UserID returns the user ID string for the given username
func UserID(ctx context.Context, db DBAccessor, username string) (string, error) {
	var (
		userID string
		query  = `SELECT id FROM users WHERE username = $1`
	)
	if err := db.QueryRowContext(ctx, query, username).Scan(&userID); err != nil {
		return "", err
	}
	return userID, nil
}
