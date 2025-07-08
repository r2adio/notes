package main

import (
	"my-app/models"
	"my-app/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var todos []models.Todo
var idCounter int

func main() {
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return render(ctx, views.Index(todos))
	})
	e.PUT("/todos", func(ctx echo.Context) error {
		title := ctx.FormValue("title")
		idCounter++
		todos = append(todos, models.Todo{ID: idCounter, Title: title})
		return render(ctx, views.TodoList(todos))
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func render(ctx echo.Context, component templ.Component) error {
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}
