package routes

import (
	"disspace/controllers/categories"
	"disspace/controllers/comments"
	"disspace/controllers/reports"
	"disspace/controllers/threads"
	"disspace/controllers/user"
	"disspace/controllers/votes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTConfig            middleware.JWTConfig
	ThreadController     threads.ThreadController
	CategoriesController categories.CategoriesController
	VoteController       votes.VoteController
	UserController       user.UserController
	CommentController    comments.CommentController
	ReportController     reports.ReportController
}

func (ctrl *ControllerList) RouteRegister(e *echo.Echo) {
	baseRoute := e.Group("/v1")
	// jwtAuth := middleware.JWTWithConfig(cl.JwtConfig)

	// Threads
	baseRoute.GET("/threads", ctrl.ThreadController.GetAll)
	baseRoute.POST("/threads", ctrl.ThreadController.Create)
	baseRoute.DELETE("/threads/:id", ctrl.ThreadController.Delete)
	baseRoute.GET("/threads/:id", ctrl.ThreadController.GetByID)
	baseRoute.PATCH("/threads/:id", ctrl.ThreadController.Update)

	//categories
	baseRoute.GET("/categories", ctrl.CategoriesController.GetAll)
	baseRoute.GET("/categories/:id", ctrl.CategoriesController.GetByID)
	baseRoute.POST("/categories", ctrl.CategoriesController.Create)
	baseRoute.DELETE("/categories/:id", ctrl.CategoriesController.Delete)
	baseRoute.PATCH("/categories/:id", ctrl.CategoriesController.Update)

	// Votes
	baseRoute.POST("/users/:id/votes", ctrl.VoteController.Store)
	baseRoute.PUT("/users/:id/votes/:ref_id", ctrl.VoteController.Update)

	//user
	baseRoute.POST("/user/register", ctrl.UserController.Register) // belum encrypt password
	baseRoute.GET("/userProfile/:username", ctrl.UserController.UserProfileGetByUsername)
	baseRoute.POST("/user/login", ctrl.UserController.Login) // belum decrypt password
	baseRoute.GET("/user/:id", ctrl.UserController.GetUserByID)
	baseRoute.PATCH("user/follow/:username/:usernameTarget", ctrl.UserController.Follow)
	// baseRoute.PATCH("user/unfollow/:username/:usernameTarget")
	// baseRoute.PATCH("userProfile/:username")
	// baseRoute.PATCH("user/:username")
	// baseRoute.DELETE("user/logout/:username")

	//leaderboard
	//get leaderboard -> GET

	//moderators
	//get moderators -> GET

	// Comments
	baseRoute.POST("/users/:id/comments", ctrl.CommentController.Create)
	baseRoute.DELETE("/users/:id/comments/:thread_id", ctrl.CommentController.Delete)

	// Reports (User, Thread, Comment)
	baseRoute.PUT("/users/:id/reporting", ctrl.ReportController.Create)
}
