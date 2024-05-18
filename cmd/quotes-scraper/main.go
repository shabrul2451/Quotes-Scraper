package main

import (
	"context"
	"log"
	"os"
	"web_Scraping/internal/api"
	"web_Scraping/internal/scraper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/quotes", api.GetQuotesHandler(client))

	c := cron.New(cron.WithSeconds())
	_, err = c.AddFunc("@every 1h", func() {
		err := scraper.ScrapeAndStoreQuotes("https://www.goodreads.com/quotes", client)
		if err != nil {
			log.Printf("Error scraping and storing quotes: %v", err)
		} else {
			log.Println("Successfully scraped and stored quotes")
		}
	})
	if err != nil {
		log.Fatalf("Failed to create cron job: %v", err)
	}

	go scraper.ScrapeAndStoreQuotes("https://www.goodreads.com/quotes", client)

	// Start the cron scheduler
	c.Start()

	e.Logger.Fatal(e.Start(":8080"))
}
