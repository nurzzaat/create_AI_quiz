package quiz

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	openai "github.com/sashabaranov/go-openai"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

type QuizController struct {
	QuizRepository models.QuizRepository
}

func (qc *QuizController) Create(c *gin.Context) {
	questionCount := 2
	client := openai.NewClient("sk-proj-Yd9kX98zOyhr4TcXf9nHT3BlbkFJFxrZ2aR4QVBRrwBuagk0")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleUser,
					Content: `i'm creating quiz for students , construct me basic ` + strconv.Itoa(questionCount) + ` questions in the text language(if text in russian , then questions also must be in russia) with 4 variants and at the end of question show their answers.` +
						`Бұрынғы өткен заманда бір топ адам кеме
мен алыс сапарға шықпақ болыпты. Баратын 
жерлері алыс болғанына қарамастан, айлап 
жол жүруге бекініпті. Саяхаттамақшы болған 
адамдардың көп болғанына қарай, кеме де 
соған сай үлкен екен. Екі қабаттан тұратын 
бұл кеменің астыңғы қабаты су деңгейімен 
шамалас немесе одан сәл төмендеу болады. 
Ал жоғарғы қабаты су деңгейінен жоғары  
тұрады. Саяхатшылар орынға таласып қалмау 
үшін өзара жеребе тастап, кемегедегі орында
рына жайғасады. Кейбіріне кеменің жоғарғы 
қабаты, кейбіріне оның төменгі қабаты бұй
ырады. Астыңғы қабатта тұратындар су алу 
үшін жоғарғы қабатқа шығуы керек болады.`,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
	c.JSON(http.StatusOK, models.SuccessResponse{Result: "success"})
}
