package api

import (
	"github.com/gin-gonic/gin"
	"pizzaOrder/database"
	Pizza "pizzaOrder/pizza"
	"strconv"
	"strings"
)

func Menu(c *gin.Context) {
	db := database.GetDatabase()
	var pizzas []Pizza.Pizza
	err := db.Find(&pizzas).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, pizzas)
}

func NewFlavor(c *gin.Context) {
	flavorName := c.Param("flavorName")
	ingredientes := c.Param("ingredients")
	preco := c.Param("price")
	precoFloat, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"eeror": "Preço precisa ser um valor numérico",
		})
	}

	db := database.GetDatabase()
	pizza :=  Pizza.Pizza{
		FlavorName: flavorName,
		Ingredients: strings.Split(ingredientes, ","),
		Price: precoFloat,
	}

	result := db.Create(&pizza)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Erro ao cria sua pizza",
		})
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Sua pizza foi montado com sucesso!",
	})
}