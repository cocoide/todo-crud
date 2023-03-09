package main

import (
	"net/http"

	"github.com/cocoide/todo-crud/controller"
	"github.com/cocoide/todo-crud/model"
	"github.com/labstack/echo/v4"
)

func connect(c echo.Context) error {
	db, _ := model.DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB接続失敗しました")
	} else {
		return c.String(http.StatusOK, "DB接続しました")
	}
}

func main() {
	e := echo.New()
	db, _ := model.DB.DB()
	defer db.Close()
	e.POST("/tasks", controller.CreateTask)
	e.GET("/tasks", controller.GetTask) 
	e.GET("/tasks/:id", controller.GetTask)
	e.PUT("/tasks/:id", controller.UpdateTask)
	e.Logger.Fatal(e.Start(":8080"))
}