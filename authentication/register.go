package authentication

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Fabucik/ctf-portal/entities"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Register(ctx *gin.Context) {
	var newUser entities.User
	ctx.ShouldBindBodyWith(&newUser, binding.JSON)

	var session entities.Session
	ctx.ShouldBindBodyWith(&session, binding.JSON)

	isAdmin := IsAdmin(ctx, session)
	if !isAdmin {
		return
	}

	if doesUserExists(newUser.Username) || newUser.Username == "admin" {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "Username already exists",
		})

		return
	}

	// read user database, append new user and write back
	userJson, _ := os.ReadFile("./database/users.json")
	var users entities.Users
	json.Unmarshal(userJson, &users)

	users.Users = append(users.Users, newUser)

	writableJson, _ := json.MarshalIndent(users, "", "\t")
	os.WriteFile("./database/users.json", writableJson, 0600)

	// create entries in point database for the newly created user
	var userPoints entities.TeamPoints
	userPoints.Team = newUser.Username
	userPoints.PointAmount = 0

	var db entities.AllPoints
	pointDb, _ := os.ReadFile("./database/points.json")
	json.Unmarshal(pointDb, &db)

	db.Points = append(db.Points, userPoints)
	writableJson, _ = json.MarshalIndent(db, "", "\t")
	os.WriteFile("./database/points.json", writableJson, 0600)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register successful",
	})
}

func doesUserExists(username string) bool {
	dbJson, _ := os.ReadFile("./database/users.json")

	var users entities.Users
	json.Unmarshal(dbJson, &users)

	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Username == username {
			return true
		}
	}

	return false
}
