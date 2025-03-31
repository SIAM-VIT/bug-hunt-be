package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/siam-vit/bughunt-be/internal/controllers"
)

func TeamRoutes(e *echo.Echo) {
	r := e.Group("/teams")

	r.POST("/createTeam", controllers.CreateTeam)
	r.GET("/getAllTeams", controllers.GetAllTeams)
	r.PUT("/updateTeam", controllers.UpdateTeam)
	r.DELETE("/deleteTeam/:id", controllers.DeleteTeam)
}
