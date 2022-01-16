package reports

import (
	"context"
	"time"
)

type Domain struct {
	// ID          string    `bson:"_id,omitempty"`
	ReportedBy  string    `bson:"reported_by"`
	TargetID    string    `bson:"target_id"`
	TargetType  int       `bson:"target_type"`
	Description string    `bson:"description,omitempty"`
	UniqueID    string    `bson:"unique_id"`
	Count       int       `bson:"count,omitempty"`
	CreatedAt   time.Time `bson:"created_at"`
}

type UseCase interface {
	GetAll(ctx context.Context, sort string) ([]Domain, error)
	Create(ctx context.Context, data *Domain, id string) error
}

type Repository interface {
	GetAll(ctx context.Context, sort string) ([]Domain, error)
	Create(ctx context.Context, data *Domain, id string) error
}
