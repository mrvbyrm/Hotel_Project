package routes

import (
	"github.com/gin-gonic/gin"
	"stajprojesi/controllers"
)

func PaymentRoutes(api *gin.RouterGroup) {
	api.GET("/", controllers.GetPayment) // Tüm ödemeleri listelemek için eklenebilir
	api.GET("/:id", controllers.GetPaymentByID)
	api.POST("/", controllers.CreatePayment)
	api.PATCH("/:id/refund", controllers.RefundPayment)
}
