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
	poDetailController struct {
		PoDetailService services.PoDetailService
	}

	PoDetailController interface {
		SavePoDetail(c *gin.Context)
		UpdatePoDetail(c *gin.Context)
		GePoDetailByID(c *gin.Context)
		Delete(c *gin.Context)
	}
)

// NewPoDetailController ...
func NewPoDetailController(db *gorm.DB) *poDetailController {
	return &poDetailController{
		PoDetailService: services.NewPoDetailService(db),
	}
}

// SavePoDetail ...
func (u *poDetailController) SavePoDetail(c *gin.Context) {

	req := models.PurchaseOrderDetail{}
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

	c.JSON(http.StatusOK, u.PoDetailService.SavePoDetail(&req))
	return
}

// UpdatePoDetail ...
func (u *poDetailController) UpdatePoDetail(c *gin.Context) {
	req := models.PurchaseOrderDetail{}
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

	c.JSON(http.StatusOK, u.PoDetailService.UpdatePoDetail(&req))
	return
}

// GePoDetailByID ...
func (u *poDetailController) GePoDetailByID(c *gin.Context) {
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

	res = u.PoDetailService.GetPoDetailByID(req.ID)

	c.JSON(http.StatusOK, res)
	c.Abort()
}

func (u *poDetailController) Delete(c *gin.Context) {
	req := models.PurchaseOrderDetail{}
	res := response.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request")
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = u.PoDetailService.Delete(req.ID)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}
