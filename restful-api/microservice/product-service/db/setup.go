package db

import (
	//"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"productservice/logging"
	"productservice/models"
)

var (
	DATABASE_USER     = os.Getenv("DATABASE_USER")
	DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
	DATABASE_HOSTNAME = os.Getenv("DATABASE_HOSTNAME")
	DATABASE_PORT     = os.Getenv("DATABASE_PORT")
	DATABASE_NAME     = os.Getenv("DATABASE_NAME")
	CONNECTION_STRING = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		DATABASE_USER,
		DATABASE_PASSWORD,
		DATABASE_HOSTNAME,
		DATABASE_PORT,
		DATABASE_NAME)
)

func SetupDatabase() (*gorm.DB, error) {

	dbInstance, err := gorm.Open(mysql.Open(CONNECTION_STRING), &gorm.Config{})
	if err != nil {
		logging.Info(CONNECTION_STRING)
		logging.Error(err)
		return nil, err
	}

	dbInstance.AutoMigrate(models.Product{})
	dbInstance.AutoMigrate(models.User{})
	//dbInstance.AutoMigrate(models.FeaturedProduct{})

	return dbInstance, nil
}
