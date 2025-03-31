package models

import "github.com/google/uuid"

type Teams struct {
	TeamID        uuid.UUID `json:"team_id"`
	TeamName      string    `json:"team_name"`
	TeamMembers   []string  `json:"team_members"`
	Score         int       `json:"score"`
	TimeRemaining int       `json:"time_remaining"`
}
