package ctfsrc

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/entities"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateChallenge(ctx *gin.Context) {
	var challenge entities.Challenge
	ctx.ShouldBindBodyWith(&challenge, binding.JSON)

	var session entities.Session
	ctx.ShouldBindBodyWith(&session, binding.JSON)

	isAdmin := authentication.IsAdmin(ctx, session)
	if !isAdmin {
		return
	}

	os.Mkdir("CTFCONTENTS/"+challenge.Name, 0777)
	os.Mkdir("CTFCONTENTS/"+challenge.Name+"/FILES", 0777)

	os.WriteFile("CTFCONTENTS/"+challenge.Name+"/FLAG.TXT", []byte(challenge.Flag), 0600)

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
