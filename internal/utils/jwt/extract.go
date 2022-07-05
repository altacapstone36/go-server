package jwt

import (
	"context"
	"fmt"
	"go-hospital-server/internal/utils/config"
	"go-hospital-server/internal/utils/errors"
	"go-hospital-server/internal/utils/logger"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


var client *mongo.Database

func NewJWTConnection(mongo *mongo.Database) {
	client = mongo
}

func GetTokenData(header string, data string) (extracted interface{}, err error) {
	err = FindToken(header)
	if err != nil {
		return
	}

	extract, err := ExtractToken(header)
	if extract != nil {
		extracted = extract.(jwt.MapClaims)[data]
	}

	return
}

func GetToken(c echo.Context) (header string, err error) {
	header = c.Request().Header.Get("Authorization")
	headers := strings.Split(header, " ")

	if header == "" || len(headers) < 2 {
		err = errors.ErrNoToken
		return
	}

	header = headers[1]
	_, err = ExtractToken(header)
	return
}

func ExtractToken(tkn string) (token interface{}, err error) {
	token, err = jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return config.SERVER_SECRET, nil
	})
	if err != nil {
		return nil, err
	}
	return token.(*jwt.Token).Claims, nil
}

func FindToken(token string) error {
	var data []bson.M
	filter := bson.D{
		{Key: "access_token", Value: token},
	}

	db := client.Collection("token")
	cur, err := db.Find(context.TODO(), filter)

	if err != nil {
		logger.WriteLog(err)
	}

	if err = cur.All(context.TODO(), &data); err != nil {
		logger.WriteLog(err)
	}
	
	if data == nil {
		return errors.ErrInvalidToken
	}

	return nil
}

