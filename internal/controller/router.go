package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/ZharasDiplom/middleware"
	"github.com/nurzzaat/ZharasDiplom/pkg"
	
	_ "github.com/nurzzaat/ZharasDiplom/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	
	"github.com/nurzzaat/ZharasDiplom/internal/controller/auth"
	"github.com/nurzzaat/ZharasDiplom/internal/controller/user"
	"github.com/nurzzaat/ZharasDiplom/internal/repository"
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

	router.POST("/signup", loginController.Signup)
	router.POST("/signin" , loginController.Signin)
	router.POST("/forgot-password" , loginController.ForgotPassword)
	//router.POST("/accountrecovery/:email/:hash_pass" , loginController.AccountRecovery)
	//router.GET("/accountrecovery/:email/:hash_pass" , loginController.AccountRecoveryHTML)
	
	router.Use(middleware.JWTAuth(env.AccessTokenSecret))
	router.POST("/logout" , loginController.Logout)

	userRouter := router.Group("/user")
	{
		userRouter.GET("/profile" , userController.GetProfile)
		userRouter.POST("/reset-password" , loginController.ResetPassword)
	}
	
	
}
