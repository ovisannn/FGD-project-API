package routes

import (
	"disspace/controllers/threads"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	ThreadController threads.ThreadController
}

func (ctrl *ControllerList) RouteRegister(e *echo.Echo) {
	baseRoute := e.Group("/v1")

	// Threads
	baseRoute.GET("/threads", ctrl.ThreadController.GetAll)
	baseRoute.POST("/threads", ctrl.ThreadController.Create)
	baseRoute.DELETE("/threads/:id", ctrl.ThreadController.Delete)
	baseRoute.PATCH("/threads/:id", ctrl.ThreadController.Update)
}
