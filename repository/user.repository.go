package repository

import (
	"fmt"
	"golang-be/models"
	"golang-be/utils/constants"
	"golang-be/utils/response"
	"time"

	"github.com/astaxie/beego/logs"
	"gorm.io/gorm"
)

type (
	userRepository struct {
		DB *gorm.DB
	}

	UserRepository interface {
		GetUserByName(username string) (models.User, error)
		SaveUser(user models.User) response.Response
		UpdateUser(user models.User) response.Response
		List(data models.User) ([]models.User, error)
		Delete(ID int64) response.Response
	}
)

// NewUserRepository ..
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

// GetUserByName ...
func (u *userRepository) GetUserByName(username string) (models.User, error) {

	var user models.User
	var err error

	u.DB.Where("user_name = ?", username).Find(&user)

	fmt.Println("User => ", user)
	return user, err

}

// SaveUser ...
func (u *userRepository) SaveUser(user models.User) response.Response {
	var res response.Response

	if r := u.DB.Save(&user); r.Error != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
	}

	res.ResponseCode = constants.ERROR_RC_200
	res.ResponseDesc = constants.ERROR_RM_200
	return res
}

// Update ...
func (u *userRepository) UpdateUser(user models.User) response.Response {
	var res response.Response
	var users models.User

	err := u.DB.Model(&models.User{}).Where("id = ? ", user.ID).First(&users).Error
	if err != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
	}

	users.Email = user.Email
	users.Password = user.Password
	user.LastUpdate = time.Now()
	user.LastLogin = time.Now()
	user.CreatedAt = time.Now()

	res.ResponseCode = constants.ERROR_RC_200
	res.ResponseDesc = constants.ERROR_RM_200
	return res
}

// List ..
func (u *userRepository) List(data models.User) ([]models.User, error) {
	user := []models.User{}
	if err := u.DB.Where("id = ?", &data.ID).Find(&user).Order("id ASC").Error; err != nil {
		logs.Error("Error List MitraSupplierPriceWholesaler", err)
		return user, err
	}
	return user, nil
}

// Delete ..
func (u *userRepository) Delete(ID int64) response.Response {
	res := response.Response{}
	data := models.User{}

	if err := u.DB.Where("id = ?", ID).Delete(data).Error; err != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
	}

	if ID == 0 {
		res.ResponseCode = constants.ERROR_RC_04
		res.ResponseDesc = constants.ERROR_RM_04
		return res
	}

	res.ResponseCode = constants.ERROR_RC_200
	res.ResponseDesc = constants.ERROR_RM_200

	return res
}
