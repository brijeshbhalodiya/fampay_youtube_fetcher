package worker

import (
	"fampay_youtube_fetcher/internal/pojo"
	"fampay_youtube_fetcher/internal/service"
	"fampay_youtube_fetcher/pkg/youtube"
	"log"
	"time"

	"github.com/google/uuid"
)

type YoutubeWorker struct {
	client     *youtube.YoutubeClient
	service    *service.VideoService
	query      string
	interval   time.Duration
	maxResults int
}

func NewYouTubeWorker(
	youtubeClient *youtube.YoutubeClient,
	service *service.VideoService,
	query string,
	interval time.Duration,
	maxResults int,
) *YoutubeWorker {
	return &YoutubeWorker{
		client:     youtubeClient,
		service:    service,
		query:      query,
		interval:   interval,
		maxResults: maxResults,
	}
}

func (w *YoutubeWorker) Start() {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	if err := w.fetchAndStoreVideos(); err != nil {
		log.Printf("Error fetching videos: %v", err)
	}

	for {
		select {
		case <-ticker.C:
			if err := w.fetchAndStoreVideos(); err != nil {
				log.Printf("Error fetching videos: %v", err)
			}
		}
	}
}

func (w *YoutubeWorker) fetchAndStoreVideos() error {
	lastPublishedTime, err := w.service.GetLatestPublishedDate()
	if err != nil {
		log.Println("Failed to get the latest publish date", err)
		return err
	}
	// Default to 10 hours ago in UTC
	var lastFetchedTime string
	if lastPublishedTime == nil {
		lastFetchedTime = time.Now().UTC().Add((-1 * 2 * 24) * time.Hour).Format(time.RFC3339)
	} else {
		lastFetchedTime = lastPublishedTime.In(time.UTC).Add(1 * time.Minute).Format(time.RFC3339)
	}

	resp, err := w.client.FetchVideos(w.query, w.maxResults, lastFetchedTime)
	if err != nil {
		return err
	}

	for _, item := range resp.Items {
		video := &pojo.VideoMetaData{
			Id:               uuid.New().String(),
			VideoId:          item.ID.VideoID,
			Title:            item.Snippet.Title,
			Description:      item.Snippet.Description,
			DefaultThumbnail: item.Snippet.Thumbnails.Default.URL,
			PublishedAt:      item.Snippet.PublishedAt,
		}

		w.service.CreateVideo(video)
	}

	return nil
}
