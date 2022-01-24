package user

import (
	"disspace/app/middlewares"
	"disspace/business/user"
	"disspace/controllers"
	"disspace/controllers/user/requests"
	"disspace/controllers/user/responses"
	"disspace/helpers/messages"
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
	id := c.Param("id")
	//getting jwt payload
	token := c.Request().Header.Get("Authorization")
	payload, isOk := middlewares.ExtractClaims(token)
	if !isOk {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, messages.ErrFailedClaimJWT)
	}
	var username string = payload["username"].(string)
	//going down to usecase
	result, err := controller.UserUseCase.GetUserByID(ctx, id, username)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, responses.UserFromDomain(result))
}

func (controller *UserController) GetUserByUsername(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("username")
	//getting jwt payload
	token := c.Request().Header.Get("Authorization")
	payload, isOk := middlewares.ExtractClaims(token)
	if !isOk {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, messages.ErrFailedClaimJWT)
	}
	var username string = payload["username"].(string)
	//going down to usecase
	result, err := controller.UserUseCase.GetUserByUsername(ctx, id, username)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, responses.UserFromDomain(result))
}

func (controller *UserController) Follow(c echo.Context) error {
	ctx := c.Request().Context()
	usernameTarget := c.Param("usernameTarget")
	token := c.Request().Header.Get("Authorization")
	payload, isOk := middlewares.ExtractClaims(token)
	if !isOk {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, messages.ErrFailedClaimJWT)
	}
	var username string = payload["username"].(string)

	err := controller.UserUseCase.Follow(ctx, username, usernameTarget)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully follow user")
}

func (controller *UserController) Unfollow(c echo.Context) error {
	ctx := c.Request().Context()
	usernameTarget := c.Param("usernameTarget")
	token := c.Request().Header.Get("Authorization")
	payload, isOk := middlewares.ExtractClaims(token)
	if !isOk {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, messages.ErrFailedClaimJWT)
	}
	var username string = payload["username"].(string)
	err := controller.UserUseCase.Unfollow(ctx, username, usernameTarget)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully unfollow user")
}

func (controller *UserController) UpdateUserProfile(c echo.Context) error {
	ctx := c.Request().Context()
	dataUserProfile := requests.UserProfile{}
	c.Bind(&dataUserProfile)
	token := c.Request().Header.Get("Authorization")
	payload, isOk := middlewares.ExtractClaims(token)
	if !isOk {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, messages.ErrFailedClaimJWT)
	}
	var username string = payload["username"].(string)

	err := controller.UserUseCase.UpdateUserProfile(ctx, username, user.UserProfileDomain(dataUserProfile))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully update user profile")
}

func (controller *UserController) ChangePassword(c echo.Context) error {
	ctx := c.Request().Context()
	dataUser := requests.User{}
	c.Bind(&dataUser)
	token := c.Request().Header.Get("Authorization")
	payload, isOk := middlewares.ExtractClaims(token)
	if !isOk {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, messages.ErrFailedClaimJWT)
	}
	var username string = payload["username"].(string)
	err := controller.UserUseCase.ChangePassword(ctx, username, *dataUser.UserInfoUpdateToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully update user password")
}

func (controller *UserController) UpdateUserInfo(c echo.Context) error {
	ctx := c.Request().Context()
	dataUserInfo := requests.User{}
	c.Bind(&dataUserInfo)
	token := c.Request().Header.Get("Authorization")
	payload, isOk := middlewares.ExtractClaims(token)
	if !isOk {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, messages.ErrFailedClaimJWT)
	}
	var username string = payload["username"].(string)
	err := controller.UserUseCase.UpdateUserInfo(ctx, username, *dataUserInfo.UserInfoUpdateToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully update user information")
}

func (controller *UserController) Logout(c echo.Context) error {
	ctx := c.Request().Context()
	dataSession := requests.UserSession{
		Token:    c.Param("token"),
		Username: c.Param("username"),
	}
	err := controller.UserUseCase.Logout(ctx, dataSession.SessionToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully logout")
}

func (controller *UserController) GetModeratorsByCategoryID(c echo.Context) error {
	moderators := []responses.UserProfile{}
	ctx := c.Request().Context()
	categoryID := c.Param("categoryID")
	result, err := controller.UserUseCase.GetModerators(ctx, categoryID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	for _, i := range result {
		moderators = append(moderators, responses.UserProfileFromDomain(i))
	}
	return controllers.NewSuccessResponse(c, moderators)
}

func (controller *UserController) GetTop5User(c echo.Context) error {
	// fmt.Print("aaa")
	topUser := []responses.UserProfile{}
	ctx := c.Request().Context()
	result, err := controller.UserUseCase.GetTop5User(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	counter := 0
	for _, i := range result {
		if counter == 5 {
			break
		}
		topUser = append(topUser, responses.UserProfileFromDomain(i))
		counter += 1
	}
	return controllers.NewSuccessResponse(c, topUser)
}

func (controller *UserController) Test(c echo.Context) error {
	// a := middlewares.GetUserId(c)
	// fmt.Print(a)
	return controllers.NewSuccessResponse(c, "hello")
}
