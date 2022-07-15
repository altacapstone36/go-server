package controller

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/utils/errors/check"
	"go-hospital-server/internal/utils/jwt"

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
// @Success 200 {object} response.MessageDataJWT{data=response.User{}} success
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

	jwt, err := acon.srv.CreateToken(res.Code, res.Role, jwt.ACCESS)
	if r, ok := check.HTTP(res, err, "Create Authentication Token"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageDataJWT{
		Message: "User Logged In",
		Data:    res,
		JWT:     jwt,
	})
}

// godoc
// @Summary Forgot Password
// @Description Login and get Authorization Token
// @Tags Authorization
// @Accept json
// @Produce json
// @Param token query string true "token from find email"
// @Param body body request.ChangePassword{} true "Send Request New Password Password"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /forgot_password [post]
func (acon AuthController) ForgotPassword(c echo.Context) error {
	token := c.QueryParam("token")
	code, err := jwt.GetTokenData(token, "code", jwt.RESET)

	if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	var cp request.ChangePassword
	c.Bind(&cp)
	cp.Code = code.(string)

	if r, ok := check.HTTP(nil, cp.Validate(), "Validate Data"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err = acon.srv.ChangePassword(cp)
	if r, ok := check.HTTP(nil, err, "Change Password"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	acon.srv.RevokeToken(token)

	return c.JSON(200, response.MessageOnly{
		Message: "Password Changed",
	})
}

// godoc
// @Summary Register
// @Description Login and get Authorization Token
// @Tags Authorization
// @Accept json
// @Produce json
// @Param body body request.UserRequest{} true "Send Request New User"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /register [post]
func (acon AuthController) Register(c echo.Context) error {
	var req request.UserRequest
	c.Bind(&req)

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := acon.srv.Register(req)
	if r, ok := check.HTTP(nil, err, "Register"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(201, response.MessageOnly{
		Message: "Register success",
	})
}

// godoc
// @Summary UserProfile
// @Description Show Profile from Logged in User
// @Tags Authorization
// @Security ApiKey
// @Accept json
// @Produce json
// @Success 200 {object} response.MessageDataJWT{data=response.User{}} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /profile [get]
func (acon AuthController) Profile(c echo.Context) error {
	token, err := jwt.GetToken(c, jwt.ACCESS)
	code, _ := jwt.GetTokenData(token, "code", jwt.ACCESS)
	res, err := acon.srv.MyProfile(code.(string))
	if r, ok := check.HTTP(res, err, "Get User Profile"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "User Profile Fetched",
		Data:    res,
	})
}

// godoc
// @Summary FindEmail
// @Description Find Email and Get Token to Change Password
// @Tags Authorization
// @Accept json
// @Produce json
// @Param body body request.FindEmail{} true "Send Request Email to change password"
// @Success 200 {object} response.MessageDataJWT{data=response.User{}} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /find_email [get]
func (acon AuthController) FindEmail(c echo.Context) error {
	var fe request.FindEmail

	c.Bind(&fe)

	res, err := acon.srv.FindEmail(fe)
	if r, ok := check.HTTP(res, err, "Find Email"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	jwt, err := acon.srv.CreateToken(res.Code, res.Role, jwt.RESET)
	if r, ok := check.HTTP(res, err, "Create Authentication Token"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageDataJWT{
		Message: "Email found",
		Data:    res,
		JWT: jwt,
	})
}

// godoc
// @Summary Change Password
// @Description Change Password in Profile Level
// @Tags Authorization
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body body request.ChangePassword{} true "Send New Password"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /profile/change_password [post]
func (acon AuthController) ChangePassword(c echo.Context) error {
	token, err := jwt.GetToken(c, jwt.ACCESS)
	if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	code, err := jwt.GetTokenData(token, "code", jwt.ACCESS)

	var cp request.ChangePassword
	c.Bind(&cp)
	cp.Code = code.(string)

	if r, ok := check.HTTP(nil, cp.Validate(), "Validate Data"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err = acon.srv.ChangePassword(cp)
	if r, ok := check.HTTP(nil, err, "Change Password"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: "Password changed",
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
// @Success 200 {object} response.MessageData{data=models.Token{}} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /refresh_token [post]
func (acon AuthController) RefreshToken(c echo.Context) error {
	var t models.Token
	token, err := jwt.GetToken(c, jwt.ACCESS)

	if err == nil {
		t.RefreshToken = token
		t.AccessToken = c.QueryParam("token")

		t, err = acon.srv.RefreshToken(t)
	}

	if err != nil {
		return c.JSON(417, response.Error{
			Message: "Failed to Generate New Token",
			Error:   err.Error(),
		})
	}

	return c.JSON(200, response.MessageData{
		Message: "New Token Generated",
		Data:    t,
	})
}

// godoc
// @Summary Refresh Token
// @Description Route Path for Get New Access Token
// @Tags Authorization
// @Security ApiKey
// @Accept json
// @Produce json
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /logout [post]
func (acon AuthController) Logout(c echo.Context) error {
	token, err := jwt.GetToken(c, jwt.ACCESS)

	if err == nil {
		err = acon.srv.RevokeToken(token)
	}

	if err != nil {
		return c.JSON(417, echo.Map{
			"message": "Failed to Revoke Token",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "User Logged Out",
	})
}
