package services

import (
	"gorm.io/gorm"

	"golang-be/models"
	"golang-be/repository"
	"golang-be/utils/constants"
	"golang-be/utils/response"
)

type (
	poDetailService struct {
		Repository         repository.PoDetailRepository
		PoHeaderRepository repository.PoHeaderRepository
		ItemRepository     repository.ItemRepository
		DB                 *gorm.DB
	}

	PoDetailService interface {
		SavePoDetail(poDetail *models.PurchaseOrderDetail) response.Response
		UpdatePoDetail(poDetail *models.PurchaseOrderDetail) response.Response
		GetPoDetailByID(ID int) response.ResponseApi
		Delete(ID int) response.Response
	}
)

// NewPoDetailService ...
func NewPoDetailService(db *gorm.DB) *poDetailService {
	return &poDetailService{
		Repository:         repository.NewPoDetailRepository(db),
		PoHeaderRepository: repository.NewPoHeaderRepository(db),
		ItemRepository:     repository.NewItemRepository(db),
		DB:                 db,
	}
}

// SavePoDetail ...
func (u *poDetailService) SavePoDetail(poDetail *models.PurchaseOrderDetail) response.Response {
	var poHeader models.PurchaseOrderHeader
	var res response.Response

	item, _ := u.ItemRepository.GetItemById(poDetail.ItemId)

	poDetail.ItemId = item.ID

	// db trx
	if err := u.DB.Transaction(func(tx *gorm.DB) error {

		// save header >> balikin id yg udah di save & error
		id, err := u.PoHeaderRepository.SavePoHeader(poHeader)
		// kondisi error
		if err != nil {
			return err
		}

		// save detail >> error
		poDetail.PurchaseOrderId = id
		err = u.Repository.SavePoDetail(*poDetail)
		// kondisi error
		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
		return res
	}

	return res
}

// UpdatePoDetail ...
func (u *poDetailService) UpdatePoDetail(poDetail *models.PurchaseOrderDetail) response.Response {

	res := u.Repository.UpdatePoDetail(*poDetail)

	return res
}

// GetPoDetailByID ...
func (u *poDetailService) GetPoDetailByID(ID int) response.ResponseApi {
	var res response.ResponseApi

	data, err := u.Repository.List(models.PurchaseOrderDetail{
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
func (u *poDetailService) Delete(ID int) response.Response {
	var res response.Response

	res = u.Repository.Delete(ID)

	return res
}
