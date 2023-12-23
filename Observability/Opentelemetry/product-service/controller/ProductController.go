package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"productservice/db"
	"productservice/models"
	"time"
	"productservice/tracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

)

// GET /products

func GetProducts(c *gin.Context) {
	dbInstance, err := db.SetupDatabase()

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "GetProducts")
	defer span.End()


	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	products, err := sql.GetProducts(c)
	if err != nil {
		ServerError(c)
	}

	c.JSON(http.StatusOK, products)
}

// GET /products/<int:id>

func GetProductById(c *gin.Context) {
	dbInstance, err := db.SetupDatabase()

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "GetProductById")
	defer span.End()


	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var productID models.Product

	if err := c.ShouldBindUri(&productID); err == nil {
		product, err := sql.GetProductById(c, productID.ID)
		if err != nil {
			ServerError(c)
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

// DELETE /products/<int:id>

func DeleteProduct(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "DeleteProduct")
	defer span.End()


	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var product models.Product

	if c.ShouldBindJSON(&product) == nil {
		if sql.DeleteProduct(c, &product) == nil {
			SuccessCreated(c)
		} else {
			ServerError(c)
		}

	} else {
		ServerError(c)
	}

}

// PUT /products/<int:id>

func UpdateProduct(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "UpdateProduct")
	defer span.End()


	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var product models.Product
	var productUri models.Product

	if c.ShouldBindUri(&productUri) == nil {
		if c.ShouldBindJSON(&product) == nil {
			product.ID = productUri.ID
			if sql.UpdateProduct(c, &product) == nil {
				SuccessCreated(c)
			} else {
				ServerError(c)
			}
		}

	} else {
		ServerError(c)
	}

}

// POST /products/create

func CreateProduct(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "CreateProduct")
	defer span.End()


	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var product models.Product

	if c.ShouldBindJSON(&product) == nil {
		// Removed Owner_UUID from models.
		// uuid := uuid.NewString()
		// product.Owner_UUID = uuid
		product.Timestamp = time.Now().String()
		if sql.CreateProduct(c, &product) == nil {
			SuccessCreated(c)
		} else {
			ServerError(c)
		}

	} else {
		ServerError(c)
	}

}
