package categories

import (
	"context"
)

type Domain struct {
	ID           string `bson:"_id,omitempty" json:"_id"`
	CategoryName string `bson:"category_name" json:"category_name"`
	Description  string `bson:"description" json:"description"`
	Rules        []struct {
		No   int32  `bson:"no" json:"no"`
		Text string `bson:"text" json:"text"`
	} `bson:"rules" json:"rules"`
	Header     string `bson:"header" json:"header"`
	ColorTheme string `bson:"color_theme" json:"color_theme"`
}

type UseCase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Create(ctx context.Context, data *Domain) (Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	// Delete(ctx context.Context, id string) error
	// Update(ctx context.Context, data *Domain, id string) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Create(ctx context.Context, data *Domain) (Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	// Delete(ctx context.Context, id string) error
	// Update(ctx context.Context, data *Domain, id string) error
}
