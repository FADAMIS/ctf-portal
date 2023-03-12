package authentication

import (
	"net/http"

	"github.com/Fabucik/ctf-portal/entities"
	"github.com/gin-gonic/gin"
)

func IsAdmin(ctx *gin.Context, session entities.Session) bool {
	if !IsValidSession(session) || IsValidSession(session) && session.Username != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})

		return false
	}

	return true
}
