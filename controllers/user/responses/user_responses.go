package responses

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

func UserFromDomain(domain user.UserDomain) User {
	return User{
		FullName:  domain.FullName,
		Username:  domain.Username,
		Password:  domain.Password,
		Email:     domain.Email,
		Gender:    domain.Gender,
		Role:      domain.Role,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type UserProfile struct {
	ID          string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Username    string   `json:"username" bson:"username"`
	ProfilePict string   `json:"profile_pict" bson:"profile_pict"`
	Bio         string   `json:"bio" bson:"bio"`
	Following   []string `json:"following" bson:"following"`
	Followers   []string `json:"followers" bson:"followers"`
	Threads     []string `json:"threads" bson:"threads"`
	Reputation  int      `json:"reputation" bson:"reputation"`
}

func UserProfileFromDomain(domain user.UserProfileDomain) UserProfile {
	return UserProfile{
		ID:          domain.ID,
		Username:    domain.Username,
		ProfilePict: domain.ProfilePict,
		Bio:         domain.Bio,
		Following:   domain.Following,
		Followers:   domain.Followers,
		Threads:     domain.Threads,
		Reputation:  domain.Reputation,
	}
}

type UserSession struct {
	Token    string `json:"token" bson:"token"`
	Username string `json:"username" bson:"username"`
}

func SessionFromDomain(domain user.UserSessionDomain) UserSession {
	return UserSession{
		Token:    domain.Token,
		Username: domain.Username,
	}
}
