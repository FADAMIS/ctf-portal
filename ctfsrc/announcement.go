package ctfsrc

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Fabucik/ctf-portal/authentication"
	"github.com/Fabucik/ctf-portal/entities"
	"github.com/gin-gonic/gin"
)

func CreateAnnouncement(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !authentication.IsAdmin(ctx, session) {
		return
	}

	var announcement entities.Announcement
	ctx.Bind(&announcement)

	if DoesAnnouncementExist(announcement) {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "Announcement with this ID already exists",
		})

		return
	}

	var announcements entities.Announcements
	announcementDb, _ := os.ReadFile("./database/announcements.json")
	json.Unmarshal(announcementDb, &announcements)

	announcements.Announcements = append(announcements.Announcements, announcement)
	writableJson, _ := json.MarshalIndent(announcements, "", "\t")
	os.WriteFile("./database/announcements.json", writableJson, 0600)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Announcement successfully created",
	})
}

func GetAnnouncements(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !authentication.IsValidSession(session) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "not logged in",
		})
		return
	}

	var announcements entities.Announcements
	announcementDb, _ := os.ReadFile("./database/announcements.json")
	json.Unmarshal(announcementDb, &announcements)

	ctx.JSON(http.StatusOK, announcements)
}

func DeleteAnnouncement(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !authentication.IsAdmin(ctx, session) {
		return
	}

	var announcementToDelete entities.Announcement
	ctx.Bind(&announcementToDelete)

	var announcements entities.Announcements
	announcementDb, _ := os.ReadFile("./database/announcements.json")
	json.Unmarshal(announcementDb, &announcements)

	for i := 0; i < len(announcements.Announcements); i++ {
		if announcementToDelete.ID == announcements.Announcements[i].ID {
			announcements.Announcements = append(announcements.Announcements[:i], announcements.Announcements[i+1:]...)
		}
	}

	writableJson, _ := json.MarshalIndent(&announcements, "", "\t")
	os.WriteFile("./database/announcements.json", writableJson, 0600)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted announcement",
	})
}

func DoesAnnouncementExist(announcement entities.Announcement) bool {
	var announcements entities.Announcements
	announcementDb, _ := os.ReadFile("./database/announcements.json")
	json.Unmarshal(announcementDb, &announcements)

	for i := 0; i < len(announcements.Announcements); i++ {
		if announcement.ID == announcements.Announcements[i].ID {
			return true
		}
	}

	return false
}
