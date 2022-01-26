package comments

import (
	"context"
	"disspace/helpers/messages"
	"strings"
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
	if strings.TrimSpace(id) == "" {
		return Domain{}, messages.ErrInvalidUserID
	}

	if strings.TrimSpace(commentDomain.Text) == "" {
		return Domain{}, messages.ErrTextCannotBeEmpty
	}

	result, err := useCase.commentRepo.Create(ctx, commentDomain, id)
	if err != nil {
		return Domain{}, messages.ErrInternalServerError
	}
	return result, nil
}

func (useCase *CommentUseCase) Delete(ctx context.Context, id string, commentId string) error {
	if strings.TrimSpace(id) == "" {
		return messages.ErrInvalidUserID
	}

	if strings.TrimSpace(commentId) == "" {
		return messages.ErrInvalidCommentID
	}

	err := useCase.commentRepo.Delete(ctx, id, commentId)
	if err != nil {
		if err == messages.ErrInvalidCommentID {
			return messages.ErrInvalidCommentID
		}
		return err
	}
	return nil
}

func (useCase *CommentUseCase) Search(ctx context.Context, q string, sort string) ([]Domain, error) {
	if sort != "" && sort != "created_at" && sort != "num_votes" && sort != "num_comments" {
		return []Domain{}, messages.ErrInvalidQueryParam
	}

	result, err := useCase.commentRepo.Search(ctx, q, sort)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (useCase *CommentUseCase) GetByID(ctx context.Context, id string) (Domain, error) {
	if strings.TrimSpace(id) == "" {
		return Domain{}, messages.ErrInvalidCommentID
	}

	result, err := useCase.commentRepo.GetByID(ctx, id)
	if err != nil {
		return Domain{}, messages.ErrDataNotFound
	}
	return result, nil
}

func (useCase *CommentUseCase) GetAllInThread(ctx context.Context, threadId string, parentId string, option string) ([]Domain, error) {
	if option != "ne" && option != "" {
		return []Domain{}, messages.ErrInvalidOption
	}
	if strings.TrimSpace(threadId) == "" || strings.TrimSpace(parentId) == "" {
		return []Domain{}, messages.ErrInvalidThreadOrParent
	}

	result, err := useCase.commentRepo.GetAllInThread(ctx, threadId, parentId, option)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}
