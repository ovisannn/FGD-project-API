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

func (controller *UserController) UserProfileGetByUsername(c echo.Context) error {
	ctx := c.Request().Context()
	username := c.Param("username")

	result, err := controller.UserUseCase.UserProfileGetByUsername(ctx, username)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, responses.UserProfileFromDomain(result))
}

func (controller *UserController) Login(c echo.Context) error {
	loginInfo := requests.LoginInfo{}
	c.Bind(&loginInfo)

	ctx := c.Request().Context()
	result, err := controller.UserUseCase.Login(ctx, loginInfo.Username, loginInfo.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, responses.SessionFromDomain(result))
}

func (controller *UserController) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	// user := responses.User{}
	dataSession := requests.UserSession{}
	id := c.Param("id")
	c.Bind(&dataSession)
	// fmt.Println("id : " + id)
	result, err := controller.UserUseCase.GetUserByID(ctx, id, dataSession.SessionToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, responses.UserFromDomain(result))
}

func (controller *UserController) Follow(c echo.Context) error {
	ctx := c.Request().Context()
	dataSession := requests.UserSession{}
	c.Bind(&dataSession)

	username := c.Param("username")
	usernameTarget := c.Param("usernameTarget")

	err := controller.UserUseCase.Follow(ctx, username, usernameTarget, dataSession.SessionToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully follow user")
}
