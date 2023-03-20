package ctfsrc

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/entities"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ValidateFlag(ctx *gin.Context) {
	var flag entities.Flag
	ctx.ShouldBindBodyWith(&flag, binding.JSON)

	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !authentication.IsValidSession(session) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not logged in",
		})

		return
	}

	if !IsCtfStarted() || IsCtfExpired() {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "CTF is not yet started",
		})

		return
	}

	// return not found if challenge (directory) doesnt exist
	_, err := os.ReadDir("CTFCONTENTS/" + flag.Challenge)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Challenge not found",
		})
	}

	// read flag and compare it with input
	serverFlag, _ := os.ReadFile("CTFCONTENTS/" + flag.Challenge + "/FLAG.TXT")
	if string(serverFlag) != flag.Value {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "Wrong flag",
		})
	}

	// assign points
	points, _ := os.ReadFile("CTFCONTENTS/" + flag.Challenge + "/POINTS.TXT")
	pointsInt, _ := strconv.Atoi(string(points))

	// if flag was already answered, return status forbidden
	isSuccess := AssignPoints(pointsInt, session.Username, flag.Challenge)
	if !isSuccess {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Already answered",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Correct",
	})
}
