package main

import (
	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/ctfsrc"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/gzip"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.Use(gzip.Gzip(gzip.DefaultCompression))

	server.LoadHTMLFiles("frontend/dist_collected/index.html", "frontend/dist_collected/admin.html", "frontend/dist_collected/CTF.html")
	server.Static("/assets", "./frontend/dist_collected/assets")

	//authentication stuff
	server.POST("/api/register", authentication.Register)
	server.POST("/api/login", authentication.Login)

	// challenge upload
	server.POST("/api/upload", ctfsrc.CreateChallenge)

	// flag validation
	server.POST("/api/validate", ctfsrc.ValidateFlag)

	// lists challenges
	server.GET("/api/challenges", ctfsrc.GetChallenges)

	// team management
	server.GET("/api/teams", ctfsrc.GetTeams)
	server.DELETE("/api/teams", ctfsrc.DeleteTeam)

	// returns json with all teams and their points
	server.GET("/api/points", ctfsrc.GetAllPoints)

	// ctf start and stop
	server.POST("/api/timedstart", ctfsrc.SetTime)
	server.GET("/api/timedstart", ctfsrc.GetTime)
	server.POST("/api/manualstart", ctfsrc.SetManualStartStop)
	server.GET("/api/manualstart", ctfsrc.GetManualStartStop)

	// announcement stuff
	server.POST("/api/announcement", ctfsrc.CreateAnnouncement)
	server.GET("/api/announcement", ctfsrc.GetAnnouncements)
	server.DELETE("/api/announcement", ctfsrc.DeleteAnnouncement)

	// importing and exporting CTF challenges
	server.GET("/api/backup", ctfsrc.Export)
	server.POST("/api/backup", ctfsrc.Import)

	server.Run(":8888")
}
