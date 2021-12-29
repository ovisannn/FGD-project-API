package request

import (
	"disspace/business/categories"
)

type Categories struct {
	ID           string `bson:"_id,omitempty" json:"_id,omitempty"`
	CategoryName string `bson:"category_name,omitempty" json:"category_name,omitempty"`
	Description  string `bson:"description,omitempty" json:"description,omitempty"`
	Rules        []struct {
		No   int32  `bson:"no,omitempty" json:"no,omitempty"`
		Text string `bson:"text,omitempty" json:"text,omitempty"`
	} `bson:"rules,omitempty" json:"rules,omitempty"`
	Header     string `bson:"header,omitempty" json:"header,omitempty"`
	ColorTheme string `bson:"color_theme,omitempty" json:"color_theme,omitempty"`
}

func (record *Categories) ToDomain() *categories.Domain {
	return &categories.Domain{
		ID:           record.ID,
		CategoryName: record.CategoryName,
		Description:  record.Description,
		Rules:        record.Rules,
		Header:       record.Header,
		ColorTheme:   record.ColorTheme,
	}
}
