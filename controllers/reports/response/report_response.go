package response

import (
	"disspace/business/reports"
	comResp "disspace/controllers/comments/response"
	threadResp "disspace/controllers/threads/responses"
	userResp "disspace/controllers/user/responses"
	"time"
)

type ReportResponse struct {
	// ID          string    `json:"_id,omitempty"`
	ReportedBy  string                    `json:"reported_by"`
	TargetID    string                    `json:"target_id"`
	TargetType  int                       `json:"target_type"`
	Thread      threadResp.ThreadResponse `json:"thread,omitempty"`
	Comment     comResp.CommentResponse   `json:"comment,omitempty"`
	User        userResp.UserProfile      `json:"user,omitempty"`
	Description string                    `json:"description,omitempty"`
	UniqueID    string                    `json:"unique_id"`
	Count       int                       `json:"count,omitempty"`
	CreatedAt   time.Time                 `json:"created_at"`
}

func FromDomain(domain reports.Domain) ReportResponse {
	return ReportResponse{
		// ID:          domain.ID,
		ReportedBy:  domain.ReportedBy,
		TargetID:    domain.TargetID,
		TargetType:  domain.TargetType,
		Thread:      threadResp.FromDomain(domain.Thread),
		Comment:     comResp.FromDomain(domain.Comment),
		User:        userResp.UserProfileFromDomain(domain.User),
		Description: domain.Description,
		UniqueID:    domain.UniqueID,
		Count:       domain.Count,
		CreatedAt:   domain.CreatedAt,
	}
}
