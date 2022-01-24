package comments

import (
	"context"
	"disspace/helpers/messages"
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

func (useCase *CommentUseCase) Delete(ctx context.Context, id string, threadId string) error {
	err := useCase.commentRepo.Delete(ctx, id, threadId)
	if err != nil {
		if err == messages.ErrInvalidThreadID {
			return messages.ErrInvalidThreadID
		}
		return messages.ErrInvalidUserID
	}
	return nil
}

func (useCase *CommentUseCase) Search(ctx context.Context, q string, sort string) ([]Domain, error) {
	result, err := useCase.commentRepo.Search(ctx, q, sort)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (useCase *CommentUseCase) GetByID(ctx context.Context, id string) (Domain, error) {
	result, err := useCase.commentRepo.GetByID(ctx, id)
	if err != nil {
		return Domain{}, messages.ErrDataNotFound
	}
	return result, nil
}

func (useCase *CommentUseCase) GetAllInThread(ctx context.Context, threadId string, parentId string) ([]Domain, error) {
	result, err := useCase.commentRepo.GetAllInThread(ctx, threadId, parentId)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}
