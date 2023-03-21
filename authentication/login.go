package authentication

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/Fabucik/ctf-portal/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Login(ctx *gin.Context) {
	var credentials entities.User
	ctx.Bind(&credentials)

	dbJson, _ := os.ReadFile("./database/users.json")
	var dbContent entities.Users
	json.Unmarshal(dbJson, &dbContent)

	for i := 0; i < len(dbContent.Users); i++ {
		// if credentials match, create session
		if dbContent.Users[i].Username == credentials.Username && dbContent.Users[i].Password == credentials.Password {
			session := createSessionCookie(credentials.Username)
			ctx.SetCookie("session", session, 6*60*60, "/", "localhost", false, true)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
			})
			return
		}
	}

	ctx.JSON(http.StatusUnauthorized, gin.H{
		"message": "Login failed",
	})
}

func createSessionCookie(username string) string {

	// create new session with all the parameters
	var session entities.Session

	id := uuid.New()
	now := time.Now()

	session.ID = id.String()
	session.Username = username
	session.ExpiresIn = now.Unix() + 6*60*60

	returnSession, _ := json.Marshal(session)

	// read session database, append newly created session and write it back
	var sessions entities.Sessions
	sessionsJson, _ := os.ReadFile("./database/session-cookies.json")
	json.Unmarshal(sessionsJson, &sessions)

	sessions.Sessions = append(sessions.Sessions, session)

	writableJson, _ := json.MarshalIndent(sessions, "", "\t")
	os.WriteFile("./database/session-cookies.json", writableJson, 0600)

	return string(returnSession)
}

func isExpired(sessionTime int64) bool {
	now := time.Now()

	return sessionTime < now.Unix()
}
