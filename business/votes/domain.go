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
	Store(ctx context.Context, data *Domain, id string) error
	Update(ctx context.Context, status int, id string, refid string) error
}

type Repository interface {
	Store(ctx context.Context, data *Domain, id string) error
	Update(ctx context.Context, status int, id string, refid string) error
}
