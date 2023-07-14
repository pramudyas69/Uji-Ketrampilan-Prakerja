package usecase

import (
	"github.com/labstack/echo/v4"
	"uji/domain"
)

type SosmedUseCase struct {
	sosmedRepo domain.SosmedRepository
}

func NewSosmedUseCase(repo domain.SosmedRepository) domain.SosmedUseCase {
	return SosmedUseCase{
		sosmedRepo: repo,
	}
}

func (s SosmedUseCase) CreateSosmedUC(sosmed *domain.SocialMedia) error {
	return s.sosmedRepo.CreateSosmedRepository(sosmed)
}

func (s SosmedUseCase) GetSosmedsUC(ctx echo.Context) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s SosmedUseCase) UpdateSosmedUC(ctx echo.Context) (*domain.SocialMedia, error) {
	//TODO implement me
	panic("implement me")
}

func (s SosmedUseCase) DeleteSosmedUC(ctx echo.Context) (*domain.SocialMedia, error) {
	//TODO implement me
	panic("implement me")
}
