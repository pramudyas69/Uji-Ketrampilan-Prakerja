package domain

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

type SocialMedia struct {
	ID             uint      `json:"id" gorm:"primarykey"`
	Name           string    `json:"name" gorm:"NOT NULL;type:varchar(255);"`
	SocialMediaURL string    `json:"social_media_url" gorm:"NOT NULL;type:text;"`
	UserID         uint32    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {

	return nil
}

type SosmedUseCase interface {
	CreateSosmedUC(sosmed *SocialMedia) error
	GetSosmedsUC(ctx echo.Context) (*User, error)
	UpdateSosmedUC(ctx echo.Context) (*SocialMedia, error)
	DeleteSosmedUC(ctx echo.Context) (*SocialMedia, error)
}

type SosmedRepository interface {
	CreateSosmedRepository(sosmed *SocialMedia) error
	GetSosmedsRepository(ctx echo.Context) (*User, error)
	UpdateSosmedRepository(ctx echo.Context) (*SocialMedia, error)
	DeleteSosmedRepository(ctx echo.Context) (*SocialMedia, error)
}
