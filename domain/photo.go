package domain

import (
	"time"
)

type Photo struct {
	ID        uint32    `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"NOT NULL;type:varchar(255);" valid:"required"`
	Caption   string    `json:"caption" gorm:"type:varchar(255);"`
	PhotoURL  string    `json:"photo_url" gorm:"NOT NULL;type:text;" valid:"required"`
	UserID    uint32    `json:"user_id"`
	Comment   []Comment `json:"comments,omitempty" gorm:"foreignKey:photo_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoUpdateInput struct {
	ID       uint32 `json:"id" gorm:"primarykey"`
	Title    string `json:"title" gorm:"NOT NULL;type:varchar(255);"`
	Caption  string `json:"caption" gorm:"type:varchar(255);"`
	PhotoURL string `json:"photo_url" gorm:"NOT NULL;type:text;"`
}

type PhotoUseCase interface {
	CreatePhotoUC(photo *Photo) error
	GetPhotosUC(photo *[]Photo) (*[]Photo, error)
	UpdatePhotoUC(id uint, photo *Photo) (*Photo, error)
	DeletePhotoUC(id uint) error
}

type PhotoRepository interface {
	CreatePhotoRepository(photo *Photo) error
	GetPhotosRepository(photo *[]Photo) (*[]Photo, error)
	UpdatePhotoRepository(id uint, photo *Photo) (*Photo, error)
	DeletePhotoRepository(id uint) error
}
