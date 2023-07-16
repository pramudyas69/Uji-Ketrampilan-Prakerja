package domain

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Message   string    `json:"message" gorm:"NOT NULL;type:text;" valid:"required"`
	UserID    uint32    `json:"user_id"`
	PhotoID   string    `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}

type CommentUseCase interface {
	CreateCommentUC(ctx echo.Context) (*Comment, error)
	GetCommentsUC(ctx echo.Context) (*User, error)
	UpdateCommentUC(ctx echo.Context) (*Comment, error)
	DeleteCommentUC(ctx echo.Context) (*Comment, error)
}

type CommentRepository interface {
	CreateCommentRepository(ctx echo.Context) (*Comment, error)
	GetCommentsRepository(ctx echo.Context) (*User, error)
	UpdateCommentRepository(ctx echo.Context) (*Comment, error)
	DeleteCommentRepository(ctx echo.Context) (*Comment, error)
}
