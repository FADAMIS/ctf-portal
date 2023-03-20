package ctfsrc

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/entities"
	"github.com/gin-gonic/gin"
)

func GetChallenges(ctx *gin.Context) {
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

	var allChallenges entities.ReturnChallenges

	challengeDirs, _ := os.ReadDir("CTFCONTENTS")
	allChallenges.Challenges = make([]entities.ReturnChallenge, len(challengeDirs))

	for i, entry := range challengeDirs {
		points, _ := os.ReadFile("CTFCONTENTS/" + entry.Name() + "/POINTS.TXT")
		description, _ := os.ReadFile("CTFCONTENTS/" + entry.Name() + "/DESCRIPTION.TXT")

		allChallenges.Challenges[i].Name = entry.Name()
		allChallenges.Challenges[i].Points, _ = strconv.Atoi(string(points))
		allChallenges.Challenges[i].Description = string(description)

		challengeFiles, _ := os.ReadDir("CTFCONTENTS/" + entry.Name() + "/FILES")
		for _, file := range challengeFiles {
			contents, _ := os.ReadFile("CTFCONTENTS/" + entry.Name() + "/FILES/" + file.Name())

			var fileInfo entities.ChallengeFile
			fileInfo.FileName = file.Name()
			fileInfo.Base64 = base64.StdEncoding.EncodeToString(contents)
			allChallenges.Challenges[i].Files = append(allChallenges.Challenges[i].Files, fileInfo)
		}
	}

	ctx.JSON(http.StatusOK, allChallenges)
}
