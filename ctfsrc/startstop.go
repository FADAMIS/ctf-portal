package ctfsrc

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/entities"
	"github.com/gin-gonic/gin"
)

func SetTime(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !authentication.IsAdmin(ctx, session) {
		return
	}

	var ctfTime entities.AutomaticStart
	ctx.Bind(&ctfTime)

	var startInfo entities.StartStopInfo
	startInfoDb, _ := os.ReadFile("./database/time.json")
	json.Unmarshal(startInfoDb, &startInfo)

	startInfo.Automatic = ctfTime
	startInfo.Automatic.IsValid = true
	startInfo.Manual.IsValid = false

	writableJson, _ := json.MarshalIndent(startInfo, "", "\t")
	os.WriteFile("./database/time.json", writableJson, 0600)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully set time",
	})
}

func GetTime(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !authentication.IsValidSession(session) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not logged in",
		})

		return
	}

	var ctfTime entities.StartStopInfo
	timeDb, _ := os.ReadFile("./database/time.json")
	json.Unmarshal(timeDb, &ctfTime)

	ctx.JSON(http.StatusOK, ctfTime.Automatic)
}

func SetManualStartStop(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !authentication.IsAdmin(ctx, session) {
		return
	}

	var manualStart entities.ManualStart
	ctx.Bind(&manualStart)

	var startInfo entities.StartStopInfo
	startInfoDb, _ := os.ReadFile("./database/time.json")
	json.Unmarshal(startInfoDb, &startInfo)

	startInfo.Manual = manualStart
	startInfo.Manual.IsValid = true
	startInfo.Automatic.IsValid = false

	writableJson, _ := json.MarshalIndent(&startInfo, "", "\t")
	os.WriteFile("./database/time.json", writableJson, 0600)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "CTF running status is " + strconv.FormatBool(manualStart.IsStarted),
	})
}

func GetManualStartStop(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !authentication.IsValidSession(session) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not logged in",
		})

		return
	}

	var ctfTime entities.ManualStart
	startDb, _ := os.ReadFile("./database/time.json")
	json.Unmarshal(startDb, &ctfTime)

	ctx.JSON(http.StatusOK, ctfTime)
}

func IsCtfExpired() bool {
	var startInfo entities.StartStopInfo
	timeDb, _ := os.ReadFile("./database/time.json")
	json.Unmarshal(timeDb, &startInfo)

	// if automatic start is used
	if startInfo.Automatic.IsValid {
		now := time.Now()
		return startInfo.Automatic.StopTime < now.Unix()
	}

	// if manual start is used
	return !startInfo.Manual.IsStarted
}

func IsCtfStarted() bool {
	var startInfo entities.StartStopInfo
	timeDb, _ := os.ReadFile("./database/time.json")
	json.Unmarshal(timeDb, &startInfo)

	// if automatic start is used
	if startInfo.Automatic.IsValid {
		now := time.Now()
		return startInfo.Automatic.StartTime < now.Unix()
	}

	// if manual start is used
	return startInfo.Manual.IsStarted
}
