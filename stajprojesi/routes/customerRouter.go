package routes

import (
	"github.com/gin-gonic/gin"
	"stajprojesi/controllers"
)

func CustomerRoutes(api *gin.RouterGroup) {
	api.GET("/", controllers.GetCustomers)
	api.GET("/:customer_id", controllers.GetCustomer)
	api.POST("/signup", controllers.CustomerSignUp)
	api.PATCH("/:customer_id", controllers.UpdateCustomer)
	api.DELETE("/:customer_id", controllers.DeleteCustomer)
}
