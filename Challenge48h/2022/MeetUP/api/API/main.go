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

	router.LoadHTMLGlob("src/html/*")
	router.Static("/js", "src/js")

	// secure := router.Group("/", middlewares.GetJWT)

	router.GET("/", controllers.ShowIndexPage)
	router.GET("/event", controllers.ShowIndexPage)
	router.GET("/login", controllers.ShowLoginPage)
	router.POST("/login", controllers.ShowLoginPage)
	router.GET("/register", controllers.ShowRegisterPage)
	router.GET("/register/user", controllers.ShowRegisterUserPage)
	router.POST("/register/user", controllers.ShowRegisterUserPage)
	router.GET("/register/orga", controllers.ShowRegisterOrgaPage)
	router.POST("/register/orga", controllers.ShowRegisterOrgaPage)
	router.GET("/orga", controllers.ShowOrgaPage)
	router.GET("/orga/event", controllers.ShowCreateEventPage)
	router.POST("/orga/event", controllers.ShowCreateEventPage)
	router.GET("/event/:id", controllers.ShowEventPage)
	router.GET("/profile", controllers.ShowProfilPage)

	//orga

	//admin
	// router.GET("/admin/",controllers.data) à définir les datas
	router.GET("/admin/users", controllers.ShowAllUserPage)
	router.GET("/admin/orgas", controllers.ShowAllOrgasPage)
	router.GET("/admin/admins", controllers.ShowAllAdminPage)

	// error 404
	router.NoRoute(func(c *gin.Context) { controllers.ShowErrorPage(c, 404) })

	// router.Static("/images", "./images") // for static path

	// Set path of API
	api := router.Group("/api")

	// Public route
	public := api.Group("/")

	// Users CRUD - WIP
	public.GET("/user/:id", controllers.GetUserById)
	public.POST("/user/login", controllers.LoginUser)
	public.POST("/user", controllers.PostUser)

	// Orga
	public.GET("/orga/:id", controllers.GetOrgaById)
	public.POST("/orga/login", controllers.LoginOrga)
	public.GET("/orgas", controllers.GetOrgas)
	public.POST("/orga", controllers.PostOrga)

	// Category
	public.GET("/categories", controllers.GetCategories)
	public.GET("/category/:id", controllers.GetCategoryById)

	// Event
	public.GET("/", controllers.RootEvent)
	public.GET("/event/:id", controllers.GetEventById)

	// Admin
	public.POST("/admin/login", controllers.LoginAdmin)

	// private routes
	// private := api.Group("/", middlewares.JWT)
	private := api.Group("/")

	// User
	private.DELETE("/user/:id", controllers.DeleteUserById)
	private.GET("/users", controllers.GetUsers)

	// Organisateur CRUD - WIP
	private.DELETE("/orga/:id", controllers.DeleteOrgaById)

	// Admin CRUD - WIP
	private.GET("/admin/:id", controllers.GetAdminById)
	private.GET("/admins", controllers.GetAdmins)

	// Category CRUD - DONE (without security)
	private.PUT("/category/:id", controllers.PutCategoryById)
	private.DELETE("/category/:id", controllers.DeleteCategoryById)
	private.POST("/category", controllers.PostCategory)

	// Event CRUD - TODO
	private.POST("/event", controllers.PostEvent)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":80")
}
