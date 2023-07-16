package postgre

import (
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
	"uji/database/redis/repository"
	"uji/domain"
)

type PhotoRepository struct {
	db        *gorm.DB
	redisrepo repository.RedisRepository
}

func NewPhotoRepository(db *gorm.DB, redisRepo repository.RedisRepository) domain.PhotoRepository {
	return &PhotoRepository{
		db,
		redisRepo,
	}
}

func (p *PhotoRepository) CreatePhotoRepository(photo *domain.Photo) error {
	ctx := context.Background()
	err := p.redisrepo.DeleteKey(ctx, "users", "photos")
	if err != nil {
		return errors.New("error when clearing data in redis!")
	}
	return p.db.Debug().Create(&photo).Error
}

func (p *PhotoRepository) GetPhotosRepository(photo *[]domain.Photo) (*[]domain.Photo, error) {
	ctx := context.Background()

	// Coba mendapatkan data dari Redis
	res, err := p.redisrepo.GetValue(ctx, "photos")
	if err == nil {
		err = json.Unmarshal([]byte(res), &photo)
		if err != nil {
			return nil, err
		}
		return photo, nil
	}

	if err := p.db.Preload("Comment").Find(&photo).Error; err != nil {
		return nil, err
	}

	// Simpan data ke Redis
	dataJSON, err := json.Marshal(photo)
	if err != nil {
		return nil, err
	}
	err = p.redisrepo.SetValue(ctx, "photos", dataJSON, 1*time.Hour)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (p *PhotoRepository) UpdatePhotoRepository(id uint, photo *domain.Photo) (*domain.Photo, error) {
	ctx := context.Background()
	var existingUser domain.Photo

	err := p.db.Where("id = ?", id).First(&existingUser).Error
	if err != nil {
		return nil, errors.New("record not found")
	}

	if photo.Title != "" {
		existingUser.Title = photo.Title
	}
	if photo.PhotoURL != "" {
		existingUser.PhotoURL = photo.PhotoURL
	}
	if photo.Caption != "" {
		existingUser.Caption = photo.Caption
	}

	err = p.redisrepo.DeleteKey(ctx, "users", "photos")
	if err != nil {
		return nil, errors.New("error when clearing data in redis!")
	}

	err = p.db.Save(&existingUser).Error

	return &existingUser, err
}

func (p *PhotoRepository) DeletePhotoRepository(id uint) error {
	var photo domain.Photo

	ctx := context.Background()
	err := p.redisrepo.DeleteKey(ctx, "users", "photos")
	if err != nil {
		return errors.New("error when clearing data in redis!")
	}

	err = p.db.Where("id = ?", id).First(&photo).Error
	if err != nil {
		return errors.New("record not found!")
	}

	return p.db.Unscoped().Delete(&photo).Error
}
