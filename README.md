# Quotes Scraper

## Overview

This project scrapes quotes from a website (e.g., [Goodreads](https://www.goodreads.com/quotes)) and stores them in a MongoDB database. It provides an API to retrieve the stored quotes. The scraper runs initially and then every hour using a cron job.

## Project Structure

```plaintext
.
├── cmd
│   └── quotes-scraper
│       └── main.go
├── internal
│   ├── api
│   │   └── handler.go
│   ├── scraper
│   │   └── scraper.go
│   └── storage
│       └── mongo.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```
## Features
- Scrapes quotes from a specified website.
- Stores quotes in a MongoDB database.
- Provides an API to retrieve quotes.
- Runs a cron job to scrape quotes every hour.

## Building and Running the Project
#### Build and Run with Docker Compose 
1. To build and run the project, use Docker Compose:
```cmd
docker compose up --build
```

## API Endpoints
- GET /quotes 
- Retrieves all stored quotes.
```cmd
curl http://localhost:8080/quotes
```