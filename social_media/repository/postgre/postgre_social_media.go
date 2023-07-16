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

type SosmedRepository struct {
	DB        *gorm.DB
	redisRepo repository.RedisRepository
}

func NewSosmedRepsitory(DB *gorm.DB, redisRepo repository.RedisRepository) domain.SosmedRepository {
	return &SosmedRepository{
		DB,
		redisRepo,
	}
}

func (s *SosmedRepository) CreateSosmedRepository(sosmed *domain.SocialMedia) error {
	ctx := context.Background()
	err := s.redisRepo.DeleteKey(ctx, "users", "sosmeds")
	if err != nil {
		return errors.New("error when clearing data in redis!")
	}
	return s.DB.Debug().Create(&sosmed).Error
}

func (s *SosmedRepository) GetSosmedsRepository(sosmed []*domain.SocialMedia) ([]*domain.SocialMedia, error) {
	ctx := context.Background()

	// Coba mendapatkan data dari Redis
	res, err := s.redisRepo.GetValue(ctx, "sosmeds")
	if err == nil {
		err = json.Unmarshal([]byte(res), &sosmed)
		if err != nil {
			return nil, err
		}
		return sosmed, nil
	}

	if err := s.DB.Find(&sosmed).Error; err != nil {
		return nil, err
	}

	// Simpan data ke Redis
	dataJSON, err := json.Marshal(sosmed)
	if err != nil {
		return nil, err
	}
	err = s.redisRepo.SetValue(ctx, "sosmeds", dataJSON, 1*time.Hour)
	if err != nil {
		return nil, err
	}

	return sosmed, nil
}

func (s *SosmedRepository) UpdateSosmedRepository(id uint, sosmed *domain.SocialMedia) (*domain.SocialMedia, error) {
	var existingUser domain.SocialMedia
	ctx := context.Background()

	err := s.DB.Where("id = ?", id).First(&existingUser).Error
	if err != nil {
		return nil, errors.New("record not found")
	}

	if sosmed.Name != "" {
		existingUser.Name = sosmed.Name
	}
	if sosmed.SocialMediaURL != "" {
		existingUser.SocialMediaURL = sosmed.SocialMediaURL
	}

	err = s.redisRepo.DeleteKey(ctx, "users", "sosmeds")
	if err != nil {
		return nil, errors.New("error when clearing data in redis!")
	}

	err = s.DB.Save(&existingUser).Error

	return &existingUser, nil
}

func (s *SosmedRepository) DeleteSosmedRepository(id uint) error {
	var sosmed domain.SocialMedia
	ctx := context.Background()

	err := s.DB.Where("id = ?", id).First(&sosmed).Error
	if err != nil {
		return errors.New("record not found!")
	}

	err = s.redisRepo.DeleteKey(ctx, "users", "sosmeds")
	if err != nil {
		return errors.New("error when clearing data in redis!")
	}

	return s.DB.Unscoped().Delete(&sosmed).Error
}
