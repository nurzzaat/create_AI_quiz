package quiz

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

// @Tags		Student
// @Accept		json
// @Produce	json
// @Param		quizId		path	int					true	"quizId"
// @Param		data	body	models.Submission	true	"data"
// @Security	ApiKeyAuth
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/quiz/submit/{quizId} [post]
func (qc *QuizController) Submit(c *gin.Context) {
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
	quizId, _ := strconv.Atoi(c.Param("quizId"))

	var submission models.Submission
	if err := c.ShouldBind(&submission); err != nil {
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

	for _, answer := range submission.Answers {
		submission.Answer += answer + `,`
	}
	submission.Answer = submission.Answer[:len(submission.Answer)-1]
	fmt.Println(submission.Answer)

	err := qc.QuizRepository.Submit(c, quizId, userID, submission)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_SUBMIT_QUIZ",
					Message: "Can't submit quiz",
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
