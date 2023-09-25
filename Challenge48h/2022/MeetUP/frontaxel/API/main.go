package main

import (
	"meetupAPI/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Tells to gin if we are in a dev environment or not
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	// Tells to gin to force color in shell
	gin.ForceConsoleColor()

	router := gin.Default()

	router.Use(cors.Default())

	// router.Static("/images", "./images") // for static path

	// router.GET("/", controllers.Base)

	// User CRUD
	router.GET("/user/:id", controllers.GetUserById)
	router.PUT("/user/:id", controllers.PutUserById)
	router.GET("/users", controllers.GetUsers)

	// Category CRUD
	router.GET("/categories", controllers.GetCategories)
	router.GET("/category/:id", controllers.GetCategoryById)
	router.PUT("/category/:id", controllers.PutCategoryById)
	router.DELETE("/category/:id", controllers.DeleteCategoryById)
	router.POST("/category", controllers.PostCategory)

	router.GET("/", controllers.RootEvent)

	// publicGroup := router.Group("/")
	// {
	// 	// Admin
	// 	publicGroup.POST("/admin/login", public.LoginAdmin)

	// 	// User
	// 	publicGroup.GET("/user/:id", public.GetUserById)
	// 	publicGroup.POST("/user", public.PostUser)
	// 	publicGroup.POST("/user/login", public.LoginUser)

	// 	// Wishlist
	// 	publicGroup.GET("/user/:id/wishlist", public.GetUserWishlist)
	// 	publicGroup.PUT("/user/:id/wishlist", public.PutUserWishlist)

	// 	// Category

	// 	// Product
	// 	publicGroup.GET("/products", public.GetProducts)
	// 	publicGroup.GET("/product/:id", public.GetProductById)

	// 	// Newsletters
	// 	publicGroup.GET("/newsletters", public.GetNewsletters)
	// 	publicGroup.POST("/newsletter", public.PostNewsletter)

	// 	// Should be private under

	// 	// User
	// 	publicGroup.GET("/users", public.GetUsers)
	// 	publicGroup.PUT("/user/:id", public.PutUserById)
	// 	publicGroup.DELETE("/user/:id", public.DeleteUserById)

	// 	// Category

	// 	// Product
	// 	publicGroup.PUT("/product/:id", public.PutProductById)
	// 	publicGroup.DELETE("/product/:id", public.DeleteProductById)
	// 	publicGroup.POST("/product", public.PostProduct)
	// }

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()

}
