package messages

import "errors"

var (
	ErrInvalidThreadID       = errors.New("invalid thread id")
	ErrInvalidCategoriesID   = errors.New("invalid categories id")
	ErrInvalidUserID         = errors.New("invalid user id")
	ErrInvalidReferenceID    = errors.New("invalid reference id")
	ErrUnauthorizedUser      = errors.New("unauthorized user")
	ErrReferenceNotFound     = errors.New("reference not found")
	ErrDataNotFound          = errors.New("data not found")
	ErrDuplicatedData        = errors.New("duplicated data")
	ErrUsernameAlreadyExist  = errors.New("username already exist")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrSessionNotFound       = errors.New("session not found")
	ErrInvalidSession        = errors.New("invalid session")
	ErrAlreadyLoggedIn       = errors.New("already logged in")
	ErrUpdateFailed          = errors.New("update failed")
	ErrUserAlreadyFollowed   = errors.New("user already followed")
	ErrFailedClaimJWT        = errors.New("failed claiming jwt payload")
	ErrInvalidCommentID      = errors.New("invalid comment id")
	ErrInvalidQueryParam     = errors.New("invalid query param")
	ErrTextCannotBeEmpty     = errors.New("text body cannot be empty")
	ErrInvalidThreadOrParent = errors.New("invalid thread or param id")
	ErrEmptyTitle            = errors.New("title cannot be empty")
	ErrInvalidOption         = errors.New("invalid option")
)
