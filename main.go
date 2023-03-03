package main

import (
	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/login", authentication.Register)
	server.GET("/login", authentication.Login)

	server.Run(":8888")
}
