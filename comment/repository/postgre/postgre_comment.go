package postgre

import (
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
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

func (c *CommentRepository) CreateCommentRepository(comment *domain.Comment) error {
	ctx := context.Background()
	err := c.redisRepo.DeleteKey(ctx, "users", "photos", "comments")
	if err != nil {
		return errors.New("error when clearing data in redis!")
	}
	return c.db.Debug().Create(&comment).Error
}

func (c *CommentRepository) GetCommentsRepository(comment *[]domain.Comment) (*[]domain.Comment, error) {
	ctx := context.Background()

	// Coba mendapatkan data dari Redis
	res, err := c.redisRepo.GetValue(ctx, "comments")
	if err == nil {
		err = json.Unmarshal([]byte(res), &comment)
		if err != nil {
			return nil, err
		}
		return comment, nil
	}

	if err := c.db.Find(&comment).Error; err != nil {
		return nil, err
	}

	// Simpan data ke Redis
	dataJSON, err := json.Marshal(comment)
	if err != nil {
		return nil, err
	}
	err = c.redisRepo.SetValue(ctx, "comments", dataJSON, 1*time.Hour)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *CommentRepository) UpdateCommentRepository(id uint, comment *domain.Comment) (*domain.Comment, error) {
	ctx := context.Background()
	var newComment *domain.Comment

	err := c.db.Where("id = ?", id).First(&newComment).Error
	if err != nil {
		return nil, errors.New("record not found")
	}

	newComment.Message = comment.Message

	err = c.redisRepo.DeleteKey(ctx, "users", "photos", "comments")
	if err != nil {
		return nil, errors.New("error when clearing data in redis!")
	}

	err = c.db.Save(&newComment).Error

	return newComment, err
}

func (c *CommentRepository) DeleteCommentRepository(id uint) error {
	var comment domain.Comment

	ctx := context.Background()
	err := c.redisRepo.DeleteKey(ctx, "users", "photos", "comments")
	if err != nil {
		return errors.New("error when clearing data in redis!")
	}

	err = c.db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return errors.New("record not found!")
	}

	return c.db.Unscoped().Delete(&comment).Error
}
