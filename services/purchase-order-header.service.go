package services

import (
	"gorm.io/gorm"

	"golang-be/models"
	"golang-be/repository"
	"golang-be/utils/constants"
	"golang-be/utils/response"
)

type (
	poHeaderService struct {
		Repository repository.PoHeaderRepository
	}

	PoHeaderService interface {
		SavePoHeader(poHeader *models.PurchaseOrderHeader) response.Response
		UpdatePoHeader(poHeader *models.PurchaseOrderHeader) response.Response
		GetPoHeaderByID(ID int) response.ResponseApi
		Delete(ID int) response.Response
	}
)

// NewPoHeaderService ...
func NewPoHeaderService(db *gorm.DB) *poHeaderService {
	return &poHeaderService{
		Repository: repository.NewPoHeaderRepository(db),
	}
}

// SavePoHeader ...
func (u *poHeaderService) SavePoHeader(poHeader *models.PurchaseOrderHeader) response.Response {

	res := response.Response{}

	_, err := u.Repository.SavePoHeader(*poHeader)

	if err != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
		return res
	}

	res.ResponseCode = constants.ERROR_RC_200
	res.ResponseDesc = constants.ERROR_RM_200

	return res
}

// UpdatePoHeader ...
func (u *poHeaderService) UpdatePoHeader(poHeader *models.PurchaseOrderHeader) response.Response {

	res := u.Repository.UpdatePoHeader(*poHeader)

	return res
}

// GetPoHeaderByID ...
func (u *poHeaderService) GetPoHeaderByID(ID int) response.ResponseApi {
	var res response.ResponseApi

	data, err := u.Repository.List(models.PurchaseOrderHeader{
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
func (u *poHeaderService) Delete(ID int) response.Response {
	var res response.Response

	res = u.Repository.Delete(ID)

	return res
}
