package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/siam-vit/bughunt-be/internal/controllers"
)

func TeamRoutes(e *echo.Echo) {
	r := e.Group("/teams")

	r.POST("/createTeam", controllers.CreateUser)
	r.GET("/getAllTeams", controllers.GetAllUsers)
	r.PUT("/updateTeam", controllers.UpdateUser)
	r.DELETE("/deleteTeam/:id", controllers.DeleteUser)
}
