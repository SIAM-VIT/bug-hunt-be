package services

import (
	"github.com/google/uuid"
	"github.com/siam-vit/bughunt-be/internal/database"
)

func AddPoints(points uint, id uuid.UUID) error {
	db := database.DB.Db
	_, err := db.Exec(`UPDATE teams SET score=score+$1 WHERE team_id=$2`, points, id)
	if err != nil {
		return err
	}
	return nil
}

func ModifyPoints(points uint, id uuid.UUID) error {
	db := database.DB.Db
	_, err := db.Exec(`UPDATE teams SET score=$1 WHERE team_id=$2`, points, id)
	if err != nil {
		return err
	}
	return nil
}
