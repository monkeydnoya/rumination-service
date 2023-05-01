package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Username string
	Email    string
}

type BlogResponse struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Author      User      `json:"author"`
}

type Blog struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Author      User      `json:"-"`
}

type DBResponse struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName       string             `json:"firstname" bson:"firstname"`
	LastName        string             `json:"lastname" bson:"lastname"`
	UserName        string             `json:"username" bson:"username"`
	Email           string             `json:"email" bson:"email"`
	Password        string             `json:"password" bson:"password"`
	PasswordConfirm string             `json:"passwordConfirm" bson:"passwordConfirm"`
	Role            []string           `json:"role"`
	Verified        bool               `json:"verified" bson:"verified"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
}

type UserValidated struct {
	ID        string    `json:"id,omitempty"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Role      []string  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
