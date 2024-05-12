package quiz

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

//	@Tags		Quiz
//	@Accept		json
//	@Produce	json
//	@Param		quizId		query	int	true	"quizId"
//	@Param		studentId	query	int	true	"studentId"
//	@Security	ApiKeyAuth
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/students/quiz/add [post]
func (qc *QuizController) AddStudentToQuizPOST(c *gin.Context) {
	roleID := c.GetUint("roleID")
	if roleID == 2 {
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
	quizId, _ := strconv.Atoi(c.Query("quizId"))
	studentId, _ := strconv.Atoi(c.Query("studentId"))

	err := qc.QuizRepository.AddStudentToQuiz(c, quizId , studentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_ADD_TO_QUIZ",
					Message: "Can't add to quiz",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(200, models.SuccessResponse{Result: "Success"})
}


//	@Tags		Quiz
//	@Accept		json
//	@Produce	json
//	@Param		quizId		path	int	true	"quizId"
//	@Param		studentId	path	int	true	"studentId"
//	@Security	ApiKeyAuth
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/students/quiz/{quizId}/add/{studentId} [put]
func (qc *QuizController) AddStudentToQuizPUT(c *gin.Context) {
	roleID := c.GetUint("roleID")
	if roleID == 2 {
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
	studentId, _ := strconv.Atoi(c.Param("studentId"))

	err := qc.QuizRepository.AddStudentToQuiz(c, quizId , studentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_ADD_TO_QUIZ",
					Message: "Can't add to quiz",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(200, models.SuccessResponse{Result: "Success"})
}
