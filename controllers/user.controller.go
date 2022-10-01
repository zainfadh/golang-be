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
	userController struct {
		UserService services.UserService
	}

	UserController interface {
		SaveUser(c *gin.Context)
		UpdateUser(c *gin.Context)
		GeUserByID(c *gin.Context)
		Delete(c *gin.Context)
	}
)

// NewUserController ...
func NewUserController(db *gorm.DB) *userController {
	return &userController{
		UserService: services.NewUserService(db),
	}
}

// SaveUser ...
func (u *userController) SaveUser(c *gin.Context) {

	req := models.User{}
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

	c.JSON(http.StatusOK, u.UserService.SaveUser(&req))
	return
}

// UpdateUser ...
func (u *userController) UpdateUser(c *gin.Context) {
	req := models.User{}
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

	c.JSON(http.StatusOK, u.UserService.UpdateUser(&req))
	return
}

// GeUserByID ...
func (u *userController) GeUserByID(c *gin.Context) {
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

	res = u.UserService.GetUserByID(req.ID)

	c.JSON(http.StatusOK, res)
	c.Abort()
}

func (u *userController) Delete(c *gin.Context) {
	req := models.User{}
	res := response.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request")
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	res = u.UserService.Delete(req.ID)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}
