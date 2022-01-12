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

func (UseCase *UserUseCase) Login(ctx context.Context, username string, password string) (UserSessionDomain, error) {
	result, err := UseCase.userRepo.Login(ctx, username, password)
	if err != nil {
		return UserSessionDomain{}, err
	}
	// idconv, errConv := strconv.ParseUint(result.ID, 16, 64)
	// if errConv != nil {
	// 	panic(errConv)
	// }
	userToken, errToken := UseCase.ConfigJwt.GenerateToken(result.ID)
	if errToken != nil {
		return UserSessionDomain{}, err
	}
	newSession := UserSessionDomain{
		Token:    userToken,
		Username: result.Username,
	}
	errSession := UseCase.userRepo.InsertSession(ctx, newSession)
	if errSession != nil {
		return UserSessionDomain{}, err
	}
	// insert new session
	return newSession, nil
}
