package router

import (
	"github.com/gin-gonic/gin"
	Customer "pizzaOrder/customer/api"
	Order "pizzaOrder/order/api"
	Pizza "pizzaOrder/pizza/api"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		pizza := main.Group("pizza")//, authentication.Auth()) // Tudo feito :)
		{
			pizza.GET("/menu", Pizza.Menu)
			pizza.POST("/sugestao-sabor", Pizza.NewFlavor)
		}

		order := main.Group("order")//, authentication.Auth()) // Tudo feito :)
		{
			order.GET("/hisotry", Order.History)
			order.GET("/ultimo-pedido", Order.LastOrder)
			order.POST("/novo-pedido", Order.NovoPedido)
		}

		customer := main.Group("customer") // Tudo feito :)
		{
			customer.POST("/signin", Customer.CreateCustomer)
			customer.POST("/login", Customer.Login)
		}
	}

	return router
}