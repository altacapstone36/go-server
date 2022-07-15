package middleware

import (
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/utils/errors/check"
	"go-hospital-server/internal/utils/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
)


func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := jwt.GetToken(c, jwt.ACCESS)

		if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
			return c.JSON(r.Code, r.Result)
		}

		err = jwt.FindToken(token)

		if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
			return c.JSON(r.Code, r.Result)
		}

		return next(c)
	}
}

func AdminPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := jwt.GetToken(c, jwt.ACCESS)
		role, err := jwt.GetTokenData(token, "role", jwt.ACCESS)

		if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
			return c.JSON(r.Code, r.Result)
		}

		if role != "administrator" && role != "admin" {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: "access for this route only for administrator",
			})
		}

		return next(c)
	}
}

func DoctorPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := jwt.GetToken(c, jwt.ACCESS)
		role, err := jwt.GetTokenData(token, "role", jwt.ACCESS)

		if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
			return c.JSON(r.Code, r.Result)
		}

		if role != "doctor" {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: "access for this route only for doctor",
			})
		}

		return next(c)
	}
}

func NursePermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := jwt.GetToken(c, jwt.ACCESS)
		role, err := jwt.GetTokenData(token, "role", jwt.ACCESS)

		if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
			return c.JSON(r.Code, r.Result)
		}

		if role != "nurse" {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: "access for this route only for nurse",
			})
		}

		return next(c)
	}
}

func DoctorNursePermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := jwt.GetToken(c, jwt.ACCESS)
		role, err := jwt.GetTokenData(token, "role", jwt.ACCESS)

		if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
			return c.JSON(r.Code, r.Result)
		}

		if role != "doctor" && role != "nurse" {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: "access for this route only for doctor and nurse",
			})
		}

		return next(c)
	}
}

