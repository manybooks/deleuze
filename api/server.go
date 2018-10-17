package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	client "github.com/manybooks/deleuze"
)

func main() {
	e := echo.New()
	e.GET("/interpret", handleGet)
	e.Start(":1111")
}

func handleGet(c echo.Context) error {
	text := c.QueryParam("text")
	ans := client.Reveal(text)
	return c.String(http.StatusOK, fmt.Sprintf("You said %v.", ans.Answer))
}
