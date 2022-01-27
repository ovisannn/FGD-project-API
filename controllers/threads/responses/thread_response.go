package responses

import (
	"disspace/business/threads"
	commentResp "disspace/controllers/comments/response"
	userResp "disspace/controllers/user/responses"
	votesResp "disspace/controllers/votes/response"
	"time"
)

type ThreadResponse struct {
	ID          string                        `json:"_id"`
	Username    string                        `json:"username,omitempty"`
	User        userResp.UserProfile          `json:"user,omitempty"`
	CategoryID  string                        `json:"category_id,omitempty"`
	Title       string                        `json:"title,omitempty"`
	Content     string                        `json:"content,omitempty"`
	ImageUrl    string                        `json:"image_url,omitempty"`
	NumVotes    int                           `json:"num_votes"`
	NumComments int                           `json:"num_comments"`
	Comments    []commentResp.CommentResponse `json:"comments,omitempty"`
	Votes       []votesResp.VoteReponse       `json:"votes,omitempty"`
	CreatedAt   time.Time                     `json:"created_at,omitempty"`
	UpdatedAt   time.Time                     `json:"updated_at,omitempty"`
}

func FromDomain(domain threads.Domain) ThreadResponse {
	var comments []commentResp.CommentResponse
	for _, getComment := range domain.Comments {
		comments = append(comments, commentResp.FromDomain(getComment))
	}

	var votes []votesResp.VoteReponse
	for _, getVote := range domain.Votes {
		votes = append(votes, votesResp.VoteReponse(getVote))
	}

	return ThreadResponse{
		ID:          domain.ID,
		Username:    domain.Username,
		User:        userResp.UserProfileFromDomain(domain.User),
		CategoryID:  domain.CategoryID,
		Title:       domain.Title,
		Content:     domain.Content,
		ImageUrl:    domain.ImageUrl,
		NumVotes:    domain.NumVotes,
		NumComments: domain.NumComments,
		Comments:    comments,
		Votes:       votes,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
