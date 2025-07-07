package routes

import (
	"github.com/gin-gonic/gin"
	"stajprojesi/controllers"
)

func UserRoutes(api *gin.RouterGroup) {
	// User-related endpoints
	api.GET("/", controllers.GetUsers)
	api.GET("/:user_id", controllers.GetUser)
	api.PATCH("/:user_id", controllers.UpdateUser)
	api.DELETE("/:user_id", controllers.DeleteUser)

	// Authentication-related endpoints
	api.POST("/signup", controllers.SignUp)
	api.POST("/login", controllers.Login)
	api.POST("/logout", controllers.Logout)
}
