package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/siam-vit/bughunt-be/internal/controllers"
)

func QuestionRoutes(e *echo.Echo) {
	r := e.Group("/questions")
	r.POST("/createQuestion", controllers.CreateQuestion)
	r.GET("/getQuestionsByDifficulty", controllers.GetQuestionsByDifficulty)
	r.POST("/createTestCase", controllers.CreateTestCase)
	r.GET("/getAllTestCases", controllers.GetAllTestCases)

}
