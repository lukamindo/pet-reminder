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

// UserByEmail returns User
func UserByEmail(c context.Context, db sqlx.ExtContext, email string) (*User, error) {
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

// UserCreate create user in database
func UserCreate(c context.Context, db sqlx.ExtContext, user User) (*int, error) {
	query, args, err := sqlx.Named(`
	INSERT INTO users
		(username
		,email
		,password)
	VALUES
		(:username
		,LOWER(:email)
		,:password)
	RETURNING id`, user)
	if err != nil {
		return nil, err
	}
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	var playerID int
	err = db.QueryRowxContext(c, query, args...).Scan(&playerID)
	if err != nil {
		return nil, err
	}
	return &playerID, nil
}
