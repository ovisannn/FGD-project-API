package comments

import (
	"context"
	"time"
)

type Domain struct {
	ID        string    `bson:"_id,omitempty"`
	ThreadID  string    `bson:"thread_id"`
	UserID    string    `bson:"user_id"`
	Text      string    `bson:"text"`
	NumVotes  int       `bson:"num_votes"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type UseCase interface {
	Create(ctx context.Context, data *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string, threadId string) error
}

type Repository interface {
	Create(ctx context.Context, data *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string, threadId string) error
}
