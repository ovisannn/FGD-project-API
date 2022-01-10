package requests

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
