package routes

import (
	"github.com/gin-gonic/gin"
	"stajprojesi/controllers"
)

func ReservationRoot(api *gin.RouterGroup) {
	api.GET("/", controllers.GetReservations)
	api.POST("/", controllers.CreateReservation)
	api.GET("/:id", controllers.GetReservationByID)
	api.PATCH("/:id", controllers.UpdateReservation)
	api.DELETE("/:id", controllers.DeleteReservation)
}
