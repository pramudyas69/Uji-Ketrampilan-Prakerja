package usecase

import (
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

func (c CommentUseCase) CreateCommentUC(comment *domain.Comment) error {
	return c.commentRepo.CreateCommentRepository(comment)
}

func (c CommentUseCase) GetCommentsUC(comment *[]domain.Comment) (*[]domain.Comment, error) {
	return c.commentRepo.GetCommentsRepository(comment)
}

func (c CommentUseCase) UpdateCommentUC(id uint, comment *domain.Comment) (*domain.Comment, error) {
	return c.commentRepo.UpdateCommentRepository(id, comment)
}

func (c CommentUseCase) DeleteCommentUC(id uint) error {
	return c.commentRepo.DeleteCommentRepository(id)
}
