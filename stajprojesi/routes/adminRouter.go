package routes

import (
	"github.com/gin-gonic/gin"
	"stajprojesi/controllers"
)

func AdminRoutes(api *gin.RouterGroup) {
	api.GET("/", controllers.GetAdmins)
	api.GET("/:id", controllers.GetAdmin)
	api.POST("/signup", controllers.AdminSignUp)
	api.POST("/login", controllers.AdminLogin)
	api.PATCH("/:id", controllers.UpdateAdmin)
	api.DELETE("/:id", controllers.DeleteAdmin)
}
