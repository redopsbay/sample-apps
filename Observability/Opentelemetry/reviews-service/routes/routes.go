package routes

import (
	"github.com/gin-gonic/gin"
	"reviews/controller"
)

func SetupRoute(rg *gin.Engine) {
	routerGroup := rg.Group("/")
	ReviewsRoutes(routerGroup)
}

func ReviewsRoutes(rg *gin.RouterGroup) {
	reviews := rg.Group("/reviews")
	reviews.GET("/", controller.GetReviews)
	reviews.POST("/create", controller.AddReviews)
	reviews.POST("/update/:id", controller.UpdateReviews)
	reviews.PUT("/:id", controller.DeleteReviews)
	reviews.GET("/:id", controller.GetReviewsById)
	reviews.GET("/product/:product_id", controller.GetReviewsByProductId)
}
