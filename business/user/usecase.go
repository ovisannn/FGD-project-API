package user

import (
	"context"
	"time"
)

type UserUseCase struct {
	userRepo       Repository
	contextTimeout time.Duration
}

func NewUserUseCase(UserRepository Repository, timeout time.Duration) UseCase {
	return &UserUseCase{
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
