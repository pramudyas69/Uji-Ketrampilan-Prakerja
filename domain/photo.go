package domain

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

type Photo struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"NOT NULL;type:varchar(255);" valid:"required"`
	Caption   string    `json:"caption" gorm:"type:varchar(255);"`
	PhotoURL  string    `json:"photo_url" gorm:"NOT NULL;type:text;" valid:"required"`
	UserID    uint32    `json:"user_id"`
	Comment   []Comment `json:"comments,omitempty" gorm:"foreignKey:photo_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	println(govalidator.IsURL(p.PhotoURL))
	isURL := govalidator.IsURL(p.PhotoURL)
	if !isURL {
		return errors.New("url not valid")
	}

	return
}
func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}

type PhotoUseCase interface {
	CreatePhotoUC(ctx echo.Context) (*Photo, error)
	GetPhotosUC(ctx echo.Context) (*User, error)
	UpdatePhotoUC(ctx echo.Context) (*Photo, error)
	DeletePhotoUC(ctx echo.Context) (*Photo, error)
}

type PhotoRepository interface {
	CreatePhotoRepository(ctx echo.Context) (*Photo, error)
	GetPhotosRepository(ctx echo.Context) (*User, error)
	UpdatePhotoRepository(ctx echo.Context) (*Photo, error)
	DeletePhotoRepository(ctx echo.Context) (*Photo, error)
}
