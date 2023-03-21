package ctfsrc

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/entities"
	"github.com/gin-gonic/gin"
)

func GetTeams(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !authentication.IsAdmin(ctx, session) {
		return
	}

	var users entities.Users
	teamDb, _ := os.ReadFile("./database/users.json")
	json.Unmarshal(teamDb, &users)

	ctx.JSON(http.StatusOK, users)
}

func DeleteTeam(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !authentication.IsAdmin(ctx, session) {
		return
	}

	var userToDelete entities.User
	ctx.Bind(&userToDelete)

	var users entities.Users
	userDb, _ := os.ReadFile("./database/users.json")
	json.Unmarshal(userDb, &users)

	var points entities.AllPoints
	pointDb, _ := os.ReadFile("./database/points.json")
	json.Unmarshal(pointDb, &points)

	var sessions entities.Sessions
	sessionDb, _ := os.ReadFile("./database/session-cookies.json")
	json.Unmarshal(sessionDb, &sessions)

	check := 0
	for i := 0; i < len(users.Users); i++ {
		// cannot delete admin user
		if userToDelete.Username == "admin" {
			ctx.JSON(http.StatusForbidden, gin.H{
				"message": "cannot delete admin",
			})
		}

		if userToDelete.Username == users.Users[i].Username {
			users.Users = append(users.Users[:i], users.Users[i+1:]...)
			check++
		}
	}

	if check == 0 {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "User does not exist",
		})

		return
	}

	for i := 0; i < len(points.Points); i++ {
		if userToDelete.Username == points.Points[i].Team {
			points.Points = append(points.Points[:i], points.Points[i+1:]...)
		}
	}

	for i := 0; i < len(sessions.Sessions); i++ {
		if userToDelete.Username == sessions.Sessions[i].Username {
			sessions.Sessions = append(sessions.Sessions[:i], sessions.Sessions[i+1:]...)
		}
	}

	writableUserJson, _ := json.MarshalIndent(&users, "", "\t")
	os.WriteFile("./database/users.json", writableUserJson, 0600)

	writablePointJson, _ := json.MarshalIndent(&points, "", "\t")
	os.WriteFile("./database/points.json", writablePointJson, 0600)

	writableSessionJson, _ := json.MarshalIndent(&sessions, "", "\t")
	os.WriteFile("./database/session-cookies.json", writableSessionJson, 0600)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
