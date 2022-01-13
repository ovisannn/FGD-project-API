package messages

import "errors"

var (
	ErrInvalidThreadID      = errors.New("invalid thread id")
	ErrInvalidCategoriesID  = errors.New("invalid categories id")
	ErrInvalidUserID        = errors.New("invalid user id")
	ErrInvalidReferenceID   = errors.New("invalid reference id")
	ErrUnauthorizedUser     = errors.New("unautorized user")
	ErrReferenceNotFound    = errors.New("reference not found")
	ErrDataNotFound         = errors.New("data not found")
	ErrDuplicatedData       = errors.New("duplicated data")
	ErrUsernameAlreadyExist = errors.New("username already exist")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrSessionNotFound      = errors.New("session not found")
	ErrInvalidSession       = errors.New("invalid session")
	ErrAlreadyLoggedIn      = errors.New("already logged in")
)
