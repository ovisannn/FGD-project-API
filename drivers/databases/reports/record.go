package reports

import (
	"disspace/business/reports"
	comDb "disspace/drivers/databases/comments"
	threadDb "disspace/drivers/databases/threads"
	userDb "disspace/drivers/databases/user"
	"time"
)

type Report struct {
	// ID          string    `json:"_id,omitempty" bson:"_id,omitempty"`
	ReportedBy  string             `json:"reported_by" bson:"reported_by" param:"id"`
	TargetID    string             `json:"target_id" bson:"target_id" param:"target_id"`
	TargetType  int                `json:"target_type" bson:"target_type"`
	Thread      threadDb.Thread    `json:"thread,omitempty" bson:"thread,omitempty"`
	Comment     comDb.Comment      `json:"comment,omitempty" bson:"comment,omitempty"`
	User        userDb.UserProfile `json:"user,omitempty" bson:"user,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	UniqueID    string             `json:"unique_id" bson:"unique_id"`
	Count       int                `json:"count,omitempty" bson:"count,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

func (record *Report) ToDomain() reports.Domain {
	return reports.Domain{
		// ID:          record.ID,
		ReportedBy:  record.ReportedBy,
		TargetID:    record.TargetID,
		TargetType:  record.TargetType,
		Thread:      record.Thread.ToDomain(),
		Comment:     record.Comment.ToDomain(),
		User:        record.User.UserProfileToDomain(),
		Description: record.Description,
		UniqueID:    record.UniqueID,
		Count:       record.Count,
		CreatedAt:   record.CreatedAt,
	}
}

func FromDomain(domain reports.Domain) Report {
	return Report{
		// ID:          domain.ID,
		ReportedBy:  domain.ReportedBy,
		TargetID:    domain.TargetID,
		TargetType:  domain.TargetType,
		Description: domain.Description,
		UniqueID:    domain.UniqueID,
		Count:       domain.Count,
		CreatedAt:   domain.CreatedAt,
	}
}
