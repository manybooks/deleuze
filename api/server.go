package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/interpret", handleGet)
	e.Start(":1111")
}

func handleGet(c echo.Context) error {
	text := c.QueryParam("text")
	return c.String(http.StatusOK, fmt.Sprintf("You said %v.", text))
}
