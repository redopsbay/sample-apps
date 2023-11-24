package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "unknown server error",
	})
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"error": "forbidden",
	})
}

func SuccessCreated(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
