package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/siam-vit/bughunt-be/internal/controllers"
)

func LeaderBoardRoutes(e *echo.Echo) {
	r := e.Group("/leaderboard")
	r.POST("/addPoints/:id", controllers.AddPoints)
	r.PUT("/modifyPoints/:id", controllers.ModifyPoints)

	r.POST("/startTimer", controllers.StartTimer)
	r.GET("/getTimer", controllers.GetTimeLeft)

}
