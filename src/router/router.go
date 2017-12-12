package router

import (
	"api/groups"
	"api/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {

	echoInstance := echo.New()

	/** CREATE GROUPS **/
	adminGroup := echoInstance.Group("/admin")
	memberGroup := echoInstance.Group("/member")

	/** Declare routes **/
	groups.SetMainGroup(echoInstance)
	groups.SetAdminGroup(adminGroup)
	groups.SetMemberGroup(memberGroup)

	/** Set middlewares **/
	middlewares.SetMainMiddleware(echoInstance)
	middlewares.SetAdminMiddleware(adminGroup)
	middlewares.SetMemberMiddleware(memberGroup)

	return echoInstance

}
