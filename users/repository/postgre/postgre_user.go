package postgre

import (
	"errors"
	"gorm.io/gorm"
	"uji/domain"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepsitory(DB *gorm.DB) domain.UserRepository {
	return &UserRepository{
		DB,
	}
}

func (u *UserRepository) UserRegisterRepository(user *domain.User) error {
	err := u.DB.Where("username = ?", user.Username).Take(&user).Error
	if err == nil {
		return errors.New("Username Duplicate!")
	}

	err = u.DB.Where("email = ?", user.Email).Take(&user).Error
	if err == nil {
		return errors.New("Email Duplicate!")
	}
	return u.DB.Debug().Create(&user).Error
}

func (u *UserRepository) UserLoginRepository(user *domain.User) error {
	return u.DB.Where("email = ?", user.Email).Take(&user).Error
}

func (u *UserRepository) GetUserByIdRepository(id uint32) (*domain.User, error) {
	var user domain.User

	err := u.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u *UserRepository) GetUsersRepository(user []*domain.User) ([]*domain.User, error) {
	if err := u.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) UpdateUserRepository(id uint32, user *domain.User) (*domain.User, error) {
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

	err = u.DB.Save(&existingUser).Error

	return &existingUser, nil
}

func (u *UserRepository) DeleteUserRepository(id uint32) error {
	var user domain.User

	err := u.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return errors.New("record not found!")
	}

	return u.DB.Unscoped().Delete(&user).Error
}
