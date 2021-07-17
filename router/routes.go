package routes

import (
	"github.com/gin-gonic/gin"
	"restYT/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		pizza := main.Group("pizza")
		{
			pizza.GET("",)
			pizza.GET("",)
			pizza.GET("",)
			pizza.GET("",)
		}
		order := main.Group("order")
		{
			order.GET("",)
			order.GET("",)
			order.GET("",)
		}
		customer := main.Group("customer")
		{
			customer.POST("/", controllers.CreateUser)
			customer.POST("/login", controllers.Login)
		}
	}

	return router
}