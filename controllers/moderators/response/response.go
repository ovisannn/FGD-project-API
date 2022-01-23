package responses

import "disspace/business/moderators"

type Response struct {
	ID         string `bson:"_id,omitempty" json:"_id,omitempty"`
	Username   string `bson:"username,omitempty" json:"username,omitempty"`
	CategoryID string `bson:"categori_id,omitempty" json:"categori_id,omitempty"`
}

func FromDomain(domain moderators.Domain) Response {
	return Response{
		ID:         domain.ID,
		Username:   domain.Username,
		CategoryID: domain.CategoryID,
	}
}
