package main

import (
	"github.com/gin-gonic/gin"
	"reviews/routes"
)

func main() {
	router := gin.Default()
	routes.SetupRoute(router)
	router.Run(":8090")
}
