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
