package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pizzaOrder/database"
	Pizza "pizzaOrder/pizza"
	"strconv"
	"strings"
)

func Menu(c *gin.Context) {
	db := database.GetDatabase()
	var pizzas Pizza.Pizza
	result := db.Find(&pizzas)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("Não foi possível trazer o menu: %v", result.Error),
		})
		return
	}

	c.JSON(200, pizzas)
}

func NewFlavor(c *gin.Context) {
	flavorName := c.Param("flavorName")
	ingredientes := c.Param("ingredients")
	preco := c.Param("price")
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Preço vindo do json: %v", preco),
	})
	precoFloat, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("Preço precisa ser um valor numérico: %v", err),
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
			"error": fmt.Sprintf("Erro ao criar sua pizza: %v", result.Error),
		})
	}

	c.JSON(200, gin.H{
		"message": "Sua pizza foi montado com sucesso!",
	})
}