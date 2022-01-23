package user

import (
	"context"
	"disspace/app/middlewares"
	"disspace/helpers/encryption"
	"disspace/helpers/messages"
	"disspace/helpers/reslicing"
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

func (UseCase *UserUseCase) UserProfileGetByUsername(ctx context.Context, username string) (UserProfileDomain, error) {
	result, err := UseCase.userRepo.UserProfileGetByUsername(ctx, username)

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

	userToken, errToken := UseCase.ConfigJwt.GenerateToken(result.Username)
	if errToken != nil {
		return UserSessionDomain{}, err
	}
	newSession := UserSessionDomain{
		Token:    userToken,
		Username: result.Username,
	}

	return newSession, nil
}

func (UseCase *UserUseCase) GetUserByID(ctx context.Context, id string, usernamePayload string) (UserDomain, error) {
	getUser, err := UseCase.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return UserDomain{}, err
	}
	if getUser.Username != usernamePayload {
		return UserDomain{}, messages.ErrUnauthorizedUser
	}

	return getUser, nil
}

func (UseCase *UserUseCase) GetUserByUsername(ctx context.Context, username string, usernamePayload string) (UserDomain, error) {

	getUser, err := UseCase.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return UserDomain{}, err
	}
	if getUser.Username != usernamePayload {
		return UserDomain{}, messages.ErrUnauthorizedUser
	}

	return getUser, nil
}

func (UseCase *UserUseCase) Follow(ctx context.Context, username string, targetUsername string) error {
	//update following
	getUserProfile, err := UseCase.userRepo.UserProfileGetByUsername(ctx, username)
	if err != nil {
		return err
	}
	getTargetProfile, err := UseCase.userRepo.UserProfileGetByUsername(ctx, targetUsername)
	if err != nil {
		return err
	}

	for _, item := range getUserProfile.Following {
		if targetUsername == item {
			return messages.ErrUserAlreadyFollowed
		}
	}
	existingFollowing := getUserProfile.Following
	existingFollowing = append(existingFollowing, targetUsername)
	updateData := UserProfileDomain{
		Following: existingFollowing,
	}
	errUpdate := UseCase.userRepo.UpdateUserProfile(ctx, username, updateData)
	if errUpdate != nil {
		return err
	}

	//update followers
	existingFollowers := getTargetProfile.Followers
	existingFollowers = append(existingFollowers, username)
	updateDataTarget := UserProfileDomain{
		Followers: existingFollowers,
	}
	errUpdateTarget := UseCase.userRepo.UpdateUserProfile(ctx, targetUsername, updateDataTarget)
	if errUpdateTarget != nil {
		return err
	}

	return nil
}

func (UseCase *UserUseCase) Unfollow(ctx context.Context, username string, targetUsername string) error {
	//update following
	getUserProfile, err := UseCase.userRepo.UserProfileGetByUsername(ctx, username)
	if err != nil {
		return err
	}
	getTargetProfile, err := UseCase.userRepo.UserProfileGetByUsername(ctx, targetUsername)
	if err != nil {
		return err
	}
	deleteFollowing, errdeleteFollowing := reslicing.DeleteItemFromSlice(getUserProfile.Following, targetUsername)
	if errdeleteFollowing != nil {
		return errdeleteFollowing
	}
	updateData := UserProfileDomain{
		Following: deleteFollowing,
	}
	errUpdate := UseCase.userRepo.UpdateUserProfile(ctx, username, updateData)
	if errUpdate != nil {
		return err
	}

	//update followers
	deleteFollowers, errdeleteFollowers := reslicing.DeleteItemFromSlice(getTargetProfile.Followers, username)
	if errdeleteFollowers != nil {
		return errdeleteFollowers
	}
	updateDataTarget := UserProfileDomain{
		Followers: deleteFollowers,
	}
	errUpdateTarget := UseCase.userRepo.UpdateUserProfile(ctx, targetUsername, updateDataTarget)
	if errUpdateTarget != nil {
		return err
	}

	return nil
}

func (UseCase *UserUseCase) UpdateUserProfile(ctx context.Context, username string, data UserProfileDomain) error {
	errUpdate := UseCase.userRepo.UpdateUserProfile(ctx, username, data)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (UseCase *UserUseCase) UpdateUserInfo(ctx context.Context, username string, data UserDomain) error {
	errUpdate := UseCase.userRepo.UpdateUserInfo(ctx, username, data)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (UseCase *UserUseCase) ChangePassword(ctx context.Context, username string, data UserDomain) error {
	encryptedPass, _ := encryption.HashPassword(data.Password)
	updateData := UserDomain{
		Password: encryptedPass,
	}
	errUpdate := UseCase.userRepo.UpdateUserInfo(ctx, username, updateData)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (UseCase *UserUseCase) Logout(ctx context.Context, dataSession UserSessionDomain) error {
	getAuthorization, err := UseCase.userRepo.ConfirmAuthorization(ctx, dataSession)
	if err != nil {
		return err
	}
	if dataSession.Username != getAuthorization.Username {
		return messages.ErrInvalidSession
	}
	errLogout := UseCase.userRepo.DeleteSession(ctx, dataSession)
	if errLogout != nil {
		return err
	}
	return nil
}
