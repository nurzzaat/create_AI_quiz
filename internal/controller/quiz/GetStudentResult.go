package quiz

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

//	@Tags		Student
//	@Accept		json
//	@Produce	json
//	@Param		quizId		path	int	true	"quizId"
//	@Security	ApiKeyAuth
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/students/quiz/{quizId}/result [get]
func (qc *QuizController) GetStudentResult(c *gin.Context) {
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
	quizId, _ := strconv.Atoi(c.Param("quizId"))
	studentId := c.GetUint("userID")

	result , err := qc.QuizRepository.GetStudentResult(c , quizId , int(studentId))
	if err != nil{
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
	c.JSON(200 , models.SuccessResponse{Result: result})
}