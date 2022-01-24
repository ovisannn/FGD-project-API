package moderators

import (
	"disspace/business/moderators"
)

type Moderators struct {
	ID         string `bson:"_id,omitempty" json:"_id"`
	Username   string `bson:"username" json:"username"`
	CategoryID string `bson:"categori_id" json:"categori_id"`
}

func (record *Moderators) SessionToDomain() moderators.Domain {
	return moderators.Domain{
		ID:         record.ID,
		Username:   record.Username,
		CategoryID: record.CategoryID,
	}
}

func FromDomain(domain moderators.Domain) Moderators {
	return Moderators{
		ID:         domain.ID,
		Username:   domain.Username,
		CategoryID: domain.CategoryID,
	}
}
