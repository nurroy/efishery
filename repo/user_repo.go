package repo

import (
	"belajar/efishery/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

// UserStruct ...
type UserStruct struct {
	db *gorm.DB
}

// UserRepoInterface ...
type UserRepoInterface interface {
	RegisterUser(data *models.User) (*models.User, error)
	GetUserById(userid string) (*models.User, error)
}

// CreateRegisterRepoImpl ...
func CreateUserRepoImpl(db *gorm.DB) UserRepoInterface {
	return &UserStruct{db}
}

// RegisterUser ...
func (r *UserStruct) RegisterUser(data *models.User) (*models.User, error) {
	tx := r.db.Begin()

	if err := tx.Debug().Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[UserRepoImpl.Insert] Error when query save data with : %w", err)
	}

	tx.Commit()

	return data, nil
}


// CreateAuth ...
func (r *UserStruct) GetUserById(userid string) (*models.User, error) {
	data := models.User{}
	err := r.db.Debug().Where("user_id = ?", userid).Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("UserRepoImpl.GetTabunganByID] Error when query get by id with error: %w", err)
	}
	return &data, nil
}

