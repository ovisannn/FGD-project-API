package response

import (
	"disspace/business/comments"
	"time"
)

type CommentResponse struct {
	ID          string    `json:"_id"`
	ThreadID    string    `json:"thread_id,omitempty"`
	UserID      string    `json:"user_id,omitempty"`
	Text        string    `json:"text,omitempty"`
	NumVotes    int       `json:"num_votes"`
	NumComments int       `json:"num_comments,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func FromDomain(domain comments.Domain) CommentResponse {
	return CommentResponse{
		ID:          domain.ID,
		ThreadID:    domain.ThreadID,
		UserID:      domain.UserID,
		Text:        domain.Text,
		NumVotes:    domain.NumVotes,
		NumComments: domain.NumComments,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
