package reports

import (
	"context"
	"time"
)

type ReportUseCase struct {
	reportRepo     Repository
	contextTimeout time.Duration
}

func NewReportUseCase(reportRepository Repository, timeout time.Duration) UseCase {
	return &ReportUseCase{
		reportRepo:     reportRepository,
		contextTimeout: timeout,
	}
}

func (useCase *ReportUseCase) Create(ctx context.Context, reportDomain *Domain, id string) error {
	err := useCase.reportRepo.Create(ctx, reportDomain, id)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *ReportUseCase) GetAll(ctx context.Context, sort string) ([]Domain, error) {
	result, err := useCase.reportRepo.GetAll(ctx, sort)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}
