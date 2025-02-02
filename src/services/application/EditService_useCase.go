package application

import (
	"demo/src/services/domain/entities"
	"demo/src/services/domain"
)

type EditServiceUseCase struct {
	serviceRepo domain.IServiceRepository
}

func NewEditServiceUseCase(serviceRepo domain.IServiceRepository) *EditServiceUseCase {
	return &EditServiceUseCase{serviceRepo: serviceRepo}
}

func (u *EditServiceUseCase) Execute(service *entities.Service) error {
	return u.serviceRepo.Update(service)
}
