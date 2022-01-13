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
	ID          string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Username    string   `json:"username" bson:"username"`
	ProfilePict string   `json:"profile_pict" bson:"profile_pict"`
	Bio         string   `json:"bio" bson:"bio"`
	Following   []string `json:"following" bson:"following"`
	Followers   []string `json:"followers" bson:"followers"`
	Threads     []string `json:"threads" bson:"threads"`
	Reputation  int      `json:"reputation" bson:"reputation"`
}

type UserSessionDomain struct {
	Token    string `json:"token" bson:"token"`
	Username string `json:"username" bson:"username"`
}

type LoginInfoDomain struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type UseCase interface {
	Register(ctx context.Context, data *UserDomain) (UserDomain, error)
	UserProfileGetByUsername(ctx context.Context, username string) (UserProfileDomain, error)
	GetUserByID(ctx context.Context, id string, dataSession UserSessionDomain) (UserDomain, error)
	Login(ctx context.Context, username string, password string) (UserSessionDomain, error) //already loged in feature
}

type Repository interface {
	Register(ctx context.Context, data *UserDomain) (UserDomain, error)
	UserProfileGetByUsername(ctx context.Context, username string) (UserProfileDomain, error)
	GetUserByID(ctx context.Context, id string) (UserDomain, error)
	Login(ctx context.Context, username string, password string) (UserDomain, error)
	InsertSession(ctx context.Context, dataSession UserSessionDomain) error
	ConfirmAuthorization(ctx context.Context, session UserSessionDomain) (UserSessionDomain, error)
}
