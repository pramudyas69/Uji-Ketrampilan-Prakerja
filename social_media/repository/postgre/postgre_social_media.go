package postgre

import (
	"errors"
	"gorm.io/gorm"
	"uji/domain"
)

type SosmedRepository struct {
	DB *gorm.DB
}

func NewSosmedRepsitory(DB *gorm.DB) domain.SosmedRepository {
	return &SosmedRepository{
		DB,
	}
}

func (s *SosmedRepository) CreateSosmedRepository(sosmed *domain.SocialMedia) error {
	return s.DB.Debug().Create(&sosmed).Error
}

func (s *SosmedRepository) GetSosmedsRepository(sosmed []*domain.SocialMedia) ([]*domain.SocialMedia, error) {
	if err := s.DB.Find(&sosmed).Error; err != nil {
		return nil, err
	}
	return sosmed, nil
}

func (s *SosmedRepository) UpdateSosmedRepository(id uint, sosmed *domain.SocialMedia) (*domain.SocialMedia, error) {
	var existingUser domain.SocialMedia

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

	err = s.DB.Save(&existingUser).Error

	return &existingUser, nil
}

func (s *SosmedRepository) DeleteSosmedRepository(id uint) error {
	var sosmed domain.SocialMedia

	err := s.DB.Where("id = ?", id).First(&sosmed).Error
	if err != nil {
		return errors.New("record not found!")
	}

	return s.DB.Unscoped().Delete(&sosmed).Error
}
