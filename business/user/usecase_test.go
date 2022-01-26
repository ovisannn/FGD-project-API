package user_test

import (
	"disspace/app/middlewares"
	"disspace/business/user"
	_userMocks "disspace/business/user/mocks"
	"time"

	// "disspace/bussiness/user"
	"testing"
)

var (
	userUseCaseMock _userMocks.UseCase
	userRepoMock    _userMocks.Repository
	userUseCase     user.UseCase

	userDomain        user.UserDomain
	userProfileDomain user.UserProfileDomain
	loginInfo         user.LoginInfoDomain
)

func TestMain(m *testing.M) {
	userUseCase = user.NewUserUseCase(&userRepoMock, time.Hour*1, middlewares.ConfigJwt{Secret: "UhYiPkGrOuP10fGd", ExpiresAt: 96})
	userDomain = user.UserDomain{
		ID:        "1",
		FullName:  "mingung",
		Username:  "joni",
		Password:  "asd",
		Email:     "joni@gmail.com",
		Gender:    "male",
		Role:      3,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	userProfileDomain = user.UserProfileDomain{
		ID:          "1",
		Username:    "mingung",
		ProfilePict: "google.com",
		Bio:         "manjadda wajadda",
		Following:   []string{"1", "2", "3"},
		Followers:   []string{"1", "2", "3"},
		Threads:     []string{"1", "2", "3"},
		Reputation:  5,
	}
	loginInfo = user.LoginInfoDomain{
		Username: "mingung",
		Password: "asd",
	}
}

func TestStore(t *testing.T) {
	t.Run("Store | Valid", func(t *testing.T) {

	})
}
