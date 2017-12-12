package groups

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func SetMemberGroup(g *echo.Group) {

	g.GET("", handlers.MemberLanding)
	g.GET("/", handlers.MemberLanding)
	g.GET("/login", handlers.LoginMember)

}
