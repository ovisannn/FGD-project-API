package user

import (
	"context"
	"time"
)

type UserDomain struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName  string    `json:"full_name,omitempty" bson:"full_name,omitempty"`
	Username  string    `json:"username,omitempty" bson:"username,omitempty"`
	Password  string    `json:"password,omitempty" bson:"password,omitempty"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty"`
	Gender    string    `json:"gender,omitempty" bson:"gender,omitempty"`
	Role      int       `json:"role,omitempty" bson:"role,omitempty"`
	Status    string    `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UserProfileDomain struct {
	ID          string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Username    string   `json:"username,omitempty" bson:"username,omitempty"`
	ProfilePict string   `json:"profile_pict,omitempty" bson:"profile_pict,omitempty"`
	Bio         string   `json:"bio,omitempty" bson:"bio,omitempty"`
	Following   []string `json:"following,omitempty" bson:"following,omitempty"`
	Followers   []string `json:"followers,omitempty" bson:"followers,omitempty"`
	Threads     []string `json:"threads,omitempty" bson:"threads,omitempty"`
	Reputation  int      `json:"reputation,omitempty" bson:"reputation,omitempty"`
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
	GetUserByUsername(ctx context.Context, username string, dataSession UserSessionDomain) (UserDomain, error)
	Login(ctx context.Context, username string, password string) (UserSessionDomain, error)
	Follow(ctx context.Context, username string, targetUsername string, dataSession UserSessionDomain) error
	Unfollow(ctx context.Context, username string, targetUsername string, dataSession UserSessionDomain) error
	UpdateUserProfile(ctx context.Context, dataSession UserSessionDomain, data UserProfileDomain) error
	UpdateUserInfo(ctx context.Context, dataSession UserSessionDomain, data UserDomain) error
}

type Repository interface {
	Register(ctx context.Context, data *UserDomain) (UserDomain, error)
	UserProfileGetByUsername(ctx context.Context, username string) (UserProfileDomain, error)
	GetUserByID(ctx context.Context, id string) (UserDomain, error)
	GetUserByUsername(ctx context.Context, username string) (UserDomain, error)
	Login(ctx context.Context, username string, password string) (UserDomain, error)
	UpdateUserProfile(ctx context.Context, username string, data UserProfileDomain) error
	UpdateUserInfo(ctx context.Context, username string, data UserDomain) error

	CheckingSession(ctx context.Context, username string) error
	InsertSession(ctx context.Context, dataSession UserSessionDomain) error
	ConfirmAuthorization(ctx context.Context, session UserSessionDomain) (UserSessionDomain, error)
}
