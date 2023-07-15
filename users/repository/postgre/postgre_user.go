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

type UserRepository struct {
	DB        *gorm.DB
	redisRepo repository.RedisRepository
}

func NewUserRepsitory(DB *gorm.DB, redisRepo repository.RedisRepository) domain.UserRepository {
	return &UserRepository{
		DB,
		redisRepo,
	}
}

func (u *UserRepository) UserRegisterRepository(user *domain.User) error {
	ctx := context.Background()
	err := u.DB.Where("username = ?", user.Username).Take(&user).Error
	if err == nil {
		return errors.New("Username Duplicate!")
	}

	err = u.DB.Where("email = ?", user.Email).Take(&user).Error
	if err == nil {
		return errors.New("Email Duplicate!")
	}
	err = u.redisRepo.DeleteKey(ctx, "users")
	if err != nil {
		return errors.New("error when clearing data in redis!")
	}

	err = u.DB.Debug().Create(&user).Error

	return err
}

func (u *UserRepository) UserLoginRepository(user *domain.User) error {
	return u.DB.Where("email = ?", user.Email).Take(&user).Error
}

func (u *UserRepository) GetUserByIdRepository(id uint32) (*domain.User, error) {
	var user domain.User

	err := u.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u *UserRepository) GetUsersRepository(user *[]domain.User) (*[]domain.User, error) {
	ctx := context.Background()

	// Coba mendapatkan data dari Redis
	res, err := u.redisRepo.GetValue(ctx, "users")
	if err == nil {
		err = json.Unmarshal([]byte(res), &user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	// Jika data tidak ditemukan di Redis, ambil data dari database
	if err := u.DB.Find(&user).Error; err != nil {
		return nil, err
	}

	// Simpan data ke Redis
	dataJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	err = u.redisRepo.SetValue(ctx, "users", dataJSON, 1*time.Hour)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) UpdateUserRepository(id uint32, user *domain.User) (*domain.User, error) {
	ctx := context.Background()
	var existingUser domain.User

	err := u.DB.Where("id = ?", id).First(&existingUser).Error
	if err != nil {
		return nil, errors.New("record not found")
	}

	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}
	if user.Age != 0 {
		existingUser.Age = user.Age
	}

	err = u.redisRepo.DeleteKey(ctx, "users")
	if err != nil {
		return nil, errors.New("error when clearing data in redis!")
	}

	err = u.DB.Save(&existingUser).Error

	return &existingUser, err
}

func (u *UserRepository) DeleteUserRepository(id uint32) error {
	var user domain.User

	err := u.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return errors.New("record not found!")
	}

	return u.DB.Unscoped().Delete(&user).Error
}
