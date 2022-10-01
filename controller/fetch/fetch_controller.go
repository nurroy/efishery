package fetch

import (
	"belajar/efishery/configs"
	"belajar/efishery/middleware"
	"belajar/efishery/usecase"
	utils "belajar/efishery/utils/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FetchController struct {
	fetchService usecase.FetchServiceInterface
	config *configs.Config
}

//CreateFetchController
func CreateFetchController(router *gin.RouterGroup, fetchService usecase.FetchServiceInterface,config *configs.Config){
	inDB := FetchController{fetchService: fetchService,config: config}

	router.GET("/storages",middleware.TokenAuthMiddleware(config), inDB.GetStorages)
	router.GET("/aggregate",middleware.TokenAuthMiddleware(config), inDB.GetAggregate)
}

func (r *FetchController) GetStorages(c *gin.Context){

	data,err := r.fetchService.GetStorages()
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something error ",1001)
		log.Printf("[CONTROLLER][FETCH][STORAGES][1001] Error when get data with error: %v\n", err)
		return
	}


	utils.SuccessData(c, http.StatusOK, data)
}

func (r *FetchController) GetAggregate(c *gin.Context){
	res,err:= r.fetchService.GetAggregate()
	if err != nil{
		return
	}
	utils.SuccessData(c, http.StatusOK, res)
}


