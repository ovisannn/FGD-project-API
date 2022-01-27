package threads

import (
	"context"
	"disspace/business/comments"
	"disspace/business/user"
	"disspace/business/votes"
	"time"
)

type Domain struct {
	ID          string                 `bson:"_id,omitempty"`
	Username    string                 `bson:"username"`
	User        user.UserProfileDomain `bson:"user"`
	CategoryID  string                 `bson:"category_id"`
	Title       string                 `bson:"title"`
	Content     string                 `bson:"content"`
	ImageUrl    string                 `bson:"image_url"`
	NumVotes    int                    `bson:"num_votes,omitempty"`
	NumComments int                    `bson:"num_comments,omitempty"`
	Votes       []votes.Domain         `bson:"votes,omitempty"`
	Comments    []comments.Domain      `bson:"comments,omitempty"`
	CreatedAt   time.Time              `bson:"created_at"`
	UpdatedAt   time.Time              `bson:"updated_at"`
}

type UseCase interface {
	GetAll(ctx context.Context, sort string) ([]Domain, error)
	Create(ctx context.Context, data *Domain) (Domain, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (Domain, error)
	Update(ctx context.Context, data *Domain, id string) error
	Search(ctx context.Context, q string, sort string) ([]Domain, error)
	GetByCategoryID(ctx context.Context, categoryId string) ([]Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context, sort string) ([]Domain, error)
	Create(ctx context.Context, data *Domain) (Domain, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (Domain, error)
	Update(ctx context.Context, data *Domain, id string) error
	Search(ctx context.Context, q string, sort string) ([]Domain, error)
	GetByCategoryID(ctx context.Context, categoryId string) ([]Domain, error)
}
