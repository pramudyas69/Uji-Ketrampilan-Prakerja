package domain

import (
	"github.com/labstack/echo/v4"
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

type SosmedUseCase interface {
	CreateSosmedUC(ctx echo.Context) (*SocialMedia, error)
	GetSosmedsUC(ctx echo.Context) (*User, error)
	UpdateSosmedUC(ctx echo.Context) (*SocialMedia, error)
	DeleteSosmedUC(ctx echo.Context) (*SocialMedia, error)
}

type SosmedRepository interface {
	CreateSosmedRepository(ctx echo.Context) (*SocialMedia, error)
	GetSosmedsRepository(ctx echo.Context) (*User, error)
	UpdateSosmedRepository(ctx echo.Context) (*SocialMedia, error)
	DeleteSosmedRepository(ctx echo.Context) (*SocialMedia, error)
}
