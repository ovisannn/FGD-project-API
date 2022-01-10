package comments

import (
	"disspace/business/comments"
	"disspace/controllers"
	"disspace/controllers/comments/request"
	"disspace/controllers/comments/response"
	"disspace/helpers/messages"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	CommentUseCase comments.UseCase
}

func NewCommentController(commentUseCase comments.UseCase) *CommentController {
	return &CommentController{
		CommentUseCase: commentUseCase,
	}
}

func (controller *CommentController) Create(c echo.Context) error {
	createComment := request.Comment{}
	c.Bind(&createComment)

	ctx := c.Request().Context()

	id := c.Param("id")
	c.Bind(&id)

	result, err := controller.CommentUseCase.Create(ctx, createComment.ToDomain(), id)
	if err != nil {
		if err == messages.ErrUnauthorizedUser {
			return controllers.NewErrorResponse(c, http.StatusUnauthorized, err)
		}
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomain(result))
}

func (controller *CommentController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	threadId := c.Param("thread_id")

	err := controller.CommentUseCase.Delete(ctx, id, threadId)
	if err != nil {
		if err == messages.ErrDataNotFound {
			return controllers.NewErrorResponse(c, http.StatusNotFound, err)
		}
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully delete comment")
}
