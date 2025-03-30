# YouTube Video Fetcher API

## Project Overview

A Go-based API service that fetches and stores YouTube videos based on search queries. The service continuously fetches videos in the background & stores video metadata in a PostgreSQL database. Additionally it provides APIs to access the stored videos.

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/brijeshbhalodiya/fampay_youtube_fetcher.git
cd fampay_youtube_fetcher
```

### 2. Environment Setup

Create a `.env` file in the root directory with the following variables:

```env
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=youtube_fetcher
YOUTUBE_API_KEY=YOUTUBE_API_KEY
SEARCH_QUERY=ipl
FETCH_INTERVAL=10s
MAX_RESULTS=50
PORT=8080
```

Replace `YOUTUBE_API_KEY` with your actual YouTube Data API key.

### 3. Running Locally

#### Using Go directly:

```bash
# Download dependencies
go mod download

# Run the application
go run cmd/api/main.go
```

#### Using Docker Compose:

Update the .env file with the proper `MONGO_URI` with `mongodb://mongodb:27017`

```bash
# Build and run the application with MongoDB
docker-compose up --build
```

The API will be available at `http://localhost:8080`

## API Endpoints

### 1. Get Latest Videos

```http
GET /api/v1/videos?limit=10&offset=0
```

Query Parameters:

- `limit` (optional): Number of videos to return (default: 10)
- `offset` (optional): Number of videos to skip (default: 0)

Response:

```json
{
  "videos": [
    {
      "id": "...",
      "video_id": "...",
      "title": "...",
      "description": "...",
      "published_at": "2025-03-30T12:54:22Z",
      "default_thumbnail": "...",
      "created_at": "2025-03-30T12:57:47.302Z",
      "updated_at": "2025-03-30T12:57:47.302Z"
    }
  ]
}
```

### 2. Search Videos

```http
GET /api/v1/videos/search?q=query&limit=10&offset=0
```

Query Parameters:

- `q` (required): Search query
- `limit` (optional): Number of videos to return (default: 10)
- `offset` (optional): Number of videos to skip (default: 0)

Response format is the same as Get Latest Videos.
