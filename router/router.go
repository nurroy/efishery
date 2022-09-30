package router

import (
	"belajar/efishery/configs"
	"belajar/efishery/controller/user"
	"belajar/efishery/models"
	"belajar/efishery/repo"
	"belajar/efishery/usecase"
	"github.com/gin-gonic/gin"
)

func Router(config *configs.Config) *gin.Engine{
	router := gin.New()

	db := configs.ConnectDB()
	models.InitTable(db)

	userRepo := repo.CreateUserRepoImpl(db)
	userService := usecase.CreateUserServiceImpl(userRepo,config)


	v1 := router.Group("api/v1")

	{
		newRoute := v1.Group("efishery")
		user.CreateUserController(newRoute, userService)
	}
	return router

}
