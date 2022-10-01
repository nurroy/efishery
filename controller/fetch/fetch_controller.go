package fetch

import (
	"belajar/efishery/middleware"
	"belajar/efishery/usecase"
	utils "belajar/efishery/utils/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FetchController struct {
	fetchService usecase.FetchServiceInterface
}

//CreateFetchController
func CreateFetchController(router *gin.RouterGroup, fetchService usecase.FetchServiceInterface){
	inDB := FetchController{fetchService: fetchService}

	router.GET("/storages",middleware.TokenAuthMiddleware(), inDB.GetStorages)
	router.GET("/aggregate",middleware.TokenAuthMiddleware(), inDB.GetAggregate)
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


