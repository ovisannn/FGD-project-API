package request

import (
	"disspace/business/reports"
	"time"
)

type Report struct {
	ReportedBy  string    `json:"reported_by" bson:"reported_by" param:"id"`
	TargetID    string    `json:"target_id" bson:"target_id" param:"target_id"`
	TargetType  int       `json:"target_type" bson:"target_type"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	UniqueID    string    `json:"unique_id" bson:"unique_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}

func (request *Report) ToDomain() *reports.Domain {
	return &reports.Domain{
		ReportedBy: request.ReportedBy,
		TargetID:   request.TargetID,
		TargetType: request.TargetType,
		UniqueID:   request.ReportedBy + request.TargetID,
		CreatedAt:  time.Now(),
	}
}
