package votes

import (
	"context"
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

