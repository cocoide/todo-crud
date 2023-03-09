package controller

import (
	"net/http"

	"github.com/cocoide/todo-crud/model"
	"github.com/labstack/echo/v4"
)

func CreateTask(c echo.Context) error {
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	model.DB.Create(&task)
	return c.JSON(http.StatusCreated, task)
}

func GetTasks(c echo.Context) error {
	Tasks := []model.Task{}
	model.DB.Find(&Tasks)
	return c.JSON(http.StatusOK, Tasks)
}

func GetTask(c echo.Context) error {
	Task := model.Task{}
	if err := c.Bind(&Task); err != nil {
		return err
	}
	model.DB.Take(&Task)
	return c.JSON(http.StatusOK, Task)
}

func UpdateTask(c echo.Context) error {
	Task := model.Task{}
	if err := c.Bind(&Task); err != nil {
		return err
	}
	model.DB.Save(&Task)
	return c.JSON(http.StatusOK, Task)
}