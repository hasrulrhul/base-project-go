package route

import (
	"base-project-go/app/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

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
		Route.POST("/foo", controllers.IndexPost)
		Route.GET("/halo", controllers.Hello)

		Route.POST("/login", controllers.Login)
		Route.POST("/Register", controllers.Register)

		option := Route.Group("/option")
		{
			option.GET("", controllers.IndexOption)
			option.POST("", controllers.CreateOption)
			option.GET("/:id", controllers.ShowOption)
			option.PUT("/:id", controllers.UpdateOption)
			option.DELETE("/:id", controllers.DeleteOption)
		}

		role := Route.Group("/role")
		{
			role.GET("", controllers.IndexRole)
			role.POST("", controllers.CreateRole)
			role.GET("/:id", controllers.ShowRole)
			role.PUT("/:id", controllers.UpdateRole)
			role.DELETE("/:id", controllers.DeleteRole)
		}

		user := Route.Group("/user")
		{
			user.GET("", controllers.IndexUser)
			user.POST("", controllers.CreateUser)
			user.GET("/:id", controllers.ShowUser)
			user.PUT("/:id", controllers.UpdateUser)
			user.DELETE("/:id", controllers.DeleteUser)
		}

		menu := Route.Group("/menu")
		{
			menu.GET("", controllers.IndexMenu)
			menu.POST("", controllers.CreateMenu)
			menu.GET("/:id", controllers.ShowMenu)
			menu.PUT("/:id", controllers.UpdateMenu)
			menu.DELETE("/:id", controllers.DeleteMenu)
		}

		usermenu := Route.Group("/user-menu")
		{
			usermenu.GET("", controllers.IndexUserMenu)
			usermenu.POST("", controllers.CreateUserMenu)
			usermenu.GET("/:id", controllers.ShowUserMenu)
			usermenu.PUT("/:id", controllers.UpdateUserMenu)
			usermenu.DELETE("/:id", controllers.DeleteUserMenu)
		}

		Route.POST("/upload", controllers.UploadFile)
		Route.POST("/uploads", controllers.UploadFile2)
		Route.POST("/delete-file/:id", controllers.DeleteFile)
	}

	return r
}
