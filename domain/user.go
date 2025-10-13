package domain

import "time"

type User struct {
	Id        int       `json:"id" db:"id"`
	UserName  string    `json:"username" validate:"required" db:"username"`
	Email     string    `json:"email" validate:"required,email" db:"email"`
	Password  string    `json:"password" validate:"required" db:"password"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"create_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}
