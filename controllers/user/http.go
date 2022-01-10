package user

import (
	"disspace/business/user"
	"disspace/controllers"
	"disspace/controllers/user/requests"
	"disspace/controllers/user/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase user.UseCase
}

func NewUserController(userUseCase user.UseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (controller *UserController) Register(c echo.Context) error {
	newUser := requests.User{}
	c.Bind(&newUser)

	ctx := c.Request().Context()
	result, err := controller.UserUseCase.Register(ctx, newUser.UserRegisterToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.UserFromDomain(result))
}
