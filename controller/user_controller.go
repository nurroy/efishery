package controller

import (
	"belajar/efishery/models"
	"belajar/efishery/usecase"
	"belajar/efishery/utils/generator"
	"belajar/efishery/utils/hash"
	utils "belajar/efishery/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
	userService usecase.UserServiceInterface
}

//CreateUserController
func CreateUserController(router *gin.RouterGroup, userService usecase.UserServiceInterface){
	inDB := UserController{userService: userService}

	router.POST("/user", inDB.UserController)
	router.POST("/login", inDB.LoginController)
	router.POST("/validate", inDB.ValidateController)
}

//UserController
func (r *UserController) UserController(c *gin.Context){
	var user *models.User
	var res models.RegisterReponse

	err := c.ShouldBindJSON(&user)
	if err != nil{
		utils.SuccessMessage(c, http.StatusInternalServerError, "Oppss, something went error")
		log.Printf("[Controller.User.User][1001] Error when decoder data from body with error : %v\n", err)
		return
	}

	data,err := r.userService.CheckData(user.UserID)
	if err != nil && data != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something went error\"",1002)
		log.Printf("[Controller.User.User][1002] Error when check data with error: %v\n", err)
		return
	}

	if data != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "data already exists found\"",1003)
		log.Println("[Controller.User.User][1003] no data found")
		return
	}

	//generate random char
	pass := generator.GenerateRandomChar(4)
	//Hashing key
	encryptionKey, err := hash.HashAndSalt([]byte(pass))
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something went error\"",1004)
		log.Printf("[Controller.User.User][1004] Error when check data with error: %v\n", err)
		return
	}

	//set user password
	user.Password = encryptionKey

	//register user
	_,err = r.userService.RegisterUser(user)
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something went error\"",1005)
		log.Printf("[Controller.User.User][1005] Error when check data with error: %v\n", err)
		return
	}

	res.Password = pass

	utils.SuccessData(c, http.StatusOK, res)
}

//UserController
func (r *UserController) LoginController(c *gin.Context){
	var user *models.User

	err := c.ShouldBindJSON(&user)
	if err != nil{
		utils.SuccessMessage(c, http.StatusInternalServerError, "Oppss, something went error")
		log.Printf("[Controller.User.Login][2001] Error when decoder data from body with error : %v\n", err)
		return
	}

	data ,err:= r.userService.CheckData(user.UserID)
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "Data not found",2002)
		log.Printf("[Controller.User.Login][2002] Error when check data with error: %v\n", err)
		return
	}

	if data == nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "username has been registered",2003)
		log.Println("[Controller.User.Login][2003] no data found")
		return
	}

	err = hash.ComparePasswords(data.Password, []byte(user.Password))
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "invalid password",2004)
		log.Printf("[Controller.User.Login][2004] Error when compare password with error: %v\n", err)
		return
	}

	res, err := r.userService.CreateAuth(user)
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something went error\"",2005)
		log.Printf("[Controller.User.Login][2005] Error when check data with error: %v\n", err)
		return
	}


	utils.SuccessData(c, http.StatusOK, res)
}


//ValidateController
func (r *UserController) ValidateController(c *gin.Context){
	var user *models.User

	err := c.ShouldBindJSON(&user)
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something went error\"",3001)
		log.Printf("[Controller.User.Validate][3001] Error when check data with error: %v\n", err)
		return
	}

	res,err := r.userService.ValidateToken(user)
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, fmt.Sprintf("%v",err),3002)
		log.Printf("[Controller.User.Validate][3002] Error when check data with error: %v\n", err)
		return
	}
	utils.SuccessData(c, http.StatusOK, res)
}


