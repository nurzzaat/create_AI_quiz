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

//	@Tags		Quiz
//	@Param		text	formData	string	true	"text"
//	@Param		count	formData	string	true	"count"
//	@Security	ApiKeyAuth
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/quiz/generate [post]
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

	// content := fmt.Sprintf(`Help me to construct basic %v questions with 4 variants and at the end of each question show its answers for the next text:
	// %v. The result must be only in next array json form , no other texts are allowed
	// 		[
    //             {
    //                     "id": 1,
    //                     "title": "How old is Nurzat?",
    //                     "variants": [
    //                             {"title": "17"},
    //                             {"title": "19"},
    //                             {"title": "21"},
    //                             {"title": "23"}
    //                     ],
    //                     "correctAnswer": "19"
    //             }
	// 		]`, count, text)

	content := fmt.Sprintf(`Помоги мне создать %v вопросов с 4 вариантами для следующего текста %v и в конце каждого вопроса должен быть правильный ответ на него 
	Ответ должен быть как в следующем формате , позволено только массив json-ов , другие ответы не принимается.
			[
                {
                        "id": 1,
                        "title": "Сколько лет Земле?",
                        "variants": [
                                {"title": "17 млн"},
                                {"title": "1 млрд"},
                                {"title": "21 млрд"},
                                {"title": "4.54 млрд"}
                        ],
                        "correctAnswer": "4.54 млрд"
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
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_RESPONSE",
					Message: "Couldn't bind data from openAI. Please , regenerate it again to get correct values.",
				},
			},
		})
		return
    }

	c.JSON(http.StatusOK, models.SuccessResponse{Result: response})
}
