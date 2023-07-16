package usecase

import (
	"uji/domain"
)

type PhotoUseCase struct {
	photoRepo domain.PhotoRepository
}

func NewPhotoRepository(photoRepo domain.PhotoRepository) domain.PhotoUseCase {
	return &PhotoUseCase{
		photoRepo,
	}
}

func (p PhotoUseCase) CreatePhotoUC(photo *domain.Photo) error {
	return p.photoRepo.CreatePhotoRepository(photo)
}

func (p PhotoUseCase) GetPhotosUC(photo *[]domain.Photo) (*[]domain.Photo, error) {
	return p.photoRepo.GetPhotosRepository(photo)
}

func (p PhotoUseCase) UpdatePhotoUC(id uint, photo *domain.Photo) (*domain.Photo, error) {
	return p.photoRepo.UpdatePhotoRepository(id, photo)
}

func (p PhotoUseCase) DeletePhotoUC(id uint) error {
	return p.photoRepo.DeletePhotoRepository(id)
}
