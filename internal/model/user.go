package model

import (
	"database/sql"
	"time"
)

// User ...
type User struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	Info      UserInfo
}

// UserInfo ...
type UserInfo struct {
	Name  string
	Email string
}
