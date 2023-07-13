package postgre

import (
	"github.com/labstack/echo/v4"
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
	return u.DB.Create(&user).Error
}

func (u *UserRepository) UserLoginRepository(user *domain.User) error {
	return u.DB.Where("email = ?", user.Email).Take(&user).Error
}

func (u *UserRepository) GetUserByIdRepository(ctx echo.Context) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) GetUsersRepository(ctx echo.Context) ([]*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) UpdateUserRepository(ctx echo.Context) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) DeleteUserRepository(ctx echo.Context) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}
