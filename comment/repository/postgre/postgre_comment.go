package postgre

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"uji/database/redis/repository"
	"uji/domain"
)

type CommentRepository struct {
	db        *gorm.DB
	redisRepo repository.RedisRepository
}

func NewCommentRepository(db *gorm.DB, redisRepo repository.RedisRepository) domain.CommentRepository {
	return &CommentRepository{
		db,
		redisRepo,
	}
}

func (c *CommentRepository) CreateCommentRepository(ctx echo.Context) (*domain.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CommentRepository) GetCommentsRepository(ctx echo.Context) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CommentRepository) UpdateCommentRepository(ctx echo.Context) (*domain.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CommentRepository) DeleteCommentRepository(ctx echo.Context) (*domain.Comment, error) {
	//TODO implement me
	panic("implement me")
}
