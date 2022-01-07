package reports

import (
	"disspace/business/reports"
	"disspace/controllers"
	"disspace/controllers/reports/request"
	"disspace/helpers/messages"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReportController struct {
	ReportUseCase reports.UseCase
}

func NewReportController(reportUseCase reports.UseCase) *ReportController {
	return &ReportController{
		ReportUseCase: reportUseCase,
	}
}

func (controller *ReportController) Create(c echo.Context) error {
	createReport := request.Report{}
	id := c.Param("id")

	c.Bind(&createReport)
	c.Bind(&id)

	ctx := c.Request().Context()
	err := controller.ReportUseCase.Create(ctx, createReport.ToDomain(), id)
	if err != nil {
		if err == messages.ErrUnauthorizedUser {
			return controllers.NewErrorResponse(c, http.StatusUnauthorized, err)
		} else if err == messages.ErrInvalidUserID {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusConflict, err)
	}
	return controllers.NewSuccessResponse(c, "successfully reported target")
}
