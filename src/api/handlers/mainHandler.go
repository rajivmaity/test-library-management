package handlers

import (
	"api/typestructs"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

//post author search to this handler
func SearchAuthors(c echo.Context) error {

	searchAuthorID := typestructs.SearchAuthorID{}

	err := c.Bind(&searchAuthorID)

	if err != nil {
		log.Printf("Failed to process searchAuthors %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	getAuthorID := typestructs.GetAuthorData{}

	if searchAuthorID.ID == "1234" {

		getAuthorID.ID = searchAuthorID.ID
		getAuthorID.Name = "Rajiv Maity"
		getAuthorID.BestBook = "How to meet a ideal women"
		getAuthorID.AliasName = "Raj is super awesome guy"

	} else {
		return c.JSON(http.StatusOK, map[string]string{
			"errormessage": "no authors found",
		})
	}

	return c.JSON(http.StatusOK, getAuthorID)

}

//post book search to this handler
func SearchBooks(c echo.Context) error {

	searchBookID := typestructs.SearchBookID{}

	err := c.Bind(&searchBookID)

	if err != nil {
		log.Printf("Failed to process searchBooks %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	getBookData := typestructs.GetBookData{}

	if searchBookID.ID == "samplebook" {

		getBookData.ID = searchBookID.ID
		getBookData.Name = "The game of thrones"
		getBookData.AuthorName = "HBO hahaha"
		getBookData.PublishDate = "1st Jan 1970"

	} else {
		return c.JSON(http.StatusOK, map[string]string{
			"errormessage": "book not found",
		})
	}

	return c.JSON(http.StatusOK, getBookData)

}
