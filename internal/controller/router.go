package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/middleware"
	"github.com/nurzzaat/create_AI_quiz/pkg"

	_ "github.com/nurzzaat/create_AI_quiz/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/nurzzaat/create_AI_quiz/internal/controller/auth"
	"github.com/nurzzaat/create_AI_quiz/internal/controller/quiz"
	"github.com/nurzzaat/create_AI_quiz/internal/controller/user"
	"github.com/nurzzaat/create_AI_quiz/internal/repository"
)

func Setup(app pkg.Application, router *gin.Engine) {
	env := app.Env
	db := app.Pql

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.Static("/images", "./images")

	loginController := &auth.AuthController{
		UserRepository: repository.NewUserRepository(db),
		Env:            env,
	}

	userController := &user.UserController{
		UserRepository: repository.NewUserRepository(db),
	}

	quizController := &quiz.QuizController{
		QuizRepository: repository.NewQuizRepository(db),
	}

	router.POST("/signup", loginController.Signup)
	router.POST("/signin", loginController.Signin)
	router.POST("/forgot-password", loginController.ForgotPassword)
	router.POST("/students/quiz/add", quizController.AddStudentToQuizPOST)
	router.PUT("/students/quiz/:quizId/add/:studentId", quizController.AddStudentToQuizPUT)

	router.Use(middleware.JWTAuth(env.AccessTokenSecret))
	router.POST("/logout", loginController.Logout)

	userRouter := router.Group("/user")
	{
		userRouter.GET("/profile", userController.GetProfile)
		userRouter.PUT("/profile", userController.UpdateProfile)
		userRouter.POST("/reset-password", loginController.ResetPassword)
	}

	quizRouter := router.Group("/quiz")
	{
		quizRouter.POST("", quizController.Create)
		quizRouter.POST("/generate", quizController.Generate)
		quizRouter.GET("/admin/:id", quizController.GetByIDAdmin)
		quizRouter.GET("/user/:id", quizController.GetByIDUser)
		quizRouter.GET("/admin", quizController.GetAllAdmin)
		quizRouter.GET("/user", quizController.GetAllUser)
		quizRouter.DELETE("/:quizId", quizController.Delete)

		quizRouter.POST("/submit/:quizId", quizController.Submit)
	}
	studentRouter := router.Group("/students")
	{
		studentRouter.GET("", userController.GetAll)
		studentRouter.GET("/:quizId/result", quizController.GetStudentsByQuizID)
		studentRouter.GET("/:quizId", quizController.GetPermittedStudentsByQuizID)
		studentRouter.DELETE("/quiz/:quizId/delete/:studentId", quizController.DeleteStudentFromQuiz)
		studentRouter.GET("/quiz/:quizId/result", quizController.GetStudentResult)
	}
}
