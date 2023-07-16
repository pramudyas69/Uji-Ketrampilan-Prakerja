package usecase

import (
	"github.com/labstack/echo/v4"
	"uji/domain"
)

type CommentUseCase struct {
	commentRepo domain.CommentRepository
}

func NewCommentUseCase(commentRepo domain.CommentRepository) domain.CommentUseCase {
	return &CommentUseCase{
		commentRepo,
	}
}

func (c CommentUseCase) CreateCommentUC(ctx echo.Context) (*domain.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c CommentUseCase) GetCommentsUC(ctx echo.Context) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (c CommentUseCase) UpdateCommentUC(ctx echo.Context) (*domain.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c CommentUseCase) DeleteCommentUC(ctx echo.Context) (*domain.Comment, error) {
	//TODO implement me
	panic("implement me")
}
