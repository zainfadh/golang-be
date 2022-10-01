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
	itemRepository struct {
		DB *gorm.DB
	}

	ItemRepository interface {
		SaveItem(item models.Item) response.Response
		UpdateItem(item models.Item) response.Response
		List(data models.Item) ([]models.Item, error)
		Delete(ID int) response.Response
		GetItemById(id int) (models.Item, error)
	}
)

// NewItemRepository ..
func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{
		DB: db,
	}
}

// SaveItem ...
func (u *itemRepository) SaveItem(item models.Item) response.Response {
	db := u.DB.Debug()
	var res response.Response

	if r := db.Save(&item); r.Error != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
		return res
	}

	res.ResponseCode = constants.ERROR_RC_200
	res.ResponseDesc = constants.ERROR_RM_200
	return res
}

// Update ...
func (u *itemRepository) UpdateItem(item models.Item) response.Response {
	db := u.DB.Debug()

	var res response.Response
	var items models.Item

	err := db.Model(&models.Item{}).Where("id = ? ", item.ID).First(&items).Error
	if err != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
		return res
	}

	items.Name = item.Name
	items.Description = item.Description
	items.Cost = item.Cost
	items.Price = item.Cost
	item.UpdatedAt = time.Now()
	item.CreatedAt = time.Now()

	res.ResponseCode = constants.ERROR_RC_200
	res.ResponseDesc = constants.ERROR_RM_200
	return res
}

// List ..
func (u *itemRepository) List(data models.Item) ([]models.Item, error) {
	db := u.DB.Debug()
	item := []models.Item{}
	if err := db.Where("id = ?", &data.ID).Find(&item).Order("id ASC").Error; err != nil {
		logs.Error("Error List MitraSupplierPriceWholesaler", err)
		return item, err
	}
	return item, nil
}

// Delete ..
func (u *itemRepository) Delete(ID int) response.Response {
	db := u.DB.Debug()
	res := response.Response{}
	data := models.Item{}

	if err := db.Where("id = ?", ID).Delete(data).Error; err != nil {
		res.ResponseCode = constants.ERROR_RC_511
		res.ResponseDesc = constants.ERROR_RM_511
		return res
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

// GetItemById ..
func (u *itemRepository) GetItemById(id int) (models.Item, error) {
	db := u.DB.Debug()
	item := models.Item{}

	err := db.Where("id=?", id).Find(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}
