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

		user := Route.Group("/user")
		{
			user.GET("/", controllers.IndexUser)
			user.POST("/", controllers.CreateUser)
			// user.POST("/", controllers.CreateOrUpdateUser)
			user.GET("/:id", controllers.ShowUser)
			user.PUT("/:id", controllers.UpdateUser)
			user.DELETE("/:id", controllers.DeleteUser)
		}

		role := Route.Group("/role")
		{
			role.GET("/", controllers.IndexRole)
			role.POST("/", controllers.CreateRole)
			role.GET("/:id", controllers.ShowRole)
			role.PUT("/:id", controllers.UpdateRole)
			role.DELETE("/:id", controllers.DeleteRole)
		}
	}

	// r.GET("/users", func(c *gin.Context) {
	// 	var result models.User
	// 	database.DB.Raw("SELECT * FROM users").Scan(&result)
	// 	c.JSON(http.StatusOK, result)
	// })
	// r.GET("/users/:id", func(c *gin.Context) {
	// 	var result models.User
	// 	ID := c.Param("id")
	// 	database.DB.Raw("SELECT * FROM users WHERE id = ?", ID).Scan(&result)
	// 	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": result})
	// })

	return r
}
