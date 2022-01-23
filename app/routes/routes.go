package routes

import (
	"disspace/controllers/categories"
	"disspace/controllers/comments"
	"disspace/controllers/reports"
	"disspace/controllers/threads"
	"disspace/controllers/user"

	// "disspace/controllers/user"

	// "disspace/controllers/user"
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
	jwtAuth := middleware.JWTWithConfig(ctrl.JWTConfig)

	// Threads
	baseRoute.GET("/threads", ctrl.ThreadController.GetAll)
	baseRoute.POST("/threads", ctrl.ThreadController.Create, jwtAuth)
	baseRoute.DELETE("/threads/:id", ctrl.ThreadController.Delete, jwtAuth)
	baseRoute.GET("/threads/:id", ctrl.ThreadController.GetByID)
	baseRoute.PATCH("/threads/:id", ctrl.ThreadController.Update, jwtAuth)

	//categories
	baseRoute.GET("/categories", ctrl.CategoriesController.GetAll)
	baseRoute.GET("/categories/:id", ctrl.CategoriesController.GetByID)
	baseRoute.POST("/categories", ctrl.CategoriesController.Create)
	baseRoute.DELETE("/categories/:id", ctrl.CategoriesController.Delete)
	baseRoute.PATCH("/categories/:id", ctrl.CategoriesController.Update)

	// Votes
	baseRoute.POST("/users/:id/votes", ctrl.VoteController.Store, jwtAuth)
	baseRoute.PUT("/users/:id/votes/:ref_id", ctrl.VoteController.Update, jwtAuth)

	//user
	baseRoute.POST("/user/register", ctrl.UserController.Register)                            //ok
	baseRoute.GET("/userProfile/:username", ctrl.UserController.UserProfileGetByUsername)     //ok
	baseRoute.POST("/user/login", ctrl.UserController.Login)                                  //ok
	baseRoute.GET("/user/id/:id", ctrl.UserController.GetUserByID, jwtAuth)                   //ok
	baseRoute.GET("/user/username/:username", ctrl.UserController.GetUserByUsername, jwtAuth) //ok
	baseRoute.PATCH("/user/follow/:usernameTarget", ctrl.UserController.Follow, jwtAuth)      //ok
	baseRoute.PATCH("/user/unfollow/:usernameTarget", ctrl.UserController.Unfollow, jwtAuth)  //ok
	baseRoute.PATCH("/userProfile/update", ctrl.UserController.UpdateUserProfile, jwtAuth)    //ok
	baseRoute.PATCH("/user/update", ctrl.UserController.UpdateUserInfo, jwtAuth)              //ok
	baseRoute.PATCH("/user/newPassword", ctrl.UserController.ChangePassword, jwtAuth)         //ok
	baseRoute.DELETE("/user/logout", ctrl.UserController.Logout)                              //rework or maybe dihapus aja

	//leaderboard
	//get leaderboard -> GET

	//moderators
	//get moderators -> GET

	// Comments
	baseRoute.POST("/users/:id/comments", ctrl.CommentController.Create, jwtAuth)
	baseRoute.DELETE("/users/:id/comments/:thread_id", ctrl.CommentController.Delete, jwtAuth)

	// Reports (User, Thread, Comment)
	baseRoute.PUT("/users/:id/reporting", ctrl.ReportController.Create, jwtAuth)
	baseRoute.GET("/reports", ctrl.ReportController.GetAll)

	// Search (Users, Threads, Comments)
	baseRoute.GET("/threads/search", ctrl.ThreadController.Search)
	baseRoute.GET("/comments/search", ctrl.CommentController.Search)

	baseRoute.GET("/test", ctrl.UserController.Test, jwtAuth)
}
