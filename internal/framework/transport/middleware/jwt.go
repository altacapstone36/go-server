package middleware

import (
	"context"
	"go-hospital-server/internal/core/entity/response"
	ujwt "go-hospital-server/internal/utils/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Database

func NewJWTConnection(mongo *mongo.Database) {
	client = mongo
}

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := ujwt.GetToken(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: err.Error(),
			})
		}

		filter := bson.D{
			{Key: "access_token", Value: token},
		}

		db := client.Collection("token")
		_, err = db.Find(context.TODO(), filter)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: "invalid or expired token",
			})
		}

		return next(c)
	}
}

func AdminPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, err := ujwt.GetTokenData(c, "role")

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: err.Error(),
			})
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
		role, err := ujwt.GetTokenData(c, "role")

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: err.Error(),
			})
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
		role, err := ujwt.GetTokenData(c, "role")

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: err.Error(),
			})
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
		role, err := ujwt.GetTokenData(c, "role")

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: err.Error(),
			})
		}


		if role != "doctor" && role != "nurse" {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: "access for this route only for doctor and nurse",
			})
		}
		return next(c)
	}
}

