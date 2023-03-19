package authentication

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Fabucik/ctf-portal/entities"
	"github.com/gin-gonic/gin"
)

func IsValidSession(session entities.Session) bool {
	dbJson, _ := os.ReadFile("./database/session-cookies.json")
	var sessions entities.Sessions
	json.Unmarshal(dbJson, &sessions)

	for i := 0; i < len(sessions.Sessions); i++ {
		if session == sessions.Sessions[i] {
			// if session exists BUT is expired, delete the session
			if isExpired(session.ExpiresIn) {
				sessions.Sessions = append(sessions.Sessions[:i], sessions.Sessions[i+1:]...)
				writableJson, _ := json.MarshalIndent(sessions, "", "\t")
				os.WriteFile("./database/session-cookies.json", writableJson, 0600)
				return false
			}

			return true
		}
	}

	return false
}

func IsAdmin(ctx *gin.Context, session entities.Session) bool {
	if !IsValidSession(session) || IsValidSession(session) && session.Username != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})

		return false
	}

	return true
}
