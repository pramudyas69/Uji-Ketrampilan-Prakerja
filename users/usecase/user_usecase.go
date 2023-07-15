package usecase

import (
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

func (u UserUseCase) GetUserByIdUc(id uint32) (*domain.User, error) {
	return u.userRepo.GetUserByIdRepository(id)
}

func (u UserUseCase) GetUsersUc(user *[]domain.User) (*[]domain.User, error) {
	return u.userRepo.GetUsersRepository(user)
}

func (u UserUseCase) UpdateUserUc(id uint32, user *domain.User) (*domain.User, error) {
	return u.userRepo.UpdateUserRepository(id, user)
}

func (u UserUseCase) DeleteUserUc(id uint32) error {
	return u.userRepo.DeleteUserRepository(id)
}
