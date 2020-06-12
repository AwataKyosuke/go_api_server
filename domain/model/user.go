package model

import (
	"time"
)

// User ユーザー情報
type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"user_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DeleteFlag  bool      `json:"delete_flag"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedUser string    `json:"created_user"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedUser string    `json:"updated_user"`
}
