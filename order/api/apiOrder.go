package api

import (
	"github.com/gin-gonic/gin"
	"pizzaOrder/database"
	"pizzaOrder/order"
	"strconv"
	"strings"
	"time"
)

func History(c *gin.Context) {
	db := database.GetDatabase()
	idUser, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID do usuário precisa ser um valor numérico",
		})
	}

	var pedidos []order.Order
	err = db.Find(&pedidos).Where("customer=?", idUser).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível trazer o histórico de pedidos",
		})
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, pedidos)

}

func LastOrder(c *gin.Context){
	id := c.Param("customer")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	var p order.Order
	err = db.First(&p, idInt).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func NovoPedido(c *gin.Context){
	idUser, err := strconv.ParseInt(c.Param("costumer"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID do usuário deve ser numérico",
		})
		return
	}
	flavors := strings.Split(c.Param("flavor"), ",")
	date, err := time.Parse("2006-01-02 15:04", c.Param("dateTime"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Erro ao converter a data e hora",
		})
		return
	}
	preco, err := strconv.ParseFloat(c.Param("total"), 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Preço deve ser um valor numérico",
		})
		return
	}

	status, err := strconv.ParseInt(c.Param("status"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Status deve ser um valor numérico",
		})
		return
	}

	var sabores []order.Flavor
	for _, v := range(flavors){
		sabor := strings.Split(v, ";")
		if sabor != nil {
			sabores = append(sabores, order.Flavor{
				FlavorId: sabor[0],
				PizzaEdge: sabor[1],
				Units: sabor[2],
				Details: sabor[3],
				Price: sabor[4],
			})
		}
	}

	db := database.GetDatabase()
	pedido := order.Order{
		CustomerID: idUser,
		Flavor: sabores,
		Date: date,
		TotalPrice: preco,
		Status: status,
	}

	dbError := db.Create(&pedido)
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "Erro ao adicionar seu pedido.",
		})
		return
	}

	c.JSON(200, pedido)
}