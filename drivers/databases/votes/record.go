package votes

import (
	"disspace/business/votes"
	"time"
)

type Vote struct {
	ID          string    `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID      string    `json:"user_id" bson:"user_id" param:"id"`
	ReferenceID string    `json:"reference_id" bson:"reference_id"`
	Status      int       `json:"status" bson:"status"`
	TimeLike    time.Time `json:"time_like" bson:"time_like"`
}

func (record *Vote) ToDomain() votes.Domain {
	return votes.Domain{
		ID:          record.ID,
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,
		Status:      record.Status,
		TimeLike:    record.TimeLike,
	}
}

func FromDomain(domain votes.Domain) Vote {
	return Vote{
		ID:          domain.ID,
		UserID:      domain.UserID,
		ReferenceID: domain.ReferenceID,
		Status:      domain.Status,
		TimeLike:    domain.TimeLike,
	}
}
