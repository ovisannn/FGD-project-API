package votes

import (
	"disspace/business/votes"
	"disspace/controllers"
	"disspace/controllers/votes/request"
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
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully liked")
}
