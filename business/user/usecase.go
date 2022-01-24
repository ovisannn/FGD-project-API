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
	loggedInCheck := UseCase.userRepo.CheckingSession(ctx, username)
	if loggedInCheck != nil {
		return UserSessionDomain{}, loggedInCheck
	}
	result, err := UseCase.userRepo.Login(ctx, username, password)
	if err != nil {
		return UserSessionDomain{}, err
	}

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

func (UseCase *UserUseCase) GetUserByID(ctx context.Context, id string, dataSession UserSessionDomain) (UserDomain, error) {
	getAuthorization, err := UseCase.userRepo.ConfirmAuthorization(ctx, dataSession)
	if err != nil {
		return UserDomain{}, err
	}
	getUser, err := UseCase.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return UserDomain{}, err
	}
	if getUser.Username != getAuthorization.Username {
		return UserDomain{}, messages.ErrInvalidSession
	}

	return getUser, nil
}

func (UseCase *UserUseCase) GetUserByUsername(ctx context.Context, username string, dataSession UserSessionDomain) (UserDomain, error) {
	getAuthorization, err := UseCase.userRepo.ConfirmAuthorization(ctx, dataSession)
	if err != nil {
		return UserDomain{}, err
	}
	getUser, err := UseCase.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return UserDomain{}, err
	}
	if getUser.Username != getAuthorization.Username {
		return UserDomain{}, messages.ErrInvalidSession
	}

	return getUser, nil
}

func (UseCase *UserUseCase) Follow(ctx context.Context, username string, targetUsername string, dataSession UserSessionDomain) error {
	getAuthorization, err := UseCase.userRepo.ConfirmAuthorization(ctx, dataSession)
	if err != nil {
		return err
	}
	getUser, err := UseCase.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}
	if getUser.Username != getAuthorization.Username {
		return messages.ErrInvalidSession
	}
	//update following
	getUserProfile, err := UseCase.userRepo.UserProfileGetByUsername(ctx, username)
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
	getTargetProfile, err := UseCase.userRepo.UserProfileGetByUsername(ctx, targetUsername)
	if err != nil {
		return err
	}
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

func (UseCase *UserUseCase) Unfollow(ctx context.Context, username string, targetUsername string, dataSession UserSessionDomain) error {
	getAuthorization, err := UseCase.userRepo.ConfirmAuthorization(ctx, dataSession)
	if err != nil {
		return err
	}
	getUser, err := UseCase.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}
	if getUser.Username != getAuthorization.Username {
		return messages.ErrInvalidSession
	}
	//update following
	getUserProfile, err := UseCase.userRepo.UserProfileGetByUsername(ctx, username)
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
	getTargetProfile, err := UseCase.userRepo.UserProfileGetByUsername(ctx, targetUsername)
	if err != nil {
		return err
	}
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

func (UseCase *UserUseCase) UpdateUserProfile(ctx context.Context, dataSession UserSessionDomain, data UserProfileDomain) error {
	getAuthorization, err := UseCase.userRepo.ConfirmAuthorization(ctx, dataSession)
	if err != nil {
		return err
	}
	if dataSession.Username != getAuthorization.Username {
		return messages.ErrInvalidSession
	}
	errUpdate := UseCase.userRepo.UpdateUserProfile(ctx, getAuthorization.Username, data)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (UseCase *UserUseCase) UpdateUserInfo(ctx context.Context, dataSession UserSessionDomain, data UserDomain) error {
	getAuthorization, err := UseCase.userRepo.ConfirmAuthorization(ctx, dataSession)
	if err != nil {
		return err
	}
	if dataSession.Username != getAuthorization.Username {
		return messages.ErrInvalidSession
	}
	errUpdate := UseCase.userRepo.UpdateUserInfo(ctx, getAuthorization.Username, data)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (UseCase *UserUseCase) ChangePassword(ctx context.Context, dataSession UserSessionDomain, data UserDomain) error {
	getAuthorization, err := UseCase.userRepo.ConfirmAuthorization(ctx, dataSession)
	if err != nil {
		return err
	}
	if dataSession.Username != getAuthorization.Username {
		return messages.ErrInvalidSession
	}
	encryptedPass, _ := encryption.HashPassword(data.Password)
	updateData := UserDomain{
		Password: encryptedPass,
	}
	errUpdate := UseCase.userRepo.UpdateUserInfo(ctx, getAuthorization.Username, updateData)
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















func (useCase *UserUseCase) Search(ctx context.Context, q string, sort string) ([]UserProfileDomain, error) {
	result, err := useCase.userRepo.Search(ctx, q, sort)
	if err != nil {
		return []UserProfileDomain{}, err
	}
	return result, nil
}
