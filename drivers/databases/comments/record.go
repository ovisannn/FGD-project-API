package comments

import (
	"disspace/business/comments"
	"time"
)

type Comment struct {
	ID          string    `json:"_id,omitempty" bson:"_id,omitempty"`
	ThreadID    string    `json:"thread_id,omitempty" bson:"thread_id,omitempty"`
	UserID      string    `json:"user_id,omitempty" bson:"user_id,omitempty" param:"id"`
	Text        string    `json:"text,omitempty" bson:"text,omitempty"`
	NumVotes    int       `json:"num_votes,omitempty" bson:"num_votes,omitempty"`
	NumComments int       `json:"num_comments,omitempty" bson:"num_comments,omitempty"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func (record *Comment) ToDomain() comments.Domain {
	return comments.Domain{
		ID:          record.ID,
		ThreadID:    record.ThreadID,
		UserID:      record.UserID,
		Text:        record.Text,
		NumVotes:    record.NumVotes,
		NumComments: record.NumComments,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
}

func FromDomain(domain comments.Domain) Comment {
	return Comment{
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
