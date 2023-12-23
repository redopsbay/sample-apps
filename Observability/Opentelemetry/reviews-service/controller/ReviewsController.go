package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reviews/db"
	"reviews/logging"
	"reviews/models"
	"strconv"
	"reviews/tracer"
	//oteltrace "go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel"
	//"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
)

// GET /reviews

func GetReviews(c *gin.Context) {

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "GetReviews")
	defer span.End()

	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	reviews, err := sql.GetReviews(c)
	if err != nil {
		ServerError(c)
	}

	c.JSON(http.StatusOK, reviews)
}

// GET /reviews/<int:id>

func GetReviewsById(c *gin.Context) {
	dbInstance, err := db.SetupDatabase()

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "GetReviewsById")
		
	defer span.End()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var reviewsID models.Reviews

	if err := c.ShouldBindUri(&reviewsID); err == nil {
		
		reviews, err := sql.GetReviewsById(c, reviewsID.ID)
		if err != nil {
			ServerError(c)
			return
		}

		c.JSON(http.StatusOK, reviews)
	}
}

// GET /reviews/<int:product_id>
func GetReviewsByProductId(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "GetReviewsById")
		
	defer span.End()


	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	reviewsProductID, err := strconv.Atoi(c.Param("product_id"))

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	reviews, err := sql.GetReviewsByProductId(c, reviewsProductID)
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	c.JSON(http.StatusOK, reviews)
}

// DELETE /reviews/<int:id>

func DeleteReviews(c *gin.Context) {
	dbInstance, err := db.SetupDatabase()

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "GetReviewsById")
		
	defer span.End()


	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var reviewsID models.Reviews

	if err := c.ShouldBindUri(&reviewsID); err == nil {
		err := sql.DeleteReviews(c, &reviewsID)
		if err != nil {
			ServerError(c)
			return
		}

		Success(c)
	}
}

// PUT /reviews/<int:id>

func UpdateReviews(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "GetReviewsById")
		
	defer span.End()


	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var reviews models.Reviews
	var reviewsUri models.Reviews

	if c.ShouldBindUri(&reviewsUri) == nil {
		if c.ShouldBindJSON(&reviews) == nil {
			reviews.ID = reviewsUri.ID
			if sql.UpdateReviews(c, &reviews) == nil {
				SuccessCreated(c)
			} else {
				ServerError(c)
			}
		}

	} else {
		ServerError(c)
	}

}

// POST /reviews/add

func AddReviews(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "GetReviewsById")
		
	defer span.End()


	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var reviews models.Reviews

	if c.ShouldBindJSON(&reviews) == nil {
		if sql.AddReviews(c, &reviews) == nil {
			SuccessCreated(c)
		} else {
			ServerError(c)
		}

	} else {
		ServerError(c)
	}

}
