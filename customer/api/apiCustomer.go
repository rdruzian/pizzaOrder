package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	login "pizzaOrder/Login"
	"pizzaOrder/auth"
	"pizzaOrder/customer"
	"pizzaOrder/database"
)

func CreateCustomer(c *gin.Context) {
	db := database.GetDatabase()
	var user customer.Customer

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível fazer o bind do JSON: " + err.Error(),
		})
		return
	}

	user.Password = auth.SHA256Encoder(user.Password)

	result := db.Create(&user).Error
	if result != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("Não foi possível salvar o Login do cliente: %v", result),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Cliente cadastrado com sucesso!",
	})
}

func Login(c *gin.Context){
	db := database.GetDatabase()

	var ld login.LoginData
	var l customer.Customer

	err := c.ShouldBindJSON(&ld)
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("Não foi possível pegar o json com dados de login %v", err),
		})
		return
	}

	dbError := db.Where("login = ?", ld.Login).First(&l).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível encontrar usuário",
		})
		return
	}

	if l.Password != auth.SHA256Encoder(ld.Senha) {
		c.JSON(401, gin.H{
			"error": "Senha inválida",
		})
		return
	}

	token, err := auth.NewJWTService().GeraToken(l.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}