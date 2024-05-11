package quiz

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

//	@Tags		Quiz
//	@Accept		json
//	@Produce	json
//	@Param		search	query	string	false	"search"
//	@Security	ApiKeyAuth
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/quiz/admin [get]
func (qc *QuizController) GetAllAdmin(c *gin.Context) {
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
	search := `%` + c.Query("search") + `%`

	quizes, err := qc.QuizRepository.GetAllAdmin(c, userID , search)
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
