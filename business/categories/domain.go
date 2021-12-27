package categories

import "context"

type Domain struct {
	ID           string `bson:"_id,omitempty"`
	CategoryName string `bson:"category_name"`
	Description  string `bson:"description"`
	Rules        string `bson:"rules"`
	Header       string `bson:"header"`
	ColorTheme   string `bson:"color_theme"`
}

type RulesDomain struct {
	No   int32  `bson:"no"`
	Text string `bson:"text"`
}

type UseCase interface {
	GetAll(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
}
