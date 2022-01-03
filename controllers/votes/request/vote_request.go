package request

import (
	"disspace/business/votes"
	"time"
)

type Vote struct {
	UserID      string    `json:"user_id" bson:"user_id"`
	ReferenceID string    `json:"reference_id" bson:"reference_id"`
	Status      int       `json:"status" bson:"status"`
	TimeLike    time.Time `json:"time_like" bson:"time_like"`
}

func (request *Vote) ToDomain() *votes.Domain {
	return &votes.Domain{
		UserID:      request.UserID,
		ReferenceID: request.ReferenceID,
		Status:      request.Status,
		TimeLike:    time.Now(),
	}
}
