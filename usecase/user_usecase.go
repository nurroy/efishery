package usecase

import (
	"belajar/efishery/auth"
	"belajar/efishery/configs"
	"belajar/efishery/models"
	"belajar/efishery/repo"
	"log"
)

type UserService struct {
	userRepo repo.UserRepoInterface
	config *configs.Config
}

type UserServiceInterface interface {
	RegisterUser(data *models.User)(*models.User,error)
	CreateAuth(data *models.User)(*models.LoginResponse,error)
	ValidateToken(data *models.User)(*models.ValidateResponse,error)
	CheckData(userid string) (*models.User,error)
}

func CreateUserServiceImpl(userRepo repo.UserRepoInterface,config *configs.Config) UserServiceInterface {
	return &UserService{userRepo,config}
}

//RegisterUser
func (r *UserService) RegisterUser(data *models.User)(*models.User,error){

	data,err := r.userRepo.RegisterUser(data)
	if err != nil{
		log.Printf("[Usecase.User.Register][1] error when register user with error : %v",err)
		return nil, err
	}

	return data,nil
}

//CreateAuth
func (r *UserService) CreateAuth(data *models.User)(*models.LoginResponse, error){
	var res models.LoginResponse
	var tokenResult string
	var authD models.Auth

	user , err := r.userRepo.GetUserById(data.UserID)
	if err != nil{
		log.Printf("[Usecase.User.Auth][1] error when get user by id with error : %v",err)
		return nil, err
	}

	authD.UserID = data.UserID
	authD.Username = user.Username
	authD.Role = user.Role

	tokenResult, err = auth.CreateToken(authD,r.config)
	if err != nil {
		log.Printf("[Usecase.User.Auth][2] error when create token with error : %v",err)
		return nil, err
	}

	res.Token = tokenResult

	return &res,nil
}

//ValidateToken
func (r *UserService) ValidateToken(data *models.User)(*models.ValidateResponse,error){
	var res models.ValidateResponse
	claims,err := auth.TokenValid(nil,data.Token,r.config)
	if err != nil{
		log.Printf("[Usecase.User.Validate][1] error when validate user with error : %v",err)
		return nil, err
	}

	res.UserID = claims["userid"].(string)
	res.Username = claims["username"].(string)
	res.Role = claims["role"].(string)

	return &res,nil
}

//CheckData is check data usecase
func (r *UserService) CheckData(userid string)(data *models.User,err error){

	data , err = r.userRepo.GetUserById(userid)
	if err != nil{
		log.Printf("[Usecase.User.Check][1] error when get user by id with error : %v",err)
		return nil, err
	}

	return data,nil
}




