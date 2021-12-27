package categories

import (
	"disspace/business/categories"
)

type Categories struct {
	ID           string `bson:"_id,omitempty"`
	CategoryName string `bson:"category_name"`
	Description  string `bson:"description"`
	Rules        string `bson:"rules"`
	Header       string `bson:"header"`
	ColorTheme   string `bson:"color_theme"`
}

type Rules struct {
	No   int32  `bson:"no"`
	Text string `bson:"text"`
}

func (record *Categories) ToDomain() categories.Domain {
	return categories.Domain{
		ID:           record.ID,
		CategoryName: record.CategoryName,
		Description:  record.Description,
		Rules:        record.Rules,
		Header:       record.Header,
		ColorTheme:   record.ColorTheme,
	}
}

func FromDomain(domain categories.Domain) Categories {
	return Categories{
		ID:           domain.ID,
		CategoryName: domain.CategoryName,
		Description:  domain.Description,
		Rules:        domain.Rules,
		Header:       domain.Header,
		ColorTheme:   domain.ColorTheme,
	}
}
