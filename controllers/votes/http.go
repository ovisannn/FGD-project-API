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

// func (controller *VoteController) Create(c echo.Context) error {
// 	createVote := request.Vote{}
// 	c.Bind(&createVote)
// 	id := c.Param("id")

// 	ctx := c.Request().Context()

// 	err := controller.VoteUseCase.Create(ctx, createVote.ToDomain(), id)
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}
// 	return controllers.NewSuccessResponse(c, "successfully liked")
// }
func (controller *VoteController) Create(c echo.Context) error {
	createVote := request.Vote{}
	c.Bind(&createVote)

	ctx := c.Request().Context()

	err := controller.VoteUseCase.Create(ctx, createVote.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully liked")
}
