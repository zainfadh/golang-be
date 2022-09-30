package services

import (
	"time"

	"gorm.io/gorm"

	"golang-be/models"
	"golang-be/repository"
	"golang-be/utils/constants"
	"golang-be/utils/response"
)

type (
	itemService struct {
		Repository repository.ItemRepository
	}

	ItemService interface {
		SaveItem(item *models.Item) response.Response
		UpdateItem(item *models.Item) response.Response
		GetItemByID(ID int64) response.ResponseApi
		Delete(ID int64) response.Response
	}
)

// NewItemService ...
func NewItemService(db *gorm.DB) *itemService {
	return &itemService{
		Repository: repository.NewItemRepository(db),
	}
}

// SaveItem ...
func (u *itemService) SaveItem(item *models.Item) response.Response {
	item.LastUpdate = time.Now()
	// brand.LastUpdateBy = dto.CurrItem

	res := u.Repository.SaveItem(*item)

	return res
}

// UpdateItem ...
func (u *itemService) UpdateItem(item *models.Item) response.Response {
	item.LastUpdate = time.Now()
	// item.LastUpdateBy = dto.CurrItem

	res := u.Repository.UpdateItem(*item)

	return res
}

// GetItemByID ...
func (u *itemService) GetItemByID(ID int64) response.ResponseApi {
	var res response.ResponseApi

	data, err := u.Repository.List(models.Item{
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
func (u *itemService) Delete(ID int64) response.Response {
	var res response.Response

	res = u.Repository.Delete(ID)

	return res
}
