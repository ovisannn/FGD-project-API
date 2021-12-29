package response

import (
	"disspace/business/categories"
)

type Categories struct {
	ID           string `bson:"_id,omitempty" json:"_id"`
	CategoryName string `bson:"category_name,omitempty" json:"category_name,omitempty"`
	Description  string `bson:"description,omitempty" json:"description,omitempty"`
	Rules        []struct {
		No   int32  `bson:"no,omitempty" json:"no,omitempty"`
		Text string `bson:"text,omitempty" json:"text,omitempty"`
	} `bson:"rules,omitempty" json:"rules,omitempty"`
	Header     string `bson:"header,omitempty" json:"header,omitempty"`
	ColorTheme string `bson:"color_theme,omitempty" json:"color_theme,omitempty"`
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
