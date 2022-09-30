package services

import (
	"time"

	"gorm.io/gorm"

	"golang-be/models"
	dto "golang-be/models/dto"
	"golang-be/repository"
	"golang-be/utils/constants"
	"golang-be/utils/response"
)

type (
	userService struct {
		Repository repository.UserRepository
	}

	UserService interface {
		AuthLogin(userDto *dto.LoginRequestDto) response.ResponseApi
		SaveUser(user *models.User) response.Response
		UpdateUser(user *models.User) response.Response
		GetUserByID(ID int64) response.ResponseApi
		Delete(ID int64) response.Response
	}
)

// NewUserService ...
func NewUserService(db *gorm.DB) *userService {
	return &userService{
		Repository: repository.NewUserRepository(db),
	}
}

// AuthLogin ...
func (a *userService) AuthLogin(userDto *dto.LoginRequestDto) response.ResponseApi {
	var res response.ResponseApi

	if userDto.Username == "" {
		res.ResponseCode = constants.ERROR_RC_50
		res.ResponseDesc = constants.ERROR_RM_50
		return res
	}

	if userDto.Password == "" {
		res.ResponseCode = constants.ERROR_RC_50
		res.ResponseDesc = constants.ERROR_RM_50
		return res
	}

	user, err := a.Repository.GetUserByName(userDto.Username)
	if err != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
		return res
	}

	if user.ID == 0 {
		res.ResponseCode = constants.ERROR_RC_50
		res.ResponseDesc = constants.ERROR_RM_50
		return res
	}

	if user.Password != userDto.Password {
		res.ResponseCode = constants.ERROR_RC_50
		res.ResponseDesc = constants.ERROR_RM_50
		return res
	}

	res.ResponseCode = constants.ERROR_RC_200
	res.ResponseDesc = constants.ERROR_RM_200

	return res
}

// SaveUser ...
func (u *userService) SaveUser(user *models.User) response.Response {
	user.LastUpdate = time.Now()
	// brand.LastUpdateBy = dto.CurrUser

	res := u.Repository.SaveUser(*user)

	return res
}

// UpdateUser ...
func (u *userService) UpdateUser(user *models.User) response.Response {
	user.LastUpdate = time.Now()
	// user.LastUpdateBy = dto.CurrUser

	res := u.Repository.UpdateUser(*user)

	return res
}

// GetUserByID ...
func (u *userService) GetUserByID(ID int64) response.ResponseApi {
	var res response.ResponseApi

	data, err := u.Repository.List(models.User{
		ID: ID,
	})

	if err != nil {
		res.Data = nil
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
		return res
	}

	res.Data = data
	res.ResponseCode = constants.ERROR_RC_200
	res.ResponseDesc = constants.ERROR_RM_200

	return res
}

// Delete ..
func (u *userService) Delete(ID int64) response.Response {
	var res response.Response

	res = u.Repository.Delete(ID)

	return res
}
