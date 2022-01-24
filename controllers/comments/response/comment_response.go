package response

import (
	"disspace/business/comments"
	userResp "disspace/controllers/user/responses"
	"time"
)

type CommentResponse struct {
	ID          string               `json:"_id"`
	ThreadID    string               `json:"thread_id,omitempty"`
	ParentID    string               `json:"parent_id,omitempty"`
	Username    string               `json:"username,omitempty"`
	User        userResp.UserProfile `json:"user,omitempty"`
	Text        string               `json:"text,omitempty"`
	NumVotes    int                  `json:"num_votes,omitempty"`
	NumComments int                  `json:"num_comments,omitempty"`
	CreatedAt   time.Time            `json:"created_at,omitempty"`
	UpdatedAt   time.Time            `json:"updated_at,omitempty"`
}

func FromDomain(domain comments.Domain) CommentResponse {
	return CommentResponse{
		ID:          domain.ID,
		ThreadID:    domain.ThreadID,
		ParentID:    domain.ParentID,
		Username:    domain.Username,
		User:        userResp.UserProfileFromDomain(domain.User),
		Text:        domain.Text,
		NumVotes:    domain.NumVotes,
		NumComments: domain.NumComments,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
