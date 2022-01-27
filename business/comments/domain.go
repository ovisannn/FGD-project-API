package comments

import (
	"context"
	"disspace/business/user"
	"disspace/business/votes"
	"time"
)

type Domain struct {
	ID          string                 `bson:"_id,omitempty"`
	ThreadID    string                 `bson:"thread_id"`
	ParentID    string                 `bson:"parent_id"`
	Username    string                 `bson:"username"`
	User        user.UserProfileDomain `bson:"user"`
	Text        string                 `bson:"text"`
	NumVotes    int                    `bson:"num_votes"`
	NumComments int                    `bson:"num_comments"`
	Votes       []votes.Domain         `bson:"votes,omitempty"`
	CreatedAt   time.Time              `bson:"created_at"`
	UpdatedAt   time.Time              `bson:"updated_at"`
}

type UseCase interface {
	Create(ctx context.Context, data *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string, commentId string) error
	Search(ctx context.Context, q string, sort string) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	GetAllInThread(ctx context.Context, threadId string, parentId string, option string) ([]Domain, error)
}

type Repository interface {
	Create(ctx context.Context, data *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string, commentId string) error
	Search(ctx context.Context, q string, sort string) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	GetAllInThread(ctx context.Context, threadId string, parentId string, option string) ([]Domain, error)
}
