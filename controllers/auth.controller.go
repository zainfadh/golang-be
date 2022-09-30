package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	dto "golang-be/models/dto"
	"golang-be/services"
	"golang-be/utils/constants"
	"golang-be/utils/response"
)

type (
	authController struct {
		UserService services.UserService
	}

	AuthController interface {
		Login(c *gin.Context)
	}
)

// NewAuthController ...
func NewAuthController(db *gorm.DB) *authController {
	return &authController{
		UserService: services.NewUserService(db),
	}
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param login body dto.LoginRequestDto true "login info"
// @Success 200 {object} models.Response
// @Router /v9/auth/login [post]
func (a *authController) Login(c *gin.Context) {
	req := dto.LoginRequestDto{}
	res := response.ResponseApi{}

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

	c.JSON(http.StatusOK, a.UserService.AuthLogin(&req))
}
