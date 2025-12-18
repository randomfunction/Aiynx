package models

import (
	"time"
)

type UserResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Dob       string    `json:"dob"` // Format: YYYY-MM-DD
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=3"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=3"` // Assuming required for update too or partial? Prompt says "UpdateUser".
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}
