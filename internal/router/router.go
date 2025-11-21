package router

import (
	"go-crud/internal/handler"
	"go-crud/internal/middleware"
	"go-crud/internal/repository"
	"go-crud/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	r := gin.Default()

	userRepo := repository.UserRepositoryImpl{}
	userSvc := service.UserServiceImpl{Repository: &userRepo}
	userHandler := handler.UserHandler{Service: &userSvc}

	adminRepo := repository.AdminRepositoryImpl{}
	adminSvc := service.AdminServiceImpl{Repository: &adminRepo}
	adminHandler := handler.AdminHandler{Service: &adminSvc}

	r.POST("/signup", adminHandler.SignUp)
	r.POST("/login", adminHandler.LogIn)

	users := r.Group("/users")
	users.Use(middleware.AuthJWT())
	{
		users.POST("/", userHandler.Create)
		users.GET("/", userHandler.GetAll)
		users.GET("/:id", userHandler.GetById)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	return r
}
