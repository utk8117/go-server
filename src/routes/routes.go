package routes

import (
	controller "assignment/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(router *gin.Engine) {
	publicRoutes := router.Group("/")
	publicRoutes.POST("fetch_details", controller.TrackUsers)               // fetches products of a particular user
	publicRoutes.POST("add_product_for_user", controller.AddProduct)        // adds products viewed for particular user
	publicRoutes.POST("check_user", controller.CheckUser)                   // checks for admin user
	publicRoutes.POST("add_product", controller.NewProduct)                 // adds new product to product table
	publicRoutes.POST("get_recommendations", controller.GetRecommendations) // get recommendation based on type since user is viewing a product he has all details about the product
}
