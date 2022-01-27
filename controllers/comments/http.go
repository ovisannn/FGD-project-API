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
	commentId := c.Param("comment_id")

	err := controller.CommentUseCase.Delete(ctx, id, commentId)
	if err != nil {
		if err == messages.ErrDataNotFound {
			return controllers.NewErrorResponse(c, http.StatusNotFound, err)
		}
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully delete comment")
}

func (controller *CommentController) Search(c echo.Context) error {
	comments := []response.CommentResponse{}
	query := c.QueryParam("q")
	sorting := c.QueryParam("sort")
	ctx := c.Request().Context()

	result, err := controller.CommentUseCase.Search(ctx, query, sorting)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNotFound, err)
	}

	for _, item := range result {
		comments = append(comments, response.FromDomain(item))
	}
	return controllers.NewSuccessResponse(c, comments)
}

func (controller *CommentController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	result, err := controller.CommentUseCase.GetByID(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNotFound, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomain(result))
}

func (controller *CommentController) GetAllInThread(c echo.Context) error {
	comments := []response.CommentResponse{}
	threadId := c.Param("thread_id")
	parentId := c.Param("parent_id")
	option := c.QueryParam("option")
	ctx := c.Request().Context()

	result, err := controller.CommentUseCase.GetAllInThread(ctx, threadId, parentId, option)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNotFound, err)
	}
	for _, item := range result {
		comments = append(comments, response.FromDomain(item))
	}

	return controllers.NewSuccessResponse(c, comments)
}
