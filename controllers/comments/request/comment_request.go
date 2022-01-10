package request

import (
	"disspace/business/comments"
	"time"
)

type Comment struct {
	ThreadID  string    `json:"thread_id" bson:"thread_id"`
	UserID    string    `json:"user_id" bson:"user_id" param:"id"`
	Text      string    `json:"text" bson:"text"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (request *Comment) ToDomain() *comments.Domain {
	return &comments.Domain{
		ThreadID:  request.ThreadID,
		UserID:    request.UserID,
		Text:      request.Text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
