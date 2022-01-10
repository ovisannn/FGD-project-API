package responses

import (
	"disspace/business/user"
	"time"
)

type User struct {
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
