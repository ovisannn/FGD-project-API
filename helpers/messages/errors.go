package messages

import "errors"

var (
	ErrInvalidThreadID     = errors.New("invalid thread id")
	ErrInvalidCategoriesID = errors.New("invalid categories id")
)
