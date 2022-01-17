package responses

import (
	"disspace/business/threads"
	commentResp "disspace/controllers/comments/response"
	"time"
)

type ThreadResponse struct {
	ID          string                        `json:"_id"`
	UserID      string                        `json:"user_id,omitempty"`
	CategoryID  string                        `json:"category_id,omitempty"`
	Title       string                        `json:"title,omitempty"`
	Content     string                        `json:"content,omitempty"`
	ImageUrl    string                        `json:"image_url,omitempty"`
	NumVotes    int                           `json:"num_votes"`
	NumComments int                           `json:"num_comments"`
	Comments    []commentResp.CommentResponse `json:"comments,omitempty"`
	CreatedAt   time.Time                     `json:"created_at,omitempty"`
	UpdatedAt   time.Time                     `json:"updated_at,omitempty"`
}

func FromDomain(domain threads.Domain) ThreadResponse {
	var comments []commentResp.CommentResponse
	for _, getComment := range domain.Comments {
		comments = append(comments, commentResp.CommentResponse(getComment))
	}

	return ThreadResponse{
		ID:          domain.ID,
		UserID:      domain.UserID,
		CategoryID:  domain.CategoryID,
		Title:       domain.Title,
		Content:     domain.Content,
		ImageUrl:    domain.ImageUrl,
		NumVotes:    domain.NumVotes,
		NumComments: domain.NumComments,
		Comments:    comments,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
