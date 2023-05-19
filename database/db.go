package database

import (
	"github.com/Fabucik/ctf-portal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=fanda dbname=ctfportal port=5432 sslmode=disable TimeZone=Europe/Prague"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities.User{}, &entities.Announcement{}, &entities.Session{}, &entities.TeamPoints{}, &entities.AutomaticStart{}, &entities.ManualStart{})

	return db, nil
}

var DB, _ = InitDB()

func GetOpenedDB() *gorm.DB {
	return DB
}

func CreateUser(db *gorm.DB, user entities.User) {
	db.Create(&user)
}

func UpdateUser(db *gorm.DB, user entities.User) {
	db.Save(&user)
}

func ReadUsers(db *gorm.DB) entities.Users {
	var users []entities.User
	db.Find(&users)

	var allUsers entities.Users
	allUsers.Users = users
	return allUsers
}

func DeleteUser(db *gorm.DB, user entities.User) {
	db.Delete(user)
	db.Delete(&entities.Session{}, "username = ?", user.Username)
	db.Delete(&entities.TeamPoints{}, "team = ?", user.Username)
}

func CreateAnnouncement(db *gorm.DB, announcement entities.Announcement) {
	db.Create(&announcement)
}

func ReadAnnouncements(db *gorm.DB) entities.Announcements {
	var announcements []entities.Announcement

	db.Find(&announcements)

	var allAnnouncements entities.Announcements
	allAnnouncements.Announcements = announcements

	return allAnnouncements
}

func DeleteAnnouncement(db *gorm.DB, announcement entities.Announcement) {
	db.Delete(&announcement)
}

func WriteSession(db *gorm.DB, session entities.Session) {
	db.Create(&session)
}

func ReadSessions(db *gorm.DB) entities.Sessions {
	var sessions []entities.Session

	db.Find(&sessions)

	var allSessions entities.Sessions
	allSessions.Sessions = sessions
	return allSessions
}

func CreateTeamPoints(db *gorm.DB, teamPoints entities.TeamPoints) {
	db.Create(&teamPoints)
}

// solved will be string split by ;
func UpdatePoints(db *gorm.DB, teamPoints entities.TeamPoints) {
	db.Save(&teamPoints)
}

func ReadAllPoints(db *gorm.DB) entities.AllPoints {
	var points []entities.TeamPoints

	db.Find(&points)

	var allPoints entities.AllPoints
	allPoints.Points = points
	return allPoints
}

func ReadTeamPoints(db *gorm.DB, team string) entities.TeamPoints {
	var points entities.TeamPoints

	db.Last(&points, "team = ?", team)

	return points
}

func WriteAutoStart(db *gorm.DB, time entities.AutomaticStart) {
	db.Savee(&time)
}

func ReadAutoStart(db *gorm.DB) entities.AutomaticStart {
	var start entities.AutomaticStart

	db.Last(&start)

	return start
}

func WriteManStart(db *gorm.DB, time entities.ManualStart) {
	db.Save(&time)
}

func ReadManStart(db *gorm.DB) entities.ManualStart {
	var start entities.ManualStart

	db.Last(&start)

	return start
}
