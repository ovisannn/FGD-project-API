package threads

import (
	"disspace/business/threads"
	"disspace/controllers"
	"disspace/controllers/threads/responses"

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
		return err
	}

	for _, item := range result {
		threads = append(threads, responses.FromDomain(item))
	}
	return controllers.NewSuccessResponse(c, threads)
}
