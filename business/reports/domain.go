package reports

import (
	"context"
	"time"
)

type Domain struct {
	ID          string    `bson:"_id,omitempty"`
	ReportedBy  string    `bson:"reported_by"`
	TargetID    string    `bson:"target_id"`
	TargetType  int       `bson:"target_type"`
	Description string    `bson:"description,omitempty"`
	UniqueID    string    `bson:"unique_id"`
	CreatedAt   time.Time `bson:"created_at"`
}

type UseCase interface {
	Create(ctx context.Context, data *Domain, id string) error
}

type Repository interface {
	Create(ctx context.Context, data *Domain, id string) error
}
