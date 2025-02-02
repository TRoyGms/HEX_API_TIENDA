package application

import (
	"demo/src/services/domain"
	"demo/src/services/domain/entities"
)

type ViewServiceUseCase struct {
	serviceRepo domain.IServiceRepository
}

func NewViewServiceUseCase(serviceRepo domain.IServiceRepository) *ViewServiceUseCase {
	return &ViewServiceUseCase{serviceRepo: serviceRepo}
}

func (u *ViewServiceUseCase) Execute(serviceID int) (*entities.Service, error) {
	return u.serviceRepo.GetByID(serviceID)
}
