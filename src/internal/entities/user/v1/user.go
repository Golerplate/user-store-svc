package entities_user_v1

import "time"

type User struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	IsVerified     bool      `json:"is_verified"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}