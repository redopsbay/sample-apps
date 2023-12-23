package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"reviews/logging"
	"reviews/models"
	"reviews/tracer"
	"go.opentelemetry.io/otel/attribute"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
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

	dbInstance.AutoMigrate(models.Reviews{})

	otelDBPlugin := otelgorm.NewPlugin(
			otelgorm.WithDBName("microservice"),
			otelgorm.WithAttributes(
				attribute.String("service.name", tracer.ServiceName),
            	attribute.String("library.language", "go"),
				attribute.String("database.driver", "mysql")),
			)

	if err := dbInstance.Use(otelDBPlugin); err != nil {
		panic(err)
	}
	

	return dbInstance, nil
}
