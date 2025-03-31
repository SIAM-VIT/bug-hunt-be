package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"team_id"`
	Name      string    `json:"team_name"`
	// TeamMembers   []string  `json:"team_members"`
	Score         int       `json:"score"`
	TimeRemaining int       `json:"time_remaining"`
}
