package handlers

import "github.com/michaelwp/trackme/internal/repository"

type PhotoHandler interface {
}

type photoHandler struct {
	repository repository.PhotoRepository
}

func NewPhotoHandler(photoRepository repository.PhotoRepository) PhotoHandler {
	return photoHandler{
		repository: photoRepository,
	}
}
