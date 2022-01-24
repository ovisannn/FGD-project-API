package response

import (
	"disspace/business/votes"
	"time"
)

type VoteReponse struct {
	ID          string    `json:"_id"`
	Username      string    `json:"username,omitempty"`
	ReferenceID string    `json:"reference_id,omitempty"`
	Status      int       `json:"status,omitempty"`
	TimeLike    time.Time `json:"time_like,omitempty"`
}

func FromDomain(domain votes.Domain) VoteReponse {
	return VoteReponse{
		ID:          domain.ID,
		Username:      domain.Username,
		ReferenceID: domain.ReferenceID,
		Status:      domain.Status,
		TimeLike:    domain.TimeLike,
	}
}
