package handler

import "ugmo/repository"

type videoHandler struct {
	videoService VideoService
}

type IdResponse struct {
	Videos []int `db:"vid_id"`
}

type VideoService interface {
	GetVideosId(repository.Video) (*IdResponse, error)
	GetImage(int) (string, error)
	GetVideo(int) (string, error)
}
