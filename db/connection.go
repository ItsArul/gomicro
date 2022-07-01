package db

import (
	"fmt"
	"os"

	"github.com/ItsArul/gomicro/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}

	RunApp struct {
		Host string
		Port string
	}

	JWTSecret string
}

var app *App
var DB *gorm.DB

func GetConnection() *gorm.DB {
	config := GetConfig()

	logon := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
	)

	db, err := gorm.Open(mysql.Open(logon), &gorm.Config{})
	if err != nil {
		utils.PanicIfError(err)
	}

	DB = db

	return db
}

func GetConfig() *App {
	if app == nil {
		app = initConfig()
	}

	return app
}

func initConfig() *App {
	var config App
	err := godotenv.Load()
	if err != nil {
		config.Database.Host = "localhost"
		config.Database.Port = "3306"
		config.Database.User = "root"
		config.Database.Password = "12345"
		config.Database.DBName = "gomicro"
		config.RunApp.Host = "localhost"
		config.RunApp.Port = "5000"
		config.JWTSecret = "secretjwt"

		return &config
	}

	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Database.User = os.Getenv("DB_USER")
	config.Database.Password = os.Getenv("DB_PASS")
	config.Database.DBName = os.Getenv("DB_NAME")
	config.RunApp.Host = os.Getenv("APP_HOST")
	config.RunApp.Port = os.Getenv("APP_PORT")
	config.JWTSecret = os.Getenv("JWT_SECRET")

	return &config

}
