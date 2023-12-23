package routes

import (
	"github.com/gin-gonic/gin"
	"ratings/controller"
)

func SetupRoute(rg *gin.Engine) {
	routerGroup := rg.Group("/")
	RatingsRoutes(routerGroup)
}

func RatingsRoutes(rg *gin.RouterGroup) {
	ratings := rg.Group("/ratings")
	ratings.GET("/", controller.GetRatings)
	ratings.POST("/create", controller.CreateRatings)
	ratings.POST("/update/:id", controller.UpdateRatings)
	ratings.PUT("/:id", controller.DeleteRatings)
	ratings.GET("/:id", controller.GetRatingsById)
	ratings.GET("/product/:product_id", controller.GetRatingsByProductId)
}
