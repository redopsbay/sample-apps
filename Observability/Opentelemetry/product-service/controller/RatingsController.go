package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"productservice/logging"
	"productservice/models"
	"productservice/tracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

)

func GetRatings(c *gin.Context) {

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "GetRatings")
	defer span.End()

	err := CheckEnv()
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	resp, err := otelhttp.Get(c, fmt.Sprintf("http://%s/ratings/", RATINGS_SERVICE))
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	defer resp.Body.Close()

	var ratings interface{}

	json.NewDecoder(resp.Body).Decode(&ratings)

	c.JSON(http.StatusOK, ratings)
}

func GetRatingsById(c *gin.Context) {

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "GetRatingsById")
	defer span.End()

	err := CheckEnv()
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	ratingsID := c.Param("id")

	resp, err := otelhttp.Get(c, fmt.Sprintf("http://%s/ratings/%v", RATINGS_SERVICE, ratingsID))
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	defer resp.Body.Close()

	var ratingsResponse models.Ratings

	json.NewDecoder(resp.Body).Decode(&ratingsResponse)

	c.JSON(http.StatusOK, ratingsResponse)
}

func GetRatingsByProductId(c *gin.Context) {

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "GetRatingsByProductId")
	defer span.End()

	err := CheckEnv()
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	ratingsProductID := c.Param("product_id")

	resp, err := otelhttp.Get(c, fmt.Sprintf("http://%s/ratings/product/%v", RATINGS_SERVICE, ratingsProductID))
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	defer resp.Body.Close()

	var ratingsResponse models.Ratings

	json.NewDecoder(resp.Body).Decode(&ratingsResponse)

	c.JSON(http.StatusOK, ratingsResponse)
}

func DeleteRatings(c *gin.Context) {

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "DeleteRatings")
	defer span.End()

	err := CheckEnv()
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	ratingsID := c.Param("id")

	var emptyByte []byte
	req, err := otelhttp.NewRequest(c, http.MethodPut, fmt.Sprintf("http://%s/ratings/%v", RATINGS_SERVICE, ratingsID),
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

func UpdateRatings(c *gin.Context) {

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "UpdateRatings")
	defer span.End()

	err := CheckEnv()
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	var ratings map[string]interface{}

	if err := c.BindJSON(&ratings); err != nil {
		ServerError(c)
	}

	json_data, err := json.Marshal(ratings)

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	resp, err := otelhttp.Post(c, fmt.Sprintf("http://%s/ratings/update", RATINGS_SERVICE),
		"application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	defer resp.Body.Close()

	SuccessCreated(c)
}

func AddRatings(c *gin.Context) {

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "AddRatings")
	defer span.End()

	err := CheckEnv()
	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	var ratings models.Ratings

	if err := c.ShouldBindJSON(&ratings); err != nil {
		ServerError(c)
	}

	json_data, err := json.Marshal(ratings)

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}
	resp, err := otelhttp.Post(c, fmt.Sprintf("http://%s/ratings/create", RATINGS_SERVICE),
		"application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		logging.Error(err)
		ServerError(c)
	}

	defer resp.Body.Close()

	SuccessCreated(c)
}
