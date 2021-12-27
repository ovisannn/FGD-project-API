package requests

import (
	"disspace/business/threads"
	"time"
)

type Thread struct {
	Title     string    `json:"title,omitempty" bson:"title,omitempty"`
	Content   string    `json:"content,omitempty" bson:"content,omitempty"`
	ImageUrl  string    `json:"image_url,omitempty" bson:"image_url,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (request *Thread) ToDomain() *threads.Domain {
	return &threads.Domain{
		Title:     request.Title,
		Content:   request.Content,
		ImageUrl:  request.ImageUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
