package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reviews/db"
	"reviews/logging"
	"reviews/models"
	"strconv"
)

// GET /reviews

func GetReviews(c *gin.Context) {
	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	reviews, err := sql.GetReviews()
	if err != nil {
		ServerError(c)
	}

	c.JSON(http.StatusOK, reviews)
}

// GET /reviews/<int:id>

func GetReviewsById(c *gin.Context) {
	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var reviewsID models.Reviews

	if err := c.ShouldBindUri(&reviewsID); err == nil {
		reviews, err := sql.GetReviewsById(reviewsID.ID)
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

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	reviewsProductID, err := strconv.Atoi(c.Param("product_id"))

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	reviews, err := sql.GetReviewsByProductId(reviewsProductID)
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	c.JSON(http.StatusOK, reviews)
}

// DELETE /reviews/<int:id>

func DeleteReviews(c *gin.Context) {
	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var reviewsID models.Reviews

	if err := c.ShouldBindUri(&reviewsID); err == nil {
		err := sql.DeleteReviews(&reviewsID)
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

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var reviews models.Reviews
	var reviewsUri models.Reviews

	if c.ShouldBindUri(&reviewsUri) == nil {
		if c.ShouldBindJSON(&reviews) == nil {
			reviews.ID = reviewsUri.ID
			if sql.UpdateReviews(&reviews) == nil {
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

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var reviews models.Reviews

	if c.ShouldBindJSON(&reviews) == nil {
		if sql.AddReviews(&reviews) == nil {
			SuccessCreated(c)
		} else {
			ServerError(c)
		}

	} else {
		ServerError(c)
	}

}
