package routes

import (
	"go-hospital-server/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewAuthRoutes(e *echo.Group, acon *controller.AuthController, middleware ...echo.MiddlewareFunc) {
	e.POST("/refresh_token", acon.RefreshToken, middleware...)
	e.POST("/logout", acon.Logout, middleware...)
	e.POST("/login", acon.Login)
	e.POST("/register", acon.Register)
	e.POST("/forgot_password", acon.ForgotPassword)
	e.POST("/find_email", acon.FindEmail)
	e.GET("/profile", acon.Profile, middleware...)
	e.POST("/profile/change_password", acon.ChangePassword, middleware...)
}
