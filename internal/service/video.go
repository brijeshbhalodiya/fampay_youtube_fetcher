package service

import (
	"fampay_youtube_fetcher/internal/pojo"
	"fampay_youtube_fetcher/internal/repository"
	"time"
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
