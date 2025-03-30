package main

import (
	"context"
	"fampay_youtube_fetcher/internal/config"
	"fampay_youtube_fetcher/internal/repository"
	"fampay_youtube_fetcher/internal/service"
	"fampay_youtube_fetcher/internal/worker"
	"fampay_youtube_fetcher/pkg/youtube"
	"log"

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

	worker := worker.NewYouTubeWorker(
		youtubeClient,
		videoService,
		cfg.SearchQuery,
		cfg.FetchInterval,
		cfg.MaxResults,
	)

	go worker.Start()
}
