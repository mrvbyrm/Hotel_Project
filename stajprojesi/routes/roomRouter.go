package routes

import (
	"github.com/gin-gonic/gin"
	"stajprojesi/controllers"
)

func RoomRoutes(api *gin.RouterGroup) {
	api.GET("/", controllers.GetRooms)
	api.GET("/:id", controllers.GetRoomByID)
	api.POST("/", controllers.CreateRoom)
	api.PATCH("/:id", controllers.UpdateRoom)
	api.DELETE("/:id", controllers.DeleteRoom)
}
