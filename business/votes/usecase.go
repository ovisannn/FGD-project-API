package votes

import (
	"context"
	"disspace/helpers/messages"
	// "disspace/helpers/messages"
	"time"
)

type VoteUseCase struct {
	voteRepo       Repository
	contextTimeout time.Duration
}

func NewVoteUseCase(voteRepository Repository, timeout time.Duration) UseCase {
	return &VoteUseCase{
		voteRepo:       voteRepository,
		contextTimeout: timeout,
	}
}

func (useCase *VoteUseCase) Store(ctx context.Context, voteDomain *Domain, id string) error {
	err := useCase.voteRepo.Store(ctx, voteDomain, id)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *VoteUseCase) Update(ctx context.Context, status int, id string, refid string) error {
	err := useCase.voteRepo.Update(ctx, status, id, refid)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *VoteUseCase) GetIsVoted(ctx context.Context, username string, refId string) (Domain, error) {
	result, err := useCase.voteRepo.GetIsVoted(ctx, username, refId)
	if err != nil {
		return Domain{}, messages.ErrDataNotFound
	}
	return result, nil
}
