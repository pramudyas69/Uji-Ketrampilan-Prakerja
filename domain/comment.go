package domain

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Message   string    `json:"message" gorm:"NOT NULL;type:text;" valid:"required"`
	UserID    uint32    `json:"user_id"`
	PhotoID   uint32    `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentUseCase interface {
	CreateCommentUC(comment *Comment) error
	GetCommentsUC(comment *[]Comment) (*[]Comment, error)
	UpdateCommentUC(id uint, comment *Comment) (*Comment, error)
	DeleteCommentUC(id uint) error
}

type CommentRepository interface {
	CreateCommentRepository(comment *Comment) error
	GetCommentsRepository(comment *[]Comment) (*[]Comment, error)
	UpdateCommentRepository(id uint, comment *Comment) (*Comment, error)
	DeleteCommentRepository(id uint) error
}
