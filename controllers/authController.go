package controllers

import (
	"go-gin-workshop/constants"
	"go-gin-workshop/dto"
	"go-gin-workshop/services"
	"go-gin-workshop/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	LoginController(ctx *gin.Context)
	RegisterController(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

// ShowAccount godoc
//
//	@Summary		Login User
//	@Description	login user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Request body	dto.LoginDto    true    "Login"
//	@Success		200	{object}	types.Response
//	@Router			/auth/login [post]
func (c *authController) LoginController(ctx *gin.Context) {
	var loginDto dto.LoginDto
	err := ctx.ShouldBind(&loginDto)
	if err != nil {
		res := utils.BuildResponseFailed(http.StatusBadRequest, constants.MESSAGE_MAPDATA_ERR, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.authService.LoginService(ctx.Request.Context(), loginDto)

	if err != nil {
		res := utils.BuildResponseFailed(http.StatusBadRequest, constants.MESSAGE_BAD_REQUEST, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(http.StatusOK, constants.MESSAGE_SUCCESS, result)
	ctx.JSON(http.StatusOK, res)
}

// ShowAccount godoc
//
//	@Summary		Register User
//	@Description	register user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Request body	dto.UserDto    true    "Register"
//	@Success		200	{object}	types.Response
//	@Router			/auth/register [post]
func (c *authController) RegisterController(ctx *gin.Context) {

	var registerDto dto.UserDto
	err := ctx.ShouldBind(&registerDto)

	if err != nil {
		res := utils.BuildResponseFailed(http.StatusBadRequest, constants.MESSAGE_MAPDATA_ERR, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.authService.RegisterService(ctx.Request.Context(), registerDto)

	if err != nil {
		res := utils.BuildResponseFailed(http.StatusBadRequest, constants.MESSAGE_BAD_REQUEST, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(http.StatusOK, constants.MESSAGE_SUCCESS, result)
	ctx.JSON(http.StatusOK, res)

}
