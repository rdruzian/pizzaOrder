package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"pizzaOrder/router"
)

type Server struct {
	port string
	server *gin.Engine
}


func NewServer() Server {
	return Server {
		port: "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := router.ConfigRoutes(s.server)

	log.Fatal(router.Run(":" + s.port))
}
