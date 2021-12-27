package threads

import (
	"disspace/business/threads"
	"disspace/controllers"
	"disspace/controllers/threads/requests"
	"disspace/controllers/threads/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ThreadController struct {
	ThreadUseCase threads.UseCase
}

func NewThreadController(threadUseCase threads.UseCase) *ThreadController {
	return &ThreadController{
		ThreadUseCase: threadUseCase,
	}
}

func (controller *ThreadController) GetAll(c echo.Context) error {
	threads := []responses.ThreadResponse{}
	ctx := c.Request().Context()

	result, err := controller.ThreadUseCase.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	for _, item := range result {
		threads = append(threads, responses.FromDomain(item))
	}
	return controllers.NewSuccessResponse(c, threads)
}

func (controller *ThreadController) Create(c echo.Context) error {
	createThread := requests.Thread{}
	c.Bind(&createThread)

	ctx := c.Request().Context()

	result, err := controller.ThreadUseCase.Create(ctx, createThread.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(result))
}

func (controller *ThreadController) GetByID(c echo.Context) error {
  ctx := c.Request().Context()

	id := c.Param("id")

	result, err := controller.ThreadUseCase.GetByID(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
    	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(result))
}

func (controller *ThreadController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	err := controller.ThreadUseCase.Delete(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully deleted thread")
}


func (controller *ThreadController) Update(c echo.Context) error {
	updateThread := requests.Thread{}
	c.Bind(&updateThread)
  
  ctx := c.Request().Context()

	id := c.Param("id")
  
	err := controller.ThreadUseCase.Update(ctx, updateThread.ToDomain(), id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully update thread")
}
