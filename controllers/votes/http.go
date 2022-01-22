package votes

import (
	"disspace/business/votes"
	"disspace/controllers"
	"disspace/controllers/votes/response"
	"disspace/controllers/votes/request"
	"disspace/helpers/messages"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VoteController struct {
	VoteUseCase votes.UseCase
}

func NewVoteController(voteUseCase votes.UseCase) *VoteController {
	return &VoteController{
		VoteUseCase: voteUseCase,
	}
}

func (controller *VoteController) Store(c echo.Context) error {
	storeVote := request.Vote{}
	c.Bind(&storeVote)

	ctx := c.Request().Context()

	id := c.Param("id")
	c.Bind(&id)

	err := controller.VoteUseCase.Store(ctx, storeVote.ToDomain(), id)
	if err != nil {
		if err == messages.ErrUnauthorizedUser {
			return controllers.NewErrorResponse(c, http.StatusUnauthorized, err)
		} else if err == messages.ErrReferenceNotFound {
			return controllers.NewErrorResponse(c, http.StatusNotFound, err)
		}
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully liked")
}

func (controller *VoteController) Update(c echo.Context) error {
	updateVote := request.UpdateVote{}
	c.Bind(&updateVote)

	ctx := c.Request().Context()

	id := c.Param("id")
	refid := c.Param("ref_id")

	err := controller.VoteUseCase.Update(ctx, updateVote.Status, id, refid)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully update vote")

}

func (controller *VoteController) GetIsVoted(c echo.Context) error {
	ctx := c.Request().Context()

	username := c.Param("id")
	refId := c.Param("ref_id")

	result, err := controller.VoteUseCase.GetIsVoted(ctx, username, refId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNotFound, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomain(result))

}
