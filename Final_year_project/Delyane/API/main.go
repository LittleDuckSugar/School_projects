package main

import (
	"delyaneAPI/public"

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

	router.Static("/images", "./images") // for static path

	// apiFinal := router.Group("/")
	// {
	// 	public := router.Group("/")
	// 	{
	// 		// Demo
	// 		public.GET("/", controllers.GetRoot)
	// 		public.POST("/upload", controllers.SaveImage)

	// 		// Admin
	// 		public.POST("/admin/login", controllers.LoginAdmin)

	// 		// User
	// 		public.GET("/user/:id", controllers.GetUserById)
	// 		public.POST("/user", controllers.PostUser)
	// 		public.POST("/user/login", controllers.LoginUser)

	// 		// Category
	// 		public.GET("/categories", controllers.GetCategories)
	// 		public.GET("/category/:id", controllers.GetCategoryById)

	// 		// Product
	// 		public.GET("/products", controllers.GetProducts)
	// 		public.GET("/product/:id", controllers.GetProductById)

	// 		// Newsletters
	// 		public.GET("/newsletters", controllers.GetNewsletters)
	// 		public.POST("/newsletter", controllers.PostNewsletter)
	// 	}

	// 	protected := apiFinal.Group("/").Use(middlewares.JWT)
	// 	{
	// 		// User
	// 		protected.GET("/users", controllers.GetUsers)
	// 		protected.PUT("/user/:id", controllers.PutUserById)
	// 		protected.DELETE("/user/:id", controllers.DeleteUserById)

	// 		// Category
	// 		protected.PUT("/category/:id", controllers.PutCategoryById)
	// 		protected.DELETE("/category/:id", controllers.DeleteCategoryById)
	// 		protected.POST("/category", controllers.PostCategory)

	// 		// Product
	// 		protected.PUT("/product/:id", controllers.PutProductById)
	// 		protected.DELETE("/product/:id", controllers.DeleteProductById)
	// 		protected.POST("/product", controllers.PostProduct)
	// 	}
	// }

	publicGroup := router.Group("/")
	{
		// Admin
		publicGroup.POST("/admin/login", public.LoginAdmin)

		// User
		publicGroup.GET("/user/:id", public.GetUserById)
		publicGroup.POST("/user", public.PostUser)
		publicGroup.POST("/user/login", public.LoginUser)

		// Wishlist
		publicGroup.GET("/user/:id/wishlist", public.GetUserWishlist)
		publicGroup.PUT("/user/:id/wishlist", public.PutUserWishlist)

		// Cart
		publicGroup.GET("/user/:id/cart", public.GetUserCart)
		// publicGroup.PUT("/user/:id/cart", public.PutUserCart)

		// publicGroup.GET("/wishlist/:id", public.GetWishlistById)

		// Category
		publicGroup.GET("/categories", public.GetCategories)
		publicGroup.GET("/category/:id", public.GetCategoryById)

		// Product
		publicGroup.GET("/products", public.GetProducts)
		publicGroup.GET("/product/:id", public.GetProductById)

		// Newsletters
		publicGroup.GET("/newsletters", public.GetNewsletters)
		publicGroup.POST("/newsletter", public.PostNewsletter)

		// Should be private under

		// User
		publicGroup.GET("/users", public.GetUsers)
		publicGroup.PUT("/user/:id", public.PutUserById)
		publicGroup.DELETE("/user/:id", public.DeleteUserById)

		// Category
		publicGroup.PUT("/category/:id", public.PutCategoryById)
		publicGroup.DELETE("/category/:id", public.DeleteCategoryById)
		publicGroup.POST("/category", public.PostCategory)

		// Product
		publicGroup.PUT("/product/:id", public.PutProductById)
		publicGroup.DELETE("/product/:id", public.DeleteProductById)
		publicGroup.POST("/product", public.PostProduct)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}

// TODO :
//  - Get wishlist with /user/:id/wishlist
//  - Middleware de test de base pour chaque table (CRUD)
