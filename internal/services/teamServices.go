package services

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/siam-vit/bughunt-be/internal/database"
	"github.com/siam-vit/bughunt-be/internal/models"
)

func CreateUser(user models.User) error {
	db := database.DB.Db

	var exists bool
	err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE name = $1)`, user.Name).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("user with name '%s' already exists", user.Name)
	}

	_, err = db.Exec(`
        INSERT INTO users (id, name, score, time_remaining)
        VALUES ($1, $2, $3, $4)`,
		uuid.New(), user.Name, user.Score, user.TimeRemaining)

	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]models.User, error) {
	db := database.DB.Db

	rows, err := db.Query(`SELECT id, name, score, time_remaining FROM users ORDER BY score DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Score, &user.TimeRemaining)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func ModifyUser(user models.User) error {
	db := database.DB.Db

	query := `UPDATE users SET `
	params := []interface{}{}
	paramCount := 1

	if user.Name != "" {
		query += fmt.Sprintf("name = $%d, ", paramCount)
		params = append(params, user.Name)
		paramCount++
	}
	if user.Score != 0 {
		query += fmt.Sprintf("score = $%d, ", paramCount)
		params = append(params, user.Score)
		paramCount++
	}
	if user.TimeRemaining != 0 {
		query += fmt.Sprintf("time_remaining = $%d, ", paramCount)
		params = append(params, user.TimeRemaining)
		paramCount++
	}

	if len(params) == 0 {
		return errors.New("no fields provided for update")
	}
	query = query[:len(query)-2]

	query += fmt.Sprintf(" WHERE id = $%d", paramCount)
	params = append(params, user.ID)

	_, err := db.Exec(query, params...)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(userID uuid.UUID) error {
	db := database.DB.Db

	_, err := db.Exec(`DELETE FROM users WHERE id = $1`, userID)
	if err != nil {
		return err
	}
	return nil
}
