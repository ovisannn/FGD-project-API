package user

import (
	"disspace/business/user"
	"time"
)

type User struct {
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

func SessionFromDomain(domain user.UserSessionDomain) UserSession {
	return UserSession{
		Token:    domain.Token,
		Username: domain.Username,
	}
}

func (record *User) UserToDomain() user.UserDomain {
	return user.UserDomain{
		ID:        record.ID,
		FullName:  record.FullName,
		Username:  record.Username,
		Password:  record.Password,
		Email:     record.Email,
		Gender:    record.Gender,
		Role:      record.Role,
		Status:    record.Status,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func (record *UserProfile) UserProfileToDomain() user.UserProfileDomain {
	return user.UserProfileDomain{
		ID:          record.ID,
		Username:    record.Username,
		ProfilePict: record.ProfilePict,
		Bio:         record.Bio,
		Following:   record.Following,
		Followers:   record.Followers,
		Threads:     record.Threads,
		Reputation:  record.Reputation,
	}
}

func UserFromDomain(domain user.UserDomain) User {
	return User{
		ID:        domain.ID,
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
