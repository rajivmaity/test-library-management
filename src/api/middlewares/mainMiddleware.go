package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetMainMiddleware(e *echo.Echo) {

	/** CREATE ALL MIDDLEWARES IN DIRECT CO-RELATION TO MAIN ECHO **/
	e.Use(serverHeader)
	/** SERVE ROOT TO THE CLIENT **/
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../static/forall",
	}))

	//e.Use(checkCookie)

}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "LibraryManagement/1.0")
		c.Response().Header().Set("ServerDesc", "Practise server for golang test")

		return next(c)
	}

}
