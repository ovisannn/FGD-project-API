package user

import (
	"context"
	"time"
)

type UserDomain struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName  string    `json:"full_name" bson:"full_name"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	Email     string    `json:"email" bson:"email"`
	Gender    string    `json:"gender" bson:"gender"`
	Role      int       `json:"role" bson:"role"`
	Status    string    `json:"status" bson:"status"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type UserProfileDomain struct {
	ID         string   `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId     string   `json:"user_id" bson:"user_id"`
	ProfilPict string   `json:"profile_pict" bson:"profile_pict"`
	Bio        string   `json:"bio" bson:"bio"`
	Following  []string `json:"following" bson:"following"`
	Followers  []string `json:"followers" bson:"followers"`
	Threads    []string `json:"threads" bson:"threads"`
	Reputation int      `json:"reputation" bson:"reputation"`
}

type UseCase interface {
	Register(ctx context.Context, data *UserDomain) (UserDomain, error)
}

type Repository interface {
	Register(ctx context.Context, data *UserDomain) (UserDomain, error)
}
