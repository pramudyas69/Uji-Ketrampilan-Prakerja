package usecase

import (
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

func (s SosmedUseCase) GetSosmedsUC(sosmed []*domain.SocialMedia) ([]*domain.SocialMedia, error) {
	return s.sosmedRepo.GetSosmedsRepository(sosmed)
}

func (s SosmedUseCase) UpdateSosmedUC(id uint, sosmed *domain.SocialMedia) (*domain.SocialMedia, error) {
	return s.sosmedRepo.UpdateSosmedRepository(id, sosmed)
}

func (s SosmedUseCase) DeleteSosmedUC(id uint) error {
	return s.sosmedRepo.DeleteSosmedRepository(id)
}
