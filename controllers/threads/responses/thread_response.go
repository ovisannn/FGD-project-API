package responses

import (
	"disspace/business/threads"
	"time"
)

type ThreadResponse struct {
	ID        string    `json:"_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain threads.Domain) ThreadResponse {
	return ThreadResponse{
		ID:        domain.ID,
		Title:     domain.Title,
		Content:   domain.Content,
		ImageUrl:  domain.ImageUrl,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
