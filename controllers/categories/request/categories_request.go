package request

import (
	"disspace/business/categories"
)

type Categories struct {
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
