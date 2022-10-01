package router

import (
	"belajar/efishery/configs"
	"belajar/efishery/controller/fetch"
	"belajar/efishery/controller/user"
	"belajar/efishery/models"
	"belajar/efishery/repo"
	"belajar/efishery/usecase"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/jellydator/ttlcache/v2"
)




func Router(config *configs.Config) *gin.Engine{
	router := gin.New()

	var cache ttlcache.SimpleCache = ttlcache.NewCache()
	cache.SetTTL(time.Duration(24 * time.Hour))

	db := configs.ConnectDB()
	models.InitTable(db)

	userRepo := repo.CreateUserRepoImpl(db)
	userService := usecase.CreateUserServiceImpl(userRepo,config)

	fetchService := usecase.CreateFetchServiceImpl(config,cache)


	v1 := router.Group("api/v1")

	{
		newRoute := v1.Group("efishery")
		user.CreateUserController(newRoute, userService)
		fetch.CreateFetchController(newRoute,fetchService,config)
	}
	return router

}
