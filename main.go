package main

import (
	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/ctfsrc"
	"github.com/Fabucik/ctf-portal/servehtml"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.LoadHTMLFiles("frontend/dist/index.html")
	server.Static("/assets", "./frontend/dist/assets")

	//authentication stuff
	server.POST("/register", authentication.Register)
	server.POST("/login", authentication.Login)

	// challenge upload
	server.POST("/upload", ctfsrc.CreateChallenge)

	// flag validation
	server.POST("/validate", ctfsrc.ValidateFlag)

	server.GET("/challenges", ctfsrc.GetChallenges)

	// html serving
	server.GET("/loginpage", servehtml.LoginHTML)

	server.Run(":8888")
}
