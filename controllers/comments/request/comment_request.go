package request

import (
	"disspace/business/comments"
	"time"
)

type Comment struct {
	ThreadID  string    `json:"thread_id" bson:"thread_id"`
	ParentID  string    `json:"parent_id" bson:"parent_id"`
	Username  string    `json:"username" bson:"username" param:"id"`
	Text      string    `json:"text" bson:"text"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (request *Comment) ToDomain() *comments.Domain {
	return &comments.Domain{
		ThreadID:  request.ThreadID,
		ParentID:  request.ParentID,
		Username:  request.Username,
		Text:      request.Text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
