package quiz

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

//	@Tags		Quiz
//	@Accept		json
//	@Produce	json
//	@Param		quiz	body	models.Quiz	true	"quiz"
//	@Security	ApiKeyAuth
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/quiz [post]
func (qc *QuizController) Create(c *gin.Context) {
	roleID := c.GetUint("roleID")
	if roleID == 2 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "Admin is required",
					Message: "You are not admin user",
				},
			},
		})
		return
	} 

	userID := c.GetUint("userID")
	var quiz models.Quiz
	if err := c.ShouldBind(&quiz); err != nil{
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_DATA",
					Message: "Error on binding data",
				},
			},
		})
		return
	}
	id  ,err := qc.QuizRepository.Create(c , quiz , userID)
	if err != nil{
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_CREATE_QUIZ",
					Message: "Can't create quiz",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(200 , models.SuccessResponse{Result: id})
}
