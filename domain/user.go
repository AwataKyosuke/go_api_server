package domain

import "time"

// User ユーザー情報
type User struct {
	ID          int
	Username    string
	Email       string
	Password    string
	DeleteFlag  bool
	CreatedAt   time.Time
	CreatedUser string
	UpdatedAt   time.Time
	UpdatedUser string
}
