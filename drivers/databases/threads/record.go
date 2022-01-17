package threads

import (
	"disspace/business/comments"
	"disspace/business/threads"
	commentDb "disspace/drivers/databases/comments"

	"time"
)

type Thread struct {
	ID          string              `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID      string              `json:"user_id,omitempty" bson:"user_id,omitempty"`
	CategoryID  string              `json:"category_id,omitempty" bson:"category_id,omitempty"`
	Title       string              `json:"title,omitempty" bson:"title,omitempty"`
	Content     string              `json:"content,omitempty" bson:"content,omitempty"`
	ImageUrl    string              `json:"image_url,omitempty" bson:"image_url,omitempty"`
	NumVotes    int                 `json:"num_votes,omitempty" bson:"num_votes,omitempty"`
	NumComments int                 `json:"num_comments,omitempty" bson:"num_comments,omitempty"`
	Comments    []commentDb.Comment `json:"comments,omitempty" bson:"comments,omitempty"`
	CreatedAt   time.Time           `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at" bson:"updated_at"`
}

func (record *Thread) ToDomain() threads.Domain {
	var comments []comments.Domain
	for _, comment := range record.Comments {
		comments = append(comments, comment.ToDomain())
	}

	return threads.Domain{
		ID:          record.ID,
		UserID:      record.UserID,
		CategoryID:  record.CategoryID,
		Title:       record.Title,
		Content:     record.Content,
		ImageUrl:    record.ImageUrl,
		NumVotes:    record.NumVotes,
		NumComments: record.NumComments,
		Comments:    comments,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
}

func FromDomain(domain threads.Domain) Thread {
	var comments []commentDb.Comment
	for _, comment := range domain.Comments {
		comments = append(comments, commentDb.FromDomain(comment))
	}

	return Thread{
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

func FromDomainUpdate(domain threads.Domain) Thread {
	return Thread{
		Title:     domain.Title,
		Content:   domain.Content,
		ImageUrl:  domain.ImageUrl,
		UpdatedAt: domain.UpdatedAt,
	}
}
