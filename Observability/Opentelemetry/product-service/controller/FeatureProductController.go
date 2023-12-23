package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"productservice/db"
	"productservice/logging"
	"productservice/models"
	"strconv"
	"time"
	"productservice/tracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

)

func GetFeaturedProduct(c *gin.Context) {

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "GetFeaturedProduct")
	defer span.End()


	numFeaturedProduct := c.Query("count")
	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	products, err := sql.GetProducts(c)
	if err != nil {
		ServerError(c)
		return
	}

	countFeatureProduct, err := strconv.Atoi(numFeaturedProduct)
	if err != nil {
		ServerError(c)
		return
	}

	fProducts, err := GetRandomFeaturedProduct(countFeatureProduct, products)
	if err != nil {
		logging.Error(err)
		ServerError(c)
		return
	}

	c.JSON(http.StatusOK, fProducts)
}

// GetRandomFeaturedProduct will randomly retrieved featured product from sliced []models.FeaturedProduct argument.
// it will also return sliced []model.FeaturedProduct that contains only the randomly retrieved FeaturedProduct.

func GetRandomFeaturedProduct(count int, featuredProduct []models.Product) ([]models.Product, error) {
	var Products []models.Product
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	if len(featuredProduct) >= count {
		for i := 0; i < count; i++ {

			Products = append(Products, featuredProduct[rand.Intn(len(featuredProduct))])
		}
		return Products, nil
	} else {
		return nil, errors.New("random count is greater than available featuredProduct.")
	}

}
