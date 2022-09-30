package controllers

import (
	"encoding/json"
	"fmt"
	"golang-be/models"
	dto "golang-be/models/dto"
	"golang-be/services"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"golang-be/utils/constants"
	"golang-be/utils/response"
)

type (
	itemController struct {
		ItemService services.ItemService
	}

	ItemController interface {
		SaveItem(c *gin.Context)
	}
)

// NewItemController ...
func NewItemController(db *gorm.DB) *itemController {
	return &itemController{
		ItemService: services.NewItemService(db),
	}
}

// SaveItem ...
func (u *itemController) SaveItem(c *gin.Context) {

	req := models.Item{}
	res := response.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request ")
		res.ResponseCode = constants.ERROR_RC_03
		res.ResponseDesc = constants.ERROR_RM_03
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, u.ItemService.SaveItem(&req))
	return
}

// UpdateItem ...
func (u *itemController) UpdateItem(c *gin.Context) {
	req := models.Item{}
	res := response.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request ")
		res.ResponseCode = constants.ERROR_RC_03
		res.ResponseDesc = constants.ERROR_RM_03
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, u.ItemService.UpdateItem(&req))
	return
}

// GeItemByID ...
func (u *itemController) GeItemByID(c *gin.Context) {
	res := response.ResponseApi{}
	req := dto.GetDataByID{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request ")
		res.ResponseCode = constants.ERROR_RC_03
		res.ResponseDesc = constants.ERROR_RM_03
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = u.ItemService.GetItemByID(req.ID)

	c.JSON(http.StatusOK, res)
	c.Abort()
}

func (u *itemController) Delete(c *gin.Context) {
	req := models.Item{}
	res := response.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request")
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = u.ItemService.Delete(req.ID)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}
