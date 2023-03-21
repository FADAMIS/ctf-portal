package ctfsrc

import (
	"archive/zip"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/entities"
	"github.com/gin-gonic/gin"
)

func Export(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !authentication.IsAdmin(ctx, session) {
		return
	}

	fileName, err := Zip()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "An error occured while exporting the CTF",
		})

		return
	}

	zipFile, _ := os.ReadFile("./backup/" + fileName)
	encoded := base64.StdEncoding.EncodeToString(zipFile)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"base64":  encoded,
	})
}

func Import(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !authentication.IsAdmin(ctx, session) {
		return
	}

	var encoded entities.Backup
	ctx.Bind(&encoded)

	decoded, _ := base64.StdEncoding.DecodeString(encoded.Base64)
	os.WriteFile("./backup/import.zip", decoded, 0600)

	ctfcontents, _ := os.ReadDir("./CTFCONTENTS")
	for _, entry := range ctfcontents {
		if entry.Name() == ".gitignore" {
			continue
		}

		os.RemoveAll("./CTFCONTENTS/" + entry.Name())
	}

	err := Unzip("./backup/import.zip")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "An error occured while importing the CTF",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully imported CTF",
	})
}

func Unzip(src string) error {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}

	defer reader.Close()

	extract := func(file *zip.File, dst string) error {
		rc, err := file.Open()
		if err != nil {
			return err
		}

		defer rc.Close()

		path := filepath.Join(dst, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(path, 0777)
		} else {
			os.MkdirAll(filepath.Dir(path), 0777)

			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
			if err != nil {
				return err
			}

			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}

		return nil
	}

	for _, f := range reader.File {
		err := extract(f, "./CTFCONTENTS")
		if err != nil {
			return err
		}
	}

	return nil
}

func Zip() (string, error) {
	fileName := "backup-" + time.Now().Format("02_01_2006__15_04_05") + ".zip"
	dstFile, err := os.Create("./backup/" + fileName)
	if err != nil {
		return "", err
	}

	zipped := zip.NewWriter(dstFile)
	err = filepath.Walk("./CTFCONTENTS", func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(filePath, "CTFCONTENTS")
		zipFile, err := zipped.Create(relPath)
		if err != nil {
			return err
		}

		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}

		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	zipped.Close()

	return fileName, nil
}
