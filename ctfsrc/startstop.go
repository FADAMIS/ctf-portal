package ctfsrc

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/database"
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

	var autoStart entities.AutomaticStart
	ctx.Bind(&autoStart)

	/*
		var startInfo entities.StartStopInfo
		startInfoDb, _ := os.ReadFile("./database/time.json")
		json.Unmarshal(startInfoDb, &startInfo)


		startInfo.Automatic = ctfTime
		startInfo.Automatic.IsValid = true
		startInfo.Manual.IsValid = false

		writableJson, _ := json.MarshalIndent(startInfo, "", "\t")
		os.WriteFile("./database/time.json", writableJson, 0600)*/

	var manualStart entities.ManualStart
	manualStart.IsValid = false
	autoStart.IsValid = true

	database.WriteAutoStart(database.GetOpenedDB(), autoStart)
	database.WriteManStart(database.GetOpenedDB(), manualStart)

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

	/*
		var ctfTime entities.StartStopInfo
		timeDb, _ := os.ReadFile("./database/time.json")
		json.Unmarshal(timeDb, &ctfTime)*/

	ctfTime := database.ReadAutoStart(database.GetOpenedDB())

	ctx.JSON(http.StatusOK, ctfTime)
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

	/*
		var startInfo entities.StartStopInfo
		startInfoDb, _ := os.ReadFile("./database/time.json")
		json.Unmarshal(startInfoDb, &startInfo)

		startInfo.Manual = manualStart
		startInfo.Manual.IsValid = true
		startInfo.Automatic.IsValid = false

		writableJson, _ := json.MarshalIndent(&startInfo, "", "\t")
		os.WriteFile("./database/time.json", writableJson, 0600)*/

	var autoStart entities.AutomaticStart
	autoStart.IsValid = false
	manualStart.IsValid = true

	database.WriteManStart(database.GetOpenedDB(), manualStart)
	database.WriteAutoStart(database.GetOpenedDB(), autoStart)

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Successfully set time",
		"isrunning": strconv.FormatBool(manualStart.IsStarted),
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

	/*
		var ctfTime entities.ManualStart
		startDb, _ := os.ReadFile("./database/time.json")
		json.Unmarshal(startDb, &ctfTime)*/

	manTime := database.ReadManStart(database.GetOpenedDB())

	ctx.JSON(http.StatusOK, manTime)
}

func IsCtfExpired() bool {
	/*
		var startInfo entities.StartStopInfo
		timeDb, _ := os.ReadFile("./database/time.json")
		json.Unmarshal(timeDb, &startInfo)*/

	autoStart := database.ReadAutoStart(database.GetOpenedDB())
	manStart := database.ReadManStart(database.GetOpenedDB())

	// if automatic start is used
	if autoStart.IsValid {
		now := time.Now()
		return autoStart.StopTime < now.Unix()
	}

	// if manual start is used
	return !manStart.IsStarted
}

func IsCtfStarted() bool {
	/*
		var startInfo entities.StartStopInfo
		timeDb, _ := os.ReadFile("./database/time.json")
		json.Unmarshal(timeDb, &startInfo)*/

	autoStart := database.ReadAutoStart(database.GetOpenedDB())
	manStart := database.ReadManStart(database.GetOpenedDB())

	// if automatic start is used
	if autoStart.IsValid {
		now := time.Now()
		return autoStart.StartTime < now.Unix()
	}

	// if manual start is used
	return manStart.IsStarted
}
