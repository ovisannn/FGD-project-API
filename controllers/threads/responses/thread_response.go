package responses

import (
	"disspace/business/threads"
	"time"
)

type ThreadResponse struct {
	ID          string    `json:"_id"`
	Title       string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	ImageUrl    string    `json:"image_url,omitempty"`
	NumVotes    int       `json:"num_votes,omitempty"`
	NumComments int       `json:"num_comments,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func FromDomain(domain threads.Domain) ThreadResponse {
	return ThreadResponse{
		ID:          domain.ID,
		Title:       domain.Title,
		Content:     domain.Content,
		ImageUrl:    domain.ImageUrl,
		NumVotes:    domain.NumVotes,
		NumComments: domain.NumComments,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
