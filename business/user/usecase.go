package user

import (
	"context"
	"disspace/app/middlewares"
	"time"
)

type UserUseCase struct {
	ConfigJwt      middlewares.ConfigJwt
	userRepo       Repository
	contextTimeout time.Duration
}

func NewUserUseCase(UserRepository Repository, timeout time.Duration, ConfigJWT middlewares.ConfigJwt) UseCase {
	return &UserUseCase{
		ConfigJwt:      ConfigJWT,
		userRepo:       UserRepository,
		contextTimeout: timeout,
	}
}

func (useCase *UserUseCase) Register(ctx context.Context, data *UserDomain) (UserDomain, error) {
	result, err := useCase.userRepo.Register(ctx, data)
	if err != nil {
		return UserDomain{}, err
	}
	return result, nil
}

func (UseCase *UserUseCase) UserProfileGetByUserID(ctx context.Context, id string) (UserProfileDomain, error) {
	result, err := UseCase.userRepo.UserProfileGetByUserID(ctx, id)

	if err != nil {
		return UserProfileDomain{}, err
	}
	return result, nil
}

func (UseCase *UserUseCase) Login(ctx context.Context, username string, password string) (UserDomain, error) {
	return UserDomain{}, nil
}
