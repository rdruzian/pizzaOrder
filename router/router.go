package router

import (
	"github.com/gin-gonic/gin"
	Customer "pizzaOrder/customer/api"
	Pizza "pizzaOrder/pizza/api"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		pizza := main.Group("pizzaOrder/pizza/api")
		{
			pizza.GET("/menu", Pizza.Menu)
			pizza.GET("/monte-sua", Pizza.MyFlavor)
			pizza.GET("/sugestao-sabor", Pizza.NewFlavor)
		}

		order := main.Group("order")
		{
			order.GET("/",)
			order.GET("/",)
			order.GET("/",)
		}

		customer := main.Group("customer")
		{
			customer.POST("/signin", )
			customer.POST("/login", Customer.Login)
			customer.GET("/loged", )
		}
	}

	return router
}