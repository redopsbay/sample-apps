package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ratings/tracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

)

func ServerError(c *gin.Context) {

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "ServerError")
	defer span.End()

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "unknown server error",
	})
}

func Forbidden(c *gin.Context) {

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "Forbidden")
	defer span.End()

	c.JSON(http.StatusForbidden, gin.H{
		"error": "forbidden",
	})
}

func SuccessCreated(c *gin.Context) {

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "SuccessCreated")
	defer span.End()

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

func Success(c *gin.Context) {

	otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	_, span := tracer.Tracer.Start(c.Request.Context(), "Success")
	defer span.End()

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
