package main

import (
	"my-app/models"
	"my-app/views"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var todos []models.Todo
var idCounter int

func main() {
	e := echo.New()

	// Serve static files (e.g., Tailwind CSS)
	e.Static("/static", "static")

	e.GET("/", func(ctx echo.Context) error {
		return render(ctx, views.Index(todos))
	})
	e.PUT("/todos", func(ctx echo.Context) error {
		title := ctx.FormValue("title")
		idCounter++
		todos = append(todos, models.Todo{ID: idCounter, Title: title})
		return render(ctx, views.TodoList(todos))
	})
	e.PUT("/todos/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid ID")
		}
		title := c.FormValue("title")
		if title == "" {
			return c.String(http.StatusBadRequest, "Title is required")
		}
		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Title = title
				return render(c, views.TodoList(todos))
			}
		}
		return c.String(http.StatusNotFound, "Todo not found")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func render(ctx echo.Context, component templ.Component) error {
	ctx.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}
