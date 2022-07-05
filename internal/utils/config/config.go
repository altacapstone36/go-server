package config

import (
	"os"

	"github.com/integralist/go-findroot/find"
	"github.com/spf13/viper"
)

var (
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST string
	MONGODB_STRING string
	MONGODB_DATABASE string
	SERVER_PORT string
	SERVER_SECRET []byte
	SMTP_SERVER string
	SMTP_PORT string
	EMAIL string
	PASSWORD string
	JWT_ACCESS_EXPIRE_TIME int64
	JWT_REFRESH_EXPIRE_TIME int64
	JWT_FORGOT_PASSWORD_EXPIRE_TIME int64
)

func LoadConfig() {
	var path string
	root, err := find.Repo()
	path = root.Path
	if err != nil {
		path = "."
	}

	viper.AddConfigPath(path + "/config")
	viper.SetConfigName("config")
	
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	SERVER_PORT = checkEnv(os.Getenv("PORT"), viper.GetString("server.PORT"))
	MONGODB_STRING = viper.GetString("mongo.STRING")
	MONGODB_DATABASE = viper.GetString("mongo.DATABASE")
	DB_DATABASE = viper.GetString("mysql.DATABASE")
	DB_USERNAME = viper.GetString("mysql.USERNAME")
	DB_PASSWORD = viper.GetString("mysql.PASSWORD")
	DB_HOST = viper.GetString("mysql.HOST")
	SERVER_SECRET= []byte(viper.GetString("server.SECRET"))
	SMTP_SERVER = viper.GetString("smtp.SERVER")
	SMTP_PORT = viper.GetString("smtp.PORT")
	EMAIL = viper.GetString("smtp.EMAIL")
	PASSWORD = viper.GetString("smtp.PASSWORD")
	JWT_ACCESS_EXPIRE_TIME = viper.GetInt64("jwt.ACCESS_TIME")
	JWT_REFRESH_EXPIRE_TIME = viper.GetInt64("jwt.REFRESH_TIME")
	JWT_FORGOT_PASSWORD_EXPIRE_TIME = viper.GetInt64("jwt.FORGOT_PASSWORD_TIME")
}

func checkEnv(env, viperStr string) string {
	if env != "" {
		return env
	}
	return viperStr
}
