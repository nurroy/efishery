package user

import (
	"belajar/efishery/models"
	"belajar/efishery/usecase"
	"belajar/efishery/utils/generator"
	"belajar/efishery/utils/hash"
	utils "belajar/efishery/utils/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Methode is the methode type
type Method int

const (
	// List of different Methods
	REGISTER Method = iota
	LOGIN
	VALIDATE
	other
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

// isValidRequest validates the status request
func isValidRequest(m Method, req *models.User) error {
	switch m {
	case REGISTER:
		if req.UserID == "" {
			return errors.New("missing userid")
		}
		if req.Username == "" {
			return errors.New("missing username")
		}
		if req.Role == "" {
			return errors.New("missing role")
		}
	case LOGIN:
		if req.UserID == ""{
			return errors.New("missing userid")
		}
		if req.Password == ""{
			return errors.New("missing password")
		}
	case VALIDATE:
		if req.Token == ""{
			return errors.New("missing token")
		}
	}

	return nil
}


//UserController
func (r *UserController) UserController(c *gin.Context){
	var user *models.User
	var res models.RegisterReponse

	err := c.ShouldBindJSON(&user)
	if err != nil{
		utils.SuccessMessage(c, http.StatusInternalServerError, "Oppss, something went error")
		log.Printf("[CONTROLLER][USER][USER][1001] Error when decoder data from body with error : %v\n", err)
		return
	}

	//check if request is valid
	err = isValidRequest(REGISTER,user)
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, err.Error(),1002)
		log.Printf("[CONTROLLER][USER][USER][1002] Error when validate request with error: %v\n", err)
		return
	}

	data,err := r.userService.CheckData(user.UserID)
	if err != nil && data != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something went error\"",1003)
		log.Printf("[CONTROLLER][USER][USER][1003] Error when check data with error: %v\n", err)
		return
	}

	if data != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "data already exists found\"",1004)
		log.Println("[CONTROLLER][USER][USER][1004] no data found")
		return
	}

	//generate random char
	pass := generator.GenerateRandomChar(4)
	//Hashing key
	encryptionKey, err := hash.HashAndSalt([]byte(pass))
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something went error\"",1005)
		log.Printf("[CONTROLLER][USER][USER][1005] Error when data hashing with error: %v\n", err)
		return
	}

	//set user password
	user.Password = encryptionKey

	//register user
	_,err = r.userService.RegisterUser(user)
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something went error\"",1006)
		log.Printf("[CONTROLLER][USER][USER][1006] Error when register user with error: %v\n", err)
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
		log.Printf("[CONTROLLER][USER][LOGIN][2001] Error when decoder data from body with error : %v\n", err)
		return
	}

	//check if request is valid
	err = isValidRequest(LOGIN,user)
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, err.Error(),2002)
		log.Printf("[CONTROLLER][USER][LOGIN][2002] Error when validate request with error: %v\n", err)
		return
	}

	//check data
	data ,err:= r.userService.CheckData(user.UserID)
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "Data not found",2003)
		log.Printf("[CONTROLLER][USER][LOGIN][2003] Error when check data with error: %v\n", err)
		return
	}

	// check if data exists
	if data == nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "username has been registered",2004)
		log.Println("[CONTROLLER][USER][LOGIN][2004] no data found")
		return
	}

	err = hash.ComparePasswords(data.Password, []byte(user.Password))
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "invalid password",2005)
		log.Printf("[CONTROLLER][USER][LOGIN][2005] Error when compare password with error: %v\n", err)
		return
	}

	res, err := r.userService.CreateAuth(user)
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something went error\"",2006)
		log.Printf("[CONTROLLER][USER][LOGIN][2006] Error when create auth with error: %v\n", err)
		return
	}


	utils.SuccessData(c, http.StatusOK, res)
}


//ValidateController
func (r *UserController) ValidateController(c *gin.Context){
	var user *models.User

	err := c.ShouldBindJSON(&user)
	if err != nil{
		utils.ErrorMessage(c, http.StatusInternalServerError, "Oppss, something went error\"",3001)
		log.Printf("[CONTROLLER][USER][VALIDATE][3001] Error when check data with error: %v\n", err)
		return
	}

	//check if request is valid
	err = isValidRequest(VALIDATE,user)
	if err != nil{
		utils.ErrorMessage(c, http.StatusBadRequest, err.Error(),3002)
		log.Printf("[CONTROLLER][USER][VALIDATE][3002] Error when validate request with error: %v\n", err)
		return
	}

	//validate token
	res,err := r.userService.ValidateToken(user)
	if err != nil{
		utils.ErrorMessage(c, http.StatusInternalServerError, fmt.Sprintf("%v",err),3003)
		log.Printf("[CONTROLLER][USER][VALIDATE][3003] Error when validate token with error: %v\n", err)
		return
	}
	utils.SuccessData(c, http.StatusOK, res)
}


