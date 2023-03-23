package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	User struct {
		ID        int       `db:"id"`
		Username  string    `db:"username"`
		Password  string    `db:"password"`
		Email     string    `db:"email"`
		CreatedAt time.Time `db:"created_at"`
	}
)

// UserByEmail returns Player
func UserByEmail(c context.Context, db *sqlx.DB, email string) (*User, error) {
	var u User
	err := sqlx.GetContext(c, db, &u, `
		SELECT 
			id,
			username,
			password,
			email,
			created_at
		FROM users
		WHERE email = $1`, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}
