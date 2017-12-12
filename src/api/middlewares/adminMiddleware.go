package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetAdminMiddleware(g *echo.Group) {

	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "Admin Log: " + `[${time_rfc3339}] 	${status} 	${method} 	${host}${path} 	${latency_human}` + "\n",
	}))

	/** SERVE ADMIN ROOT TO THE CLIENT **/
	g.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../static/admin",
	}))

}
