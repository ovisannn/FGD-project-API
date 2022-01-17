package requests

import (
	"disspace/business/user"
	"time"
)

type User struct {
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

func (request *User) UserRegisterToDomain() *user.UserDomain {
	return &user.UserDomain{
		FullName:  request.FullName,
		Username:  request.Username,
		Password:  request.Password,
		Email:     request.Email,
		Gender:    request.Gender,
		Role:      3,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (request *User) UserInfoUpdateToDomain() *user.UserDomain {
	return &user.UserDomain{
		FullName:  request.FullName,
		Username:  request.Username,
		Password:  request.Password,
		Email:     request.Email,
		Gender:    request.Gender,
		Role:      request.Role,
		Status:    request.Status,
		UpdatedAt: time.Now(),
	}
}

type UserProfile struct {
	ID          string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Username    string   `json:"username,omitempty" bson:"username,omitempty"`
	ProfilePict string   `json:"profile_pict,omitempty" bson:"profile_pict,omitempty"`
	Bio         string   `json:"bio,omitempty" bson:"bio,omitempty"`
	Following   []string `json:"following,omitempty" bson:"following,omitempty"`
	Followers   []string `json:"followers,omitempty" bson:"followers,omitempty"`
	Threads     []string `json:"threads,omitempty" bson:"threads,omitempty"`
	Reputation  int      `json:"reputation,omitempty" bson:"reputation,omitempty"`
}

func (request *UserProfile) UserProfileToDomain() *user.UserProfileDomain {
	return &user.UserProfileDomain{
		ID:          request.ID,
		Username:    request.Username,
		ProfilePict: request.ProfilePict,
		Bio:         request.Bio,
		Following:   request.Following,
		Followers:   request.Followers,
		Threads:     request.Threads,
		Reputation:  request.Reputation,
	}
}

type LoginInfo struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func (record *LoginInfo) LoginInfoToDomain() user.LoginInfoDomain {
	return user.LoginInfoDomain{
		Username: record.Username,
		Password: record.Password,
	}
}

type UserSession struct {
	Token    string `json:"token" bson:"token"`
	Username string `json:"username" bson:"username"`
}

func (record *UserSession) SessionToDomain() user.UserSessionDomain {
	return user.UserSessionDomain{
		Token:    record.Token,
		Username: record.Username,
	}
}
