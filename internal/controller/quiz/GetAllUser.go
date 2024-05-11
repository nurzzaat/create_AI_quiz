package quiz

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

//	@Tags		Student
//	@Accept		json
//	@Produce	json
//	@Param		search	query	string	false	"search"
//	@Security	ApiKeyAuth
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/quiz/user [get]
func (qc *QuizController) GetAllUser(c *gin.Context) {
	roleID := c.GetUint("roleID")
	if roleID == 1 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "User is required",
					Message: "You are not admin user",
				},
			},
		})
		return
	} 
	userID := c.GetUint("userID")
	search := `%` + c.Query("search") + `%`

	quizes, err := qc.QuizRepository.GetAllUser(c, userID , search)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_QUIZ",
					Message: "Can't get quiz",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(200, models.SuccessResponse{Result: quizes})
}