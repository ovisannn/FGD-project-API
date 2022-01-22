package comments

import (
	"context"
	"time"
)

type Domain struct {
	ID          string    `bson:"_id,omitempty"`
	ThreadID    string    `bson:"thread_id"`
	UserID      string    `bson:"user_id"`
	Text        string    `bson:"text"`
	NumVotes    int       `bson:"num_votes"`
	NumComments int       `bson:"num_comments"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

type UseCase interface {
	Create(ctx context.Context, data *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string, threadId string) error
	Search(ctx context.Context, q string, sort string) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
}

type Repository interface {
	Create(ctx context.Context, data *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string, threadId string) error
	Search(ctx context.Context, q string, sort string) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
}
