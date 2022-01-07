package reports

import (
	"disspace/business/reports"
	"time"
)

type Report struct {
	ID          string    `json:"_id,omitempty" bson:"_id,omitempty"`
	ReportedBy  string    `json:"reported_by" bson:"reported_by" param:"id"`
	TargetID    string    `json:"target_id" bson:"target_id" param:"target_id"`
	TargetType  int       `json:"target_type" bson:"target_type"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	UniqueID    string    `json:"unique_id" bson:"unique_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}

func (record *Report) ToDomain() reports.Domain {
	return reports.Domain{
		ID:          record.ID,
		ReportedBy:  record.ReportedBy,
		TargetID:    record.TargetID,
		TargetType:  record.TargetType,
		Description: record.Description,
		UniqueID:    record.UniqueID,
		CreatedAt:   record.CreatedAt,
	}
}

func FromDomain(domain reports.Domain) Report {
	return Report{
		ID:          domain.ID,
		ReportedBy:  domain.ReportedBy,
		TargetID:    domain.TargetID,
		TargetType:  domain.TargetType,
		Description: domain.Description,
		UniqueID:    domain.UniqueID,
		CreatedAt:   domain.CreatedAt,
	}
}
