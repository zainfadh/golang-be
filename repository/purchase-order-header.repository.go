package repository

import (
	"golang-be/models"
	"golang-be/utils/constants"
	"golang-be/utils/response"
	"time"

	"github.com/astaxie/beego/logs"
	"gorm.io/gorm"
)

type (
	poHeaderRepository struct {
		DB *gorm.DB
	}

	PoHeaderRepository interface {
		SavePoHeader(poHeader models.PurchaseOrderHeader) response.Response
		UpdatePoHeader(poHeader models.PurchaseOrderHeader) response.Response
		List(data models.PurchaseOrderHeader) ([]models.PurchaseOrderHeader, error)
		Delete(ID int) response.Response
	}
)

// NewPoHeaderRepository ..
func NewPoHeaderRepository(db *gorm.DB) *poHeaderRepository {
	return &poHeaderRepository{
		DB: db,
	}
}

// SavePoHeader ...
func (u *poHeaderRepository) SavePoHeader(poHeader models.PurchaseOrderHeader) response.Response {
	db := u.DB.Debug()
	var res response.Response

	if r := db.Save(&poHeader); r.Error != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
	}

	res.ResponseCode = constants.ERROR_RC_200
	res.ResponseDesc = constants.ERROR_RM_200
	return res
}

// Update ...
func (u *poHeaderRepository) UpdatePoHeader(poHeader models.PurchaseOrderHeader) response.Response {
	db := u.DB.Debug()
	var res response.Response
	var poHeaders models.PurchaseOrderHeader

	err := db.Model(&models.PurchaseOrderHeader{}).Where("id = ? ", poHeader.ID).First(&poHeaders).Error
	if err != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
	}

	poHeaders.Description = poHeader.Description
	poHeaders.Cost = poHeader.Cost
	poHeaders.Price = poHeader.Cost
	poHeader.Date = time.Now()
	poHeader.CreatedAt = time.Now()

	res.ResponseCode = constants.ERROR_RC_200
	res.ResponseDesc = constants.ERROR_RM_200
	return res
}

// List ..
func (u *poHeaderRepository) List(data models.PurchaseOrderHeader) ([]models.PurchaseOrderHeader, error) {
	db := u.DB.Debug()
	poHeader := []models.PurchaseOrderHeader{}
	if err := db.Where("id = ?", &data.ID).Find(&poHeader).Order("id ASC").Error; err != nil {
		logs.Error("Error List MitraSupplierPriceWholesaler", err)
		return poHeader, err
	}
	return poHeader, nil
}

// Delete ..
func (u *poHeaderRepository) Delete(ID int) response.Response {
	db := u.DB.Debug()
	res := response.Response{}
	data := models.PurchaseOrderHeader{}

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
