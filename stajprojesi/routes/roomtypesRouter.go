package routes

import (
	"github.com/gin-gonic/gin"
	"stajprojesi/controllers"
)

func RoomTypeRoutes(api *gin.RouterGroup) {
	// General routes for room types
	api.GET("/", controllers.GetRoomsType)              // List all room types
	api.GET("/:type_id", controllers.GetRoomTypeByID)   // Get room type by ID
	api.POST("/", controllers.CreateRoomType)           // Create new room type (Admin only)
	api.PATCH("/:type_id", controllers.UpdateRoomType)  // Update room type (Admin only)
	api.DELETE("/:type_id", controllers.DeleteRoomType) // Delete room type (Admin only)
}
