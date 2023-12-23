package routes

import (
	"github.com/gin-gonic/gin"
	"productservice/controller"
)

func SetupRoute(rg *gin.Engine) {
	routerGroup := rg.Group("/")
	ProductRoutes(routerGroup)
	UserRoutes(routerGroup)
}

func ProductRoutes(rg *gin.RouterGroup) {
	products := rg.Group("/product")
	products.GET("/", controller.GetProducts)
	products.POST("/create", controller.CreateProduct)
	products.POST("/update/:id", controller.UpdateProduct)
	products.PUT("/:id", controller.DeleteProduct)
	products.GET("/:id", controller.GetProductById)
	products.GET("/featuredproduct", controller.GetFeaturedProduct)
	products.GET("/reviews", controller.GetReviews)
	products.POST("/reviews/create", controller.AddReviews)
	products.POST("/reviews/update/:id", controller.UpdateReviews)
	products.PUT("/reviews/:id", controller.DeleteReviews)
	products.GET("/reviews/:id", controller.GetReviewsById)
	products.GET("/reviews/product/:product_id", controller.GetReviewsByProductId)
	products.GET("/ratings", controller.GetRatings)
	products.POST("/ratings/create", controller.AddRatings)
	products.POST("/ratings/update/:id", controller.UpdateRatings)
	products.PUT("/ratings/:id", controller.DeleteRatings)
	products.GET("/ratings/:id", controller.GetRatingsById)
	products.GET("/ratings/product/:product_id", controller.GetRatingsByProductId)
}

func UserRoutes(rg *gin.RouterGroup) {
	products := rg.Group("/user")
	products.GET("/", controller.GetUsers)
	products.POST("/create", controller.CreateUser)
	products.PUT("/:id", controller.DeleteUser)
	products.GET("/:id", controller.GetUserById)
	products.POST("/login", controller.UserLogin)
}
