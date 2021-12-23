package threads

import (
	"context"
	"time"
)

type ThreadUseCase struct {
	threadRepo     Repository
	contextTimeout time.Duration
}

func NewThreadUseCase(threadRepository Repository, timeout time.Duration) UseCase {
	return &ThreadUseCase{
		threadRepo:     threadRepository,
		contextTimeout: timeout,
	}
}

func (useCase *ThreadUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	result, err := useCase.threadRepo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (useCase *ThreadUseCase) Create(ctx context.Context, threadDomain *Domain) (Domain, error) {
	result, err := useCase.threadRepo.Create(ctx, threadDomain)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}
