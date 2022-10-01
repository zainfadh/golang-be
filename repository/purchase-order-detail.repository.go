package repository

import (
	"golang-be/models"
	"golang-be/utils/constants"
	"golang-be/utils/response"

	"github.com/astaxie/beego/logs"
	"gorm.io/gorm"
)

type (
	poDetailRepository struct {
		DB *gorm.DB
	}

	PoDetailRepository interface {
		SavePoDetail(poDetail models.PurchaseOrderDetail) response.Response
		UpdatePoDetail(poDetail models.PurchaseOrderDetail) response.Response
		List(data models.PurchaseOrderDetail) ([]models.PurchaseOrderDetail, error)
		Delete(ID int) response.Response
	}
)

// NewPoDetailRepository ..
func NewPoDetailRepository(db *gorm.DB) *poDetailRepository {
	return &poDetailRepository{
		DB: db,
	}
}

// SavePoDetail ...
func (u *poDetailRepository) SavePoDetail(poDetail models.PurchaseOrderDetail) response.Response {
	db := u.DB.Debug()
	var res response.Response

	if r := db.Save(&poDetail); r.Error != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
	}

	res.ResponseCode = constants.ERROR_RC_200
	res.ResponseDesc = constants.ERROR_RM_200
	return res
}

// Update ...
func (u *poDetailRepository) UpdatePoDetail(poDetail models.PurchaseOrderDetail) response.Response {
	db := u.DB.Debug()
	var res response.Response
	var poDetails models.PurchaseOrderDetail

	err := db.Model(&models.PurchaseOrderDetail{}).Where("id = ? ", poDetail.ID).First(&poDetails).Error
	if err != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
	}

	res.ResponseCode = constants.ERROR_RC_200
	res.ResponseDesc = constants.ERROR_RM_200
	return res
}

// List ..
func (u *poDetailRepository) List(data models.PurchaseOrderDetail) ([]models.PurchaseOrderDetail, error) {
	db := u.DB.Debug()
	poDetail := []models.PurchaseOrderDetail{}
	if err := db.Where("id = ?", &data.ID).Find(&poDetail).Order("id ASC").Error; err != nil {
		logs.Error("Error List MitraSupplierPriceWholesaler", err)
		return poDetail, err
	}
	return poDetail, nil
}

// Delete ..
func (u *poDetailRepository) Delete(ID int) response.Response {
	db := u.DB.Debug()
	res := response.Response{}
	data := models.PurchaseOrderDetail{}

	if err := db.Where("id = ?", ID).Delete(data).Error; err != nil {
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
