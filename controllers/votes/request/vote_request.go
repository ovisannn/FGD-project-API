package request

import (
	"disspace/business/votes"
	"time"
)

type Vote struct {
	Username    string    `json:"username" bson:"username" param:"id"`
	ReferenceID string    `json:"reference_id" bson:"reference_id"`
	Status      int       `json:"status" bson:"status"`
	TimeLike    time.Time `json:"time_like" bson:"time_like"`
}

type UpdateVote struct {
	Status int `json:"status" bson:"status"`
}

func (request *Vote) ToDomain() *votes.Domain {
	return &votes.Domain{
		Username:    request.Username,
		ReferenceID: request.ReferenceID,
		Status:      request.Status,
		TimeLike:    time.Now(),
	}
}
