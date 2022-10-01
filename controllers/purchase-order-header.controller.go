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
	poHeaderController struct {
		PoHeaderService services.PoHeaderService
	}

	PoHeaderController interface {
		SavePoHeader(c *gin.Context)
		UpdatePoHeader(c *gin.Context)
		GePoHeaderByID(c *gin.Context)
		Delete(c *gin.Context)
	}
)

// NewPoHeaderController ...
func NewPoHeaderController(db *gorm.DB) *poHeaderController {
	return &poHeaderController{
		PoHeaderService: services.NewPoHeaderService(db),
	}
}

// SavePoHeader ...
func (u *poHeaderController) SavePoHeader(c *gin.Context) {

	req := models.PurchaseOrderHeader{}
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

	c.JSON(http.StatusOK, u.PoHeaderService.SavePoHeader(&req))
	return
}

// UpdatePoHeader ...
func (u *poHeaderController) UpdatePoHeader(c *gin.Context) {
	req := models.PurchaseOrderHeader{}
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

	c.JSON(http.StatusOK, u.PoHeaderService.UpdatePoHeader(&req))
	return
}

// GePoHeaderByID ...
func (u *poHeaderController) GePoHeaderByID(c *gin.Context) {
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

	res = u.PoHeaderService.GetPoHeaderByID(req.ID)

	c.JSON(http.StatusOK, res)
	c.Abort()
}

func (u *poHeaderController) Delete(c *gin.Context) {
	req := models.PurchaseOrderHeader{}
	res := response.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request")
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = u.PoHeaderService.Delete(req.ID)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}
