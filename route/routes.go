package route

import (
	"base-project-go/app/controllers"
	"base-project-go/middleware"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	name     = "name"
	username = "username"
	email    = "email"
	role     = "role"
)

func SetupRouter() *gin.Engine {
	// r := gin.Default()
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	r.MaxMultipartMemory = 8 << 20

	Route := r.Group("/api")
	{
		Route.GET("", controllers.Index)
		Route.POST("/login", controllers.Login)
		Route.POST("/register", controllers.Register)
		Route.POST("/logout", controllers.Logout)

		Route.GET("/ping", func(c *gin.Context) {
			c.String(200, "No Authorization JWT")
		})

		auth := Route.Group("/")
		auth.Use(middleware.Authz())
		auth.Use(AuthRequired)
		{
			auth.GET("/status", controllers.Status)
			option := auth.Group("/option")
			{
				option.GET("", controllers.IndexOption)
				option.POST("", controllers.CreateOption)
				option.GET("/:id", controllers.ShowOption)
				option.PUT("/:id", controllers.UpdateOption)
				option.DELETE("/:id", controllers.DeleteOption)
			}

			role := auth.Group("/role")
			{
				role.GET("", controllers.IndexRole)
				role.POST("", controllers.CreateRole)
				role.GET("/:id", controllers.ShowRole)
				role.PUT("/:id", controllers.UpdateRole)
				role.DELETE("/:id", controllers.DeleteRole)
			}

			user := auth.Group("/user")
			{
				user.GET("", controllers.IndexUser)
				user.POST("", controllers.CreateUser)
				user.GET("/:id", controllers.ShowUser)
				user.PUT("/:id", controllers.UpdateUser)
				user.DELETE("/:id", controllers.DeleteUser)
			}

			menu := auth.Group("/menu")
			{
				menu.GET("", controllers.IndexMenu)
				menu.POST("", controllers.CreateMenu)
				menu.GET("/:id", controllers.ShowMenu)
				menu.PUT("/:id", controllers.UpdateMenu)
				menu.DELETE("/:id", controllers.DeleteMenu)
			}

			usermenu := auth.Group("/user-menu")
			{
				usermenu.GET("", controllers.IndexUserMenu)
				usermenu.POST("", controllers.CreateUserMenu)
				usermenu.GET("/:id", controllers.ShowUserMenu)
				usermenu.PUT("/:id", controllers.UpdateUserMenu)
				usermenu.DELETE("/:id", controllers.DeleteUserMenu)
			}

			auth.POST("/upload", controllers.UploadFile)
			auth.POST("/uploads", controllers.UploadFile2)
			auth.POST("/delete-file/:id", controllers.DeleteFile)

		}

	}

	return r
}

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(name)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}
