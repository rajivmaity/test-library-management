package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func SetMemberMiddleware(g *echo.Group) {

	g.Use(checkCookie)

}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")

		if err != nil {

			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "cookie is deleted or dont have any cookie now")
			}

			log.Println(err)
			return err
		}

		if cookie.Value == "ypwupzFenj" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "you don't have right cookie")

	}

}
