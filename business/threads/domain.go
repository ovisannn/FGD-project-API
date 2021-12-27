package threads

import (
	"context"
	"time"
)

type Domain struct {
	ID        string    `bson:"_id,omitempty"`
	Title     string    `bson:"title"`
	Content   string    `bson:"content"`
	ImageUrl  string    `bson:"image_url"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	// Comments []
	// Likes []
}

type UseCase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Create(ctx context.Context, data *Domain) (Domain, error)
  GetByID(ctx context.Context, id string) (Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Create(ctx context.Context, data *Domain) (Domain, error)
  GetByID(ctx context.Context, id string) (Domain, error)
}
