package moderators

import (
	"disspace/business/moderators"
	"disspace/controllers"
	responses "disspace/controllers/moderators/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ModeratorsController struct {
	ModeratorsUseCase moderators.UseCase
}

func NewModeratorsController(moderatorsUseCase moderators.UseCase) *ModeratorsController {
	return &ModeratorsController{
		ModeratorsUseCase: moderatorsUseCase,
	}
}

func (controller *ModeratorsController) GetByCategoryID(c echo.Context) error {
	moderators := []responses.Response{}
	ctx := c.Request().Context()
	categoryID := c.Param("categoryID")
	result, err := controller.ModeratorsUseCase.GetByCategoryID(ctx, categoryID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	for _, i := range result {
		moderators = append(moderators, responses.FromDomain(i))
	}

	return controllers.NewSuccessResponse(c, moderators)
}
