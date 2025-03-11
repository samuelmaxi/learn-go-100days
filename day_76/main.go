package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", SomeMiddleware(SomeHandler, SomeErrorHandler))

	e.Logger.Fatal(e.Start(":8080"))
}

func SomeHandler(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		map[string]any{"message": "Hello Word"},
	)
}

func SomeErrorHandler(c echo.Context) error {
	return c.JSON(
		http.StatusInternalServerError,
		map[string]any{"message": "Error"},
	)
}

func SomeMiddleware(next, stop echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		err := next(c)

		duration := time.Since(start)
		status := c.Response().Status
		method := c.Request().Method
		path := c.Request().URL.Path

		fmt.Printf("Method: %s,\nPath: %s,\nStatus: %d,\nDuration: %s\n", method, path, status, duration)

		return err
	}
}
