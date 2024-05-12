package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/controller"
	"github.com/nurzzaat/create_AI_quiz/pkg"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.

//	@host		77.240.39.57
//	@BasePath	/ai

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	app, err := pkg.App()

	if err != nil {
		log.Fatal(err)
	}
	//defer app.CloseDBConnection()

	ginRouter := gin.Default()
	ginRouter.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS ,HEAD")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-*, Cross-Origin-Resource-Policy , Origin, X-Requested-With, Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})
	controller.Setup(app, ginRouter)

	ginRouter.Run(fmt.Sprintf(":%s", app.Env.PORT))
}
