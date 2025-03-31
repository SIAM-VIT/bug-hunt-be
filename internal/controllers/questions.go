package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/siam-vit/bughunt-be/internal/models"
	"github.com/siam-vit/bughunt-be/internal/services"
)

func CreateQuestion(c echo.Context) error {
	var question models.Question
	if err := c.Bind(&question); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to create question",
			"data":    err.Error(),
		})
	}

	err := services.CreateQuestion(question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create question",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully created question",
		"data":    question.Question,
	})
}

func GetQuestionsByDifficulty(c echo.Context) error {
	difficulty := c.QueryParam("difficulty")
	questions, err := services.GetQuestionsByDifficulty(difficulty)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch questions",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, questions)
}

func CreateTestCase(c echo.Context) error {
	var testCase models.TestCases
	if err := c.Bind(&testCase); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to create test case",
			"data":    err.Error(),
		})
	}

	err := services.CreateTestCase(testCase)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create test case",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully created test case",
		"data":    testCase.Input,
	})
}

func GetAllTestCases(c echo.Context) error {
	testCases, err := services.GetAllTestCases()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch test cases",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, testCases)
}
