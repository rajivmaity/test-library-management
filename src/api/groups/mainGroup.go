package groups

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func SetMainGroup(e *echo.Echo) {

	/** DELARE MAIN END-POINTS **/
	e.POST("/search-authors", handlers.SearchAuthors)
	e.POST("/search-books", handlers.SearchBooks)

}
