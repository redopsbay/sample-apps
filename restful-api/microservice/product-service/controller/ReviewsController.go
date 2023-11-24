package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"productservice/logging"
	"productservice/models"
)

func GetReviews(c *gin.Context) {
	err := CheckEnv()
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	resp, err := http.Get(fmt.Sprintf("http://%s/reviews/", REVIEWS_SERVICE))
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	defer resp.Body.Close()

	var reviews interface{}

	json.NewDecoder(resp.Body).Decode(&reviews)

	c.JSON(http.StatusOK, reviews)
}

func GetReviewsById(c *gin.Context) {
	err := CheckEnv()
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	reviewsID := c.Param("id")

	resp, err := http.Get(fmt.Sprintf("http://%s/reviews/%v", REVIEWS_SERVICE, reviewsID))
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	defer resp.Body.Close()

	var reviewsResponse models.Reviews

	json.NewDecoder(resp.Body).Decode(&reviewsResponse)

	c.JSON(http.StatusOK, reviewsResponse)
}

func GetReviewsByProductId(c *gin.Context) {
	err := CheckEnv()
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	reviewsProductID := c.Param("product_id")
	resp, err := http.Get(fmt.Sprintf("http://%s/reviews/product/%v", REVIEWS_SERVICE, reviewsProductID))
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	defer resp.Body.Close()

	var reviewsResponse models.Reviews

	json.NewDecoder(resp.Body).Decode(&reviewsResponse)

	c.JSON(http.StatusOK, reviewsResponse)
}

func DeleteReviews(c *gin.Context) {
	err := CheckEnv()
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	reviewsID := c.Param("id")

	var emptyByte []byte
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s/reviews/%v", REVIEWS_SERVICE, reviewsID),
		bytes.NewBuffer(emptyByte))

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	client := &http.Client{}
	client.Do(req)

	defer req.Body.Close()

	Success(c)
}

func UpdateReviews(c *gin.Context) {
	err := CheckEnv()
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	var reviews map[string]interface{}

	if err := c.BindJSON(&reviews); err != nil {
		ServerError(c)
	}

	json_data, err := json.Marshal(reviews)

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/reviews/update", REVIEWS_SERVICE),
		"application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	defer resp.Body.Close()

	SuccessCreated(c)
}

func AddReviews(c *gin.Context) {
	err := CheckEnv()
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	var reviews models.Reviews

	if err := c.ShouldBindJSON(&reviews); err != nil {
		ServerError(c)
	}

	json_data, err := json.Marshal(reviews)

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/reviews/create", REVIEWS_SERVICE),
		"application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	defer resp.Body.Close()

	SuccessCreated(c)
}
