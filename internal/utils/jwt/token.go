package jwt

import (
	"fmt"
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/utils/config"
	"go-hospital-server/internal/utils/errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type Token string
const (
	ACCESS Token = "access"
	RESET Token = "reset"
)

type mapClaims struct {
	Code string `json:"code"`
	Role string `json:"role"`
	jwt.StandardClaims
}

// Crate New Create
// Param user_code string
// Param user_facility_id string
// Param user_level string
// Param exp_time int
func CreateToken(code,level string, types Token) (t models.Token, err error) {
	role := strings.ToLower(level)
	var expTime int64
	var sign []byte

	if types == ACCESS {

		rexpTime := time.Now().Add(time.Hour * time.Duration(config.JWT_REFRESH_EXPIRE_TIME)).Unix()
		rclaims := jwt.StandardClaims{
			ExpiresAt: rexpTime,
		}

		rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, rclaims)
		t.RefreshToken, err = rtoken.SignedString(config.ACCESS_KEY)

		expTime = time.Now().Add(time.Hour * time.Duration(config.JWT_ACCESS_EXPIRE_TIME)).Unix()
		sign = config.ACCESS_KEY
	} else if types == RESET {
		fmt.Println(config.JWT_RESET_PASSWORD_EXPIRE_TIME)
		expTime = time.Now().Add(time.Minute * time.Duration(config.JWT_RESET_PASSWORD_EXPIRE_TIME)).Unix()
		sign = config.RESET_KEY
	}

	claims := mapClaims{
		Code: code,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t.AccessToken, err = token.SignedString(sign)
	return
}

func RefreshToken(token_string models.Token) (t models.Token, err error) {
	token, err := ExtractToken(token_string.RefreshToken, ACCESS)

	if _, ok := token.(jwt.MapClaims); ok {
		tkn, _ := ExtractToken(token_string.AccessToken, ACCESS)
		code := tkn.(jwt.MapClaims)["code"]
		role := tkn.(jwt.MapClaims)["role"]
		if code != nil && role != nil {
			return CreateToken(code.(string), role.(string), ACCESS)
		}
	}
	return t, errors.New(500, "failed to generate new token")
}
