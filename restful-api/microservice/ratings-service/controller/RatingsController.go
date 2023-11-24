package controller

import (
	"github.com/gin-gonic/gin"
	//"github.com/google/uuid"
	"net/http"
	"ratings/db"
	"ratings/logging"
	"ratings/models"
	"strconv"
)

// GET /ratings

func GetRatings(c *gin.Context) {
	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	ratings, err := sql.GetRatings()
	if err != nil {
		ServerError(c)
	}

	c.JSON(http.StatusOK, ratings)
}

// GET /ratings/<int:id>

func GetRatingsById(c *gin.Context) {
	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var ratingsID models.Ratings

	if err := c.ShouldBindUri(&ratingsID); err == nil {
		ratings, err := sql.GetRatingsById(ratingsID.ID)
		if err != nil {
			ServerError(c)
			return
		}

		c.JSON(http.StatusOK, ratings)
	}
}

// GET /ratings/product/<int:product_id>
func GetRatingsByProductId(c *gin.Context) {
	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	ratingsProductID, err := strconv.Atoi(c.Param("product_id"))

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	ratings, err := sql.GetRatingsByProductId(ratingsProductID)

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	c.JSON(http.StatusOK, ratings)
}

// DELETE /ratings/<int:id>

func DeleteRatings(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var ratings models.Ratings

	if c.ShouldBindJSON(&ratings) == nil {
		if sql.DeleteRating(&ratings) == nil {
			SuccessCreated(c)
		} else {
			ServerError(c)
		}

	} else {
		ServerError(c)
	}

}

// PUT /ratings/<int:id>

func UpdateRatings(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var ratings models.Ratings
	var ratingsUri models.Ratings

	if c.ShouldBindUri(&ratingsUri) == nil {
		if c.ShouldBindJSON(&ratings) == nil {
			ratings.ID = ratingsUri.ID
			if sql.UpdateRatings(&ratings) == nil {
				SuccessCreated(c)
			} else {
				ServerError(c)
			}
		}

	} else {
		ServerError(c)
	}

}

// POST /ratings/create

func CreateRatings(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var ratings models.Ratings

	if c.ShouldBindJSON(&ratings) == nil {
		if sql.AddRatings(&ratings) == nil {
			SuccessCreated(c)
		} else {
			ServerError(c)
		}

	} else {
		ServerError(c)
	}

}
