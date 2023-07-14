package postgre

import (
	"github.com/labstack/echo/v4"
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

func (s *SosmedRepository) GetSosmedsRepository(ctx echo.Context) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SosmedRepository) UpdateSosmedRepository(ctx echo.Context) (*domain.SocialMedia, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SosmedRepository) DeleteSosmedRepository(ctx echo.Context) (*domain.SocialMedia, error) {
	//TODO implement me
	panic("implement me")
}
