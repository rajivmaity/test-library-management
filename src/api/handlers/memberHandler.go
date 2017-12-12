package handlers

import (
	"github.com/labstack/echo"
)

func LoginMember(c echo.Context) error {

	return c.JSON(200, map[string]string{
		"message": "you are in member group",
	})

}

func MemberLanding(c echo.Context) error {

	return c.JSON(200, map[string]string{
		"message": "you are in member group landing page",
	})

}
