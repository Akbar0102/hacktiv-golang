package routers

import (
	"assignment2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.GET("/orders", controllers.GetOrders)
	r.POST("/order", controllers.CreateOrder)
	r.PUT("/order/:orderId", controllers.UpdateOrderById)
	r.DELETE("/order/:orderId", controllers.DeleteOrderById)

	return r
}
