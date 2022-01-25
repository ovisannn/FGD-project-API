package response

import (
	"disspace/business/comments"
	userResp "disspace/controllers/user/responses"
	votesResp "disspace/controllers/votes/response"
	"time"
)

type CommentResponse struct {
	ID          string                  `json:"_id"`
	ThreadID    string                  `json:"thread_id,omitempty"`
	ParentID    string                  `json:"parent_id,omitempty"`
	Username    string                  `json:"username,omitempty"`
	User        userResp.UserProfile    `json:"user,omitempty"`
	Text        string                  `json:"text,omitempty"`
	NumVotes    int                     `json:"num_votes"`
	NumComments int                     `json:"num_comments"`
	Votes       []votesResp.VoteReponse `json:"votes,omitempty"`
	CreatedAt   time.Time               `json:"created_at,omitempty"`
	UpdatedAt   time.Time               `json:"updated_at,omitempty"`
}

func FromDomain(domain comments.Domain) CommentResponse {
	var votes []votesResp.VoteReponse
	for _, getVote := range domain.Votes {
		votes = append(votes, votesResp.VoteReponse(getVote))
	}

	return CommentResponse{
		ID:          domain.ID,
		ThreadID:    domain.ThreadID,
		ParentID:    domain.ParentID,
		Username:    domain.Username,
		User:        userResp.UserProfileFromDomain(domain.User),
		Text:        domain.Text,
		NumVotes:    domain.NumVotes,
		NumComments: domain.NumComments,
		Votes:       votes,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
