package comments

import (
	"disspace/business/comments"
	"disspace/drivers/databases/user"
	"time"
)

type Comment struct {
	ID          string           `json:"_id,omitempty" bson:"_id,omitempty"`
	ThreadID    string           `json:"thread_id,omitempty" bson:"thread_id,omitempty" param:"thread_id"`
	ParentID    string           `json:"parent_id,omitempty" bson:"parent_id,omitempty" param:"parent_id"`
	Username    string           `json:"username,omitempty" bson:"username,omitempty" param:"id"`
	User        user.UserProfile `json:"user,omitempty" bson:"user,omitempty"`
	Text        string           `json:"text,omitempty" bson:"text,omitempty"`
	NumVotes    int              `json:"num_votes,omitempty" bson:"num_votes,omitempty"`
	NumComments int              `json:"num_comments,omitempty" bson:"num_comments,omitempty"`
	CreatedAt   time.Time        `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" bson:"updated_at"`
}

func (record *Comment) ToDomain() comments.Domain {
	return comments.Domain{
		ID:          record.ID,
		ThreadID:    record.ThreadID,
		ParentID:    record.ParentID,
		Username:    record.Username,
		User:        record.User.UserProfileToDomain(),
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
		ParentID:    domain.ParentID,
		Username:    domain.Username,
		Text:        domain.Text,
		NumVotes:    domain.NumVotes,
		NumComments: domain.NumComments,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
