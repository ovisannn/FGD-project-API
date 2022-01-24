package threads

import (
	"context"
	"disspace/helpers/messages"
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

func (useCase *ThreadUseCase) GetAll(ctx context.Context, sort string) ([]Domain, error) {
	result, err := useCase.threadRepo.GetAll(ctx, sort)
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

func (useCase *ThreadUseCase) Delete(ctx context.Context, id string) error {
	err := useCase.threadRepo.Delete(ctx, id)
	if err != nil {
		return messages.ErrInvalidThreadID
	}
	return nil
}

func (useCase *ThreadUseCase) GetByID(ctx context.Context, id string) (Domain, error) {
	result, err := useCase.threadRepo.GetByID(ctx, id)
	if err != nil {
		return Domain{}, messages.ErrInvalidThreadID
	}
	return result, nil
}

func (useCase *ThreadUseCase) Update(ctx context.Context, threadDomain *Domain, id string) error {
	err := useCase.threadRepo.Update(ctx, threadDomain, id)
	if err != nil {
		return messages.ErrInvalidThreadID
	}
	return nil
}

func (useCase *ThreadUseCase) Search(ctx context.Context, q string, sort string) ([]Domain, error) {
	result, err := useCase.threadRepo.Search(ctx, q, sort)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (useCase *ThreadUseCase) GetByCategoryID(ctx context.Context, categoryId string) ([]Domain, error) {
	result, err := useCase.threadRepo.GetByCategoryID(ctx, categoryId)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}
