package router

import (
	"github.com/gin-gonic/gin"
	"pizzaOrder/auth/authentication"
	Customer "pizzaOrder/customer/api"
	Order "pizzaOrder/order/api"
	Pizza "pizzaOrder/pizza/api"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		pizza := main.Group("pizzaOrder/pizza/api", authentication.Auth())
		{
			pizza.GET("/menu", Pizza.Menu)
			pizza.GET("/monte-sua", Pizza.MyFlavor)
			pizza.GET("/sugestao-sabor", Pizza.NewFlavor)
		}

		order := main.Group("order", authentication.Auth())
		{
			order.GET("/", Order.History)
			order.GET("/", Order.LastOrder)
		}

		customer := main.Group("customer")
		{
			customer.POST("/signin", Customer.CreateCustomer)
			customer.POST("/login", Customer.Login)
		}
	}

	return router
}