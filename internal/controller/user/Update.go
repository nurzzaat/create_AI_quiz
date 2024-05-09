package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

// @Tags		User
// @Accept		json
// @Produce	json
// @Param		user	body	models.User	true	"user"
// @Security	ApiKeyAuth
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/user/profile [put]
func (sc *UserController) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("userID")

	var userInfo models.User
	userInfo.ID = userID
	if err := c.ShouldBind(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of signin",
				},
			},
		})
		return
	}

	id, err := sc.UserRepository.EditUser(c, userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_UPDATE_USER_PROFILE",
					Message: "Can't update profile from db",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: id})
}
