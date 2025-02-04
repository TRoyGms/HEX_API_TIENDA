package application

import (
	"demo/src/services/domain"
	"demo/src/services/domain/entities"
)

type EditServiceUseCase struct {
	serviceRepo domain.IServiceRepository
}

func NewEditServiceUseCase(serviceRepo domain.IServiceRepository) *EditServiceUseCase {
	return &EditServiceUseCase{serviceRepo: serviceRepo}
}

func (u *EditServiceUseCase) Execute(serviceID int, updatedService *entities.Service) error {
	return u.serviceRepo.EditById(serviceID, updatedService)
}
