package controller

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/utils/errors/check"
	"go-hospital-server/internal/utils/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
)


type AuthController struct {
	srv *service.AuthService
}

func NewAuthController(srv *service.AuthService) *AuthController {
	return &AuthController{srv}
}

// godoc
// @Summary Login
// @Description Login and get Authorization Token
// @Tags Authorization
// @Accept json
// @Produce json
// @Param body  body  request.Login{}  true "Send Request Email and Password"
// @Success 200 {object} response.User{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /login [post]
func (acon AuthController) Login(c echo.Context) error {
	var login request.Login
	c.Bind(&login)

	if r, ok := check.HTTP(nil, login.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	res, err := acon.srv.Login(login)
	if r, ok := check.HTTP(res, err, "User Log In"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	jwt, err := acon.srv.CreateToken(res.ID, res.Facility, res.Role)
	if r, ok := check.HTTP(res, err, "Create Authentication Token"); !ok {
		return c.JSON(r.Code, r.Result)
	}
	
	return c.JSON(http.StatusOK, response.MessageDataJWT{
		Message: "User Logged In",
		Data: res,
		JWT: jwt,
	})
}

// godoc
// @Summary Refresh Token
// @Description Route Path for Get New Access Token
// @Tags Authorization
// @Security ApiKey
// @Accept json
// @Produce json
// @Param token  query  string  true "Pass access_token Here"
// @Success 200 {object} models.Token{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /refresh_token [post]
func (acon AuthController) RefreshToken(c echo.Context) error {
	var t models.Token
	token, err := jwt.GetToken(c)

	if err == nil {
		t.RefreshToken = token
		t.AccessToken = c.QueryParam("token")

		t, err = acon.srv.RefreshToken(t)
	}

	if err != nil {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Generate New Token",
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.MessageData{
		Message: "New Token Generated",
		Data: t,
	})
}

// godoc
// @Summary Refresh Token
// @Description Route Path for Get New Access Token
// @Tags Authorization
// @Security ApiKey
// @Accept json
// @Produce json
// @Success 200 {object} models.Token{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /logout [post]
func (acon AuthController) Logout(c echo.Context) error {
	token, err := jwt.GetToken(c)

	if err == nil {
		err = acon.srv.Logout(token)

	}
	
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Revoke Token",
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User Logged Out",
	})
}
