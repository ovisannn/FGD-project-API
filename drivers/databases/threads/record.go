package threads

import (
	"disspace/business/threads"
	"time"
)

type Thread struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string    `json:"title,omitempty" bson:"title,omitempty"`
	Content   string    `json:"content,omitempty" bson:"content,omitempty"`
	ImageUrl  string    `json:"image_url,omitempty" bson:"image_url,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (record *Thread) ToDomain() threads.Domain {
	return threads.Domain{
		ID:        record.ID,
		Title:     record.Title,
		Content:   record.Content,
		ImageUrl:  record.ImageUrl,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func FromDomain(domain threads.Domain) Thread {
	return Thread{
		ID:        domain.ID,
		Title:     domain.Title,
		Content:   domain.Content,
		ImageUrl:  domain.ImageUrl,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
