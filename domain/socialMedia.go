package domain

import (
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
	GetSosmedsUC(sosmed []*SocialMedia) ([]*SocialMedia, error)
	UpdateSosmedUC(id uint, sosmed *SocialMedia) (*SocialMedia, error)
	DeleteSosmedUC(id uint) error
}

type SosmedRepository interface {
	CreateSosmedRepository(sosmed *SocialMedia) error
	GetSosmedsRepository(sosmed []*SocialMedia) ([]*SocialMedia, error)
	UpdateSosmedRepository(id uint, sosmed *SocialMedia) (*SocialMedia, error)
	DeleteSosmedRepository(id uint) error
}
