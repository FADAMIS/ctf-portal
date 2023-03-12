package main

import (
	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/servehtml"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.LoadHTMLFiles("frontend/dist/index.html")
	server.Static("/assets", "./frontend/dist/assets")

	server.POST("/register", authentication.Register)
	server.POST("/login", authentication.Login)

	server.GET("/loginpage", servehtml.LoginHTML)

	server.Run(":8888")
}
