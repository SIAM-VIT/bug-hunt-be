package main

import (
	"net/http"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/siam-vit/bughunt-be/internal/database"
	"github.com/siam-vit/bughunt-be/internal/routes"
	"github.com/siam-vit/bughunt-be/internal/utils"
)

func main() {
	utils.PrintSiamBanner()
	database.Connect()
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_custom}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"}` + "\n",
		CustomTimeFormat: "02/01/2006 15:04:05",
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	routes.TeamRoutes(e)
	routes.QuestionRoutes(e)
	routes.LeaderBoardRoutes(e)
	e.Start(":8080")
}
