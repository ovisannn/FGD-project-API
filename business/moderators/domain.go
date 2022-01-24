package moderators

import (
	"context"
)

type Domain struct {
	ID         string `bson:"_id,omitempty" json:"_id,omitempty"`
	Username   string `bson:"username,omitempty" json:"username,omitempty"`
	CategoryID string `bson:"categori_id,omitempty" json:"categori_id,omitempty"`
}

type UseCase interface {
	GetByCategoryID(ctx context.Context, idCategory string) ([]Domain, error)
}

type Repository interface {
	GetByCategoryID(ctx context.Context, idCategory string) ([]Domain, error)
}
