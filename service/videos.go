package service

import (
	"fmt"
	"ugmo/handler"
	"ugmo/repository"
)

type videoService struct {
	videosRepo repository.VideoRepository
}

func NewVideoService(videoRepo repository.VideoRepository) videoService {
	return videoService{videosRepo: videoRepo}
}

func (s videoService) GetVideosId(video repository.Video) (*handler.IdResponse, error) {
	videos, err := s.videosRepo.GetVideosId(video)
	if err != nil {
		return nil, err
	}

	result := handler.IdResponse{
		Videos: videos,
	}
	return &result, nil
}

func (s videoService) GetImage(id int) (string, error) {
	path, err := s.videosRepo.GetImage(id)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (s videoService) GetVideo(id int) (string, error) {
	path, err := s.videosRepo.GetVideo(id)
	if err != nil {
		return "", err
	}
	fmt.Println(path)
	return path, nil
}
