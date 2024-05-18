package api

import (
	"net/http"
	"web_Scraping/internal/storage"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetQuotesHandler(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		quotes, err := storage.GetQuotes(client)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch quotes from database"})
		}
		return c.JSON(http.StatusOK, quotes)
	}
}
