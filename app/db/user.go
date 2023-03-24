package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lukamindo/pet-reminder/app/request"
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

// UserByEmail returns User
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

// CreateUser inserts User
func CreateUser(c context.Context, db *sqlx.DB, urr request.RegistationParams) error {
	_, err := sqlx.NamedExecContext(c, db, `
		INSERT INTO users
			(username
			,email
			,password)
		VALUES
			(:username
			,:email
			,:password)`, urr)
	return err
}
