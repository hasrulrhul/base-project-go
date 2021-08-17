package route

import (
	"base-project-go/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	Route := r.Group("/api")
	{
		Route.GET("/", controllers.Index)
		Route.POST("/foo", controllers.IndexPost)

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
	}

	return r
}
