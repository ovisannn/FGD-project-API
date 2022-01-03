package votes

import (
	"context"
	"time"
)

type Domain struct {
	ID          string    `bson:"_id,omitempty"`
	UserID      string    `bson:"user_id"`
	ReferenceID string    `bson:"reference_id"`
	Status      int       `bson:"status"`
	TimeLike    time.Time `bson:"time_like"`
}

type UseCase interface {
	// Create(ctx context.Context, data *Domain, id string) error
	Create(ctx context.Context, data *Domain) error
}

type Repository interface {
	// Create(ctx context.Context, data *Domain, id string) error
	Create(ctx context.Context, data *Domain) error
}
