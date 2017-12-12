package groups

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func SetAdminGroup(e *echo.Group) {

	/** DELARE ADMIN-GROUP END-POINTS **/
	e.GET("", handlers.AdminHome)
	e.GET("/", handlers.AdminHome)
	e.GET("/main", handlers.AdminMain)

}
