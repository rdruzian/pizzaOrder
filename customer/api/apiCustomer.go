package api

import (
	"github.com/gin-gonic/gin"
	"pizzaOrder/auth"
	"pizzaOrder/customer"
	"pizzaOrder/database"
)

func CreateCustomer(c *gin.Context) {
	db := database.GetDatabase()
	var user customer.Customer
	var userAddress customer.Address
	var userLogin customer.LoginUser

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível fazer o bind do JSON: " + err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&userAddress)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível fazer o bind do JSON de endereço: " + err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&userLogin)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível fazer o bind do JSON com dados do login: " + err.Error(),
		})
		return
	}

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível salvar cliente",
		})
		return
	}

	userAddress.IDUser = user.ID
	userLogin.IDUser = user.ID

	result = db.Create(&userAddress)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível salvar o endereço do cliente",
		})
		return
	}

	userLogin.Password = auth.SHA256Encoder(userLogin.Password)
	result = db.Create(&userLogin)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível salvar o login do cliente",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Cliente cadastrado com sucesso!",
	})
}

func Login(c *gin.Context){
	db := database.GetDatabase()
	var l customer.LoginUser

	err := c.ShouldBindJSON(&l)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível fazer o bind do JSON: " + err.Error(),
		})
		return
	}

	var user customer.LoginUser
	dbError := db.Where("login = ?", user.User).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível encontrar usuário",
		})
		return
	}

	if user.Password != auth.SHA256Encoder(l.Password) {
		c.JSON(401, gin.H{
			"error": "Senha inválida",
		})
		return
	}

	token, err := auth.NewJWTService().GeraToken(user.IDUser)
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