package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/siam-vit/bughunt-be/internal/models"
	"github.com/siam-vit/bughunt-be/internal/services"
)

func CreateTeam(c echo.Context) error {

	var team models.Teams
	if err := c.Bind(&team); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to create team",
			"data":    err.Error(),
		})
	}

	err := services.CreateTeam(team)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create team",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully created team",
		"data":    team.TeamName,
	})

}

func GetAllTeams(c echo.Context) error {
	teams, err := services.GetAllTeams()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch teams",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, teams)
}

func UpdateTeam(c echo.Context) error {
	var team models.Teams
	if err := c.Bind(&team); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to update team",
			"data":    err.Error(),
		})
	}

	err := services.ModifyTeam(team)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to update team",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully updated team",
		"data":    team.TeamName,
	})
}

func DeleteTeam(c echo.Context) error {
	teamID := c.Param("id")
	parseTeamID, err := uuid.Parse(teamID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to delete team",
			"data":    err.Error(),
		})
	}
	err = services.DeleteTeam(parseTeamID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete team",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully deleted team",
		"data":    teamID,
	})

}
