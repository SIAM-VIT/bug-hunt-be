package services

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/siam-vit/bughunt-be/internal/database"
	"github.com/siam-vit/bughunt-be/internal/models"
)

func CreateTeam(team models.Teams) error {
	db := database.DB.Db

	// ctx := context.Background()
	// database.RedisClient.Del(ctx, "teams_by_score")
	var exists bool
	err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM teams WHERE team_name = $1)`, team.TeamName).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("team with name '%s' already exists", team.TeamName)
	}

	_, err = db.Exec(`
        INSERT INTO teams (team_id, team_name, team_members, score)
        VALUES ($1, $2, $3, $4)`,
		uuid.New(), team.TeamName, pq.Array(team.TeamMembers), team.Score)

	if err != nil {
		return err
	}
	return nil

}

func GetAllTeams() ([]models.Teams, error) {
	db := database.DB.Db

	rows, err := db.Query(`SELECT team_id, team_name, team_members, score, time_remaining FROM teams ORDER BY score DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Teams

	for rows.Next() {
		var team models.Teams
		var members pq.StringArray

		err := rows.Scan(&team.TeamID, &team.TeamName, &members, &team.Score, &team.TimeRemaining)
		if err != nil {
			return nil, err
		}

		team.TeamMembers = []string(members)
		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

func ModifyTeam(team models.Teams) error {
	db := database.DB.Db

	query := `UPDATE teams SET `
	params := []interface{}{}
	paramCount := 1

	if team.TeamName != "" {
		query += fmt.Sprintf("team_name = $%d, ", paramCount)
		params = append(params, team.TeamName)
		paramCount++
	}
	if len(team.TeamMembers) > 0 {
		query += fmt.Sprintf("team_members = $%d, ", paramCount)
		params = append(params, pq.Array(team.TeamMembers))
		paramCount++
	}
	if team.Score != 0 {
		query += fmt.Sprintf("score = $%d, ", paramCount)
		params = append(params, team.Score)
		paramCount++
	}
	if team.TimeRemaining != 0 {
		query += fmt.Sprintf("time_remaining = $%d, ", paramCount)
		params = append(params, team.TimeRemaining)
		paramCount++
	}

	if len(params) == 0 {
		return errors.New("no fields provided for update")
	}
	query = query[:len(query)-2]

	query += fmt.Sprintf(" WHERE team_id = $%d", paramCount)
	params = append(params, team.TeamID)

	_, err := db.Exec(query, params...)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTeam(teamID uuid.UUID) error {
	db := database.DB.Db

	_, err := db.Exec(`DELETE FROM teams WHERE team_id = $1`, teamID)
	if err != nil {
		return err
	}
	return nil
}
