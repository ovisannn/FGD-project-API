package comments

import (
	"context"
	"time"
)

type CommentUseCase struct {
	commentRepo    Repository
	contextTimeout time.Duration
}

func NewCommentUseCase(commentRepository Repository, timeout time.Duration) UseCase {
	return &CommentUseCase{
		commentRepo:    commentRepository,
		contextTimeout: timeout,
	}
}

func (useCase *CommentUseCase) Create(ctx context.Context, commentDomain *Domain, id string) (Domain, error) {
	result, err := useCase.commentRepo.Create(ctx, commentDomain, id)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}
