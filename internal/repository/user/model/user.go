package model

import (
	"database/sql"
	"time"
)

// User ...
type User struct {
	ID        int64        `db:"id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	Info      UserInfo     `db:""`
}

// UserInfo ...
type UserInfo struct {
	Name  string `db:"name"`
	Email string `db:"email"`
}
