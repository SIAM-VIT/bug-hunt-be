package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/siam-vit/bughunt-be/internal/models"
	"github.com/siam-vit/bughunt-be/internal/services"
)

func CreateUser(c echo.Context) error {

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to create user",
			"data":    err.Error(),
		})
	}

	err := services.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create user",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully created user",
		"data":    user.Name,
	})
}

func GetAllUsers(c echo.Context) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch users",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)
}

func UpdateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to update user",
			"data":    err.Error(),
		})
	}

	err := services.ModifyUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to update user",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully updated user",
		"data":    user.Name,
	})
}

func DeleteUser(c echo.Context) error {
	userID := c.Param("id")
	parseUserID, err := uuid.Parse(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to delete user",
			"data":    err.Error(),
		})
	}
	err = services.DeleteUser(parseUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete user",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully deleted user",
		"data":    userID,
	})
}
