package main

import (
	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/ctfsrc"
	"github.com/Fabucik/ctf-portal/servehtml"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.LoadHTMLFiles("frontend/dist_collected/index.html", "frontend/dist_collected/admin.html", "frontend/dist_collected/CTF.html")
	server.Static("/assets", "./frontend/dist_collected/assets")

	//authentication stuff
	server.POST("/register", authentication.Register)
	server.POST("/login", authentication.Login)

	// challenge upload
	server.POST("/upload", ctfsrc.CreateChallenge)

	// flag validation
	server.POST("/validate", ctfsrc.ValidateFlag)

	// lists challenges
	server.GET("/challenges", ctfsrc.GetChallenges)

	// announcement stuff
	server.POST("/announcement", ctfsrc.CreateAnnouncement)
	server.GET("/announcement", ctfsrc.GetAnnouncements)
	server.DELETE("/announcement", ctfsrc.DeleteAnnouncement)

	// html serving
	server.GET("/", servehtml.LoginHTML)
	server.GET("/dashboard", servehtml.AdminHTML)
	server.GET("/ctf", servehtml.CtfHTML)

	server.Run(":8888")
}
