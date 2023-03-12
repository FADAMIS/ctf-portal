package ctfsrc

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Fabucik/ctf-portal/entities"
)

func AssignPoints(points int, team string, challenge string) bool {
	var db entities.AllPoints
	dbJson, _ := os.ReadFile("./database/points.json")
	json.Unmarshal(dbJson, &db)

	var newPoints entities.TeamPoints
	newPoints.Team = team
	for i := 0; i < len(db.Points); i++ {
		if newPoints.Team == db.Points[i].Team {
			// if team already answered the flag return from the function
			for j := 0; j < len(db.Points[i].Solved); j++ {
				fmt.Println(j)
				if db.Points[i].Solved[j] == challenge {
					return false
				}
			}

			// update points and delete old entry
			newPoints.PointAmount = db.Points[i].PointAmount + points
			db.Points = append(db.Points[:i], db.Points[i+1:]...)
		}
	}
	newPoints.Solved = append(newPoints.Solved, challenge)

	db.Points = append(db.Points, newPoints)

	writableJson, _ := json.MarshalIndent(db, "", "\t")
	os.WriteFile("./database/points.json", writableJson, 0600)

	return true
}
