package quiz

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	openai "github.com/sashabaranov/go-openai"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

type QuizController struct {
	QuizRepository models.QuizRepository
}

// @Tags		Quiz
// @Param		text	formData	string	true	"text"
// @Param		count	formData	string	true	"count"
// @Security	ApiKeyAuth
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router	/quiz/generate [post]
func (qc *QuizController) Generate(c *gin.Context) {
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

	text := c.PostForm("text")
	count := c.PostForm("count")

	content := fmt.Sprintf(`Help me to construct basic %v questions with 4 variants and at the end of each question show its answers for the next text:
	%v. The result must be only in next array json form , no other texts are allowed
			[
                {
                        "id": 1,
                        "title": "How old is Nurzat?",
                        "variants": [
                                {"title": "17"},
                                {"title": "19"},
                                {"title": "21"},
                                {"title": "23"}
                        ],
                        "correctAnswer": "19"
                }
			]`, count, text)

	fmt.Println(content)
	client := openai.NewClient("sk-proj-Yd9kX98zOyhr4TcXf9nHT3BlbkFJFxrZ2aR4QVBRrwBuagk0")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
	response := []models.Question{}

	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &response)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

	c.JSON(http.StatusOK, models.SuccessResponse{Result: response})
}
