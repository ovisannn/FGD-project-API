package requests

import (
	"disspace/business/threads"
	"time"
)

type Thread struct {
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	ImageUrl  string    `json:"image_url" bson:"image_url"`
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
