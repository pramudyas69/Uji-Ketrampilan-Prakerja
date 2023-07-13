package usecase

import (
	"github.com/labstack/echo/v4"
	"uji/domain"
)

type UserUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) domain.UserUseCase {
	return UserUseCase{
		userRepo: repo,
	}
}

func (u UserUseCase) UserRegisterUc(user *domain.User) error {
	return u.userRepo.UserRegisterRepository(user)
}

func (u UserUseCase) UserLoginUc(user *domain.User) error {
	return u.userRepo.UserLoginRepository(user)
}

func (u UserUseCase) GetUserByIdUc(ctx echo.Context) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) GetUsersUc(ctx echo.Context) ([]*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) UpdateUserUc(ctx echo.Context) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) DeleteUserUc(ctx echo.Context) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}
