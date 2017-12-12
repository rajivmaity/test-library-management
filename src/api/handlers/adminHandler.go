package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func AdminMain(c echo.Context) error {
	return c.String(http.StatusOK, " this admin handler")
}

func AdminHome(c echo.Context) error {
	return c.String(http.StatusOK, " this admin home handler")
}
