package ctfsrc

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/entities"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateChallenge(ctx *gin.Context) {
	var challenge entities.Challenge
	ctx.ShouldBindBodyWith(&challenge, binding.JSON)

	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	isAdmin := authentication.IsAdmin(ctx, session)
	if !isAdmin {
		return
	}

	if strings.Contains(challenge.Name, ";") {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Cannot use ';' in challenge name",
		})

		return
	}

	entries, _ := os.ReadDir("CTFCONTENTS")
	for _, e := range entries {
		if challenge.Name == e.Name() {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "challenge's name is taken",
			})

			return
		}
	}

	// all ctf challenges are stored under CTFCONTENTS directory as other directories
	os.Mkdir("CTFCONTENTS/"+challenge.Name, 0777)
	os.Mkdir("CTFCONTENTS/"+challenge.Name+"/FILES", 0777)

	os.WriteFile("CTFCONTENTS/"+challenge.Name+"/FLAG.TXT", []byte(challenge.Flag), 0600)
	os.WriteFile("CTFCONTENTS/"+challenge.Name+"/POINTS.TXT", []byte(strconv.Itoa(challenge.Points)), 0600)
	os.WriteFile("CTFCONTENTS/"+challenge.Name+"/DESCRIPTION.TXT", []byte(challenge.Description), 0600)
	os.WriteFile("CTFCONTENTS/"+challenge.Name+"/COUNTRY.TXT", []byte(challenge.CountryCode), 0600)

	// write each file to FILES dir
	for i := 0; i < len(challenge.Files); i++ {
		contents, err := base64.StdEncoding.DecodeString(challenge.Files[i].Base64)
		if err != nil {
			fmt.Println(err)
		}

		f, _ := os.Create("CTFCONTENTS/" + challenge.Name + "/FILES/" + challenge.Files[i].FileName)
		defer f.Close()

		f.Write(contents)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Upload successful",
	})
}
