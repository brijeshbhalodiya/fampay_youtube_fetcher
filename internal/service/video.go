package service

import (
	"fampay_youtube_fetcher/internal/pojo"
	"fampay_youtube_fetcher/internal/repository"
	"time"
)

const (
	DEFAULT_FETCH_LIMIT = 10
)

type VideoService struct {
	videoRepo *repository.VideoRepo
}

func NewVideoService(videoRepo *repository.VideoRepo) *VideoService {
	return &VideoService{videoRepo: videoRepo}
}

func (s *VideoService) CreateVideo(video *pojo.VideoMetaData) error {
	return s.videoRepo.Create(video)
}

func (s *VideoService) GetLatestPublishedDate() (*time.Time, error) {
	return s.videoRepo.GetLatestPublishedDate()
}

func (s *VideoService) GetLatestVideos(limit, offset int) ([]pojo.VideoMetaData, error) {
	if limit <= 0 {
		limit = DEFAULT_FETCH_LIMIT
	}
	if offset < 0 {
		offset = 0
	}

	return s.videoRepo.FindLatest(limit, offset)
}

func (s *VideoService) SearchVideos(query string, limit, offset int) ([]pojo.VideoMetaData, error) {
	if limit <= 0 {
		limit = DEFAULT_FETCH_LIMIT
	}
	if offset < 0 {
		offset = 0
	}

	return s.videoRepo.Search(query, limit, offset)
}
