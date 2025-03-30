package main

import (
	"context"
	"fampay_youtube_fetcher/internal/config"
	"fampay_youtube_fetcher/internal/handler"
	"fampay_youtube_fetcher/internal/repository"
	"fampay_youtube_fetcher/internal/service"
	"fampay_youtube_fetcher/internal/worker"
	"fampay_youtube_fetcher/pkg/youtube"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.LoadConfig()

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(ctx)

	db := mongoClient.Database(cfg.MongoDBName)

	videoRepo := repository.NewVideoRepo(db)
	videoService := service.NewVideoService(videoRepo)
	youtubeClient := youtube.NewYoutubeClient(cfg.YouTubeAPIKey)
	videoHandler := handler.NewVideoHandler(videoService)

	worker := worker.NewYouTubeWorker(
		youtubeClient,
		videoService,
		cfg.SearchQuery,
		cfg.FetchInterval,
		cfg.MaxResults,
	)

	go worker.Start()

	app := fiber.New(fiber.Config{})
	app.Use(recover.New())
	app.Use(logger.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/videos", videoHandler.GetLatestVideos)
	v1.Get("/videos/search", videoHandler.SearchVideos)

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
