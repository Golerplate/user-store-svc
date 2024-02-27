package entities_user_v1

import "time"

type User struct {
	ID         string    `json:"id"`
	ExternalID string    `json:"external_id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ServiceCreateUserRequest struct {
	ExternalID string `json:"external_id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
}

type GRPCCreateUserRequest struct {
	ExternalID string `json:"external_id"`
	Email      string `json:"email"`
}
