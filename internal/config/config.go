package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI      string
	MongoDBName   string
	YouTubeAPIKey string
	SearchQuery   string
	FetchInterval time.Duration
	MaxResults    int
	Port          string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDBName:   getEnv("MONGO_DB_NAME", "youtube_fetcher"),
		YouTubeAPIKey: getEnv("YOUTUBE_API_KEY", ""),
		SearchQuery:   getEnv("SEARCH_QUERY", "golang programming"),
		FetchInterval: getDurationEnv("FETCH_INTERVAL", 10*time.Second),
		MaxResults:    getIntEnv("MAX_RESULTS", 50),
		Port:          getEnv("PORT", "8080"),
	}

}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
