package application

import (
	"demo/src/services/domain/entities"
	"demo/src/services/domain"
)

type CreateServiceUseCase struct {
	serviceRepo domain.IServiceRepository
}

func NewCreateServiceUseCase(serviceRepo domain.IServiceRepository) *CreateServiceUseCase {
	return &CreateServiceUseCase{serviceRepo: serviceRepo}
}

func (u *CreateServiceUseCase) Execute(service *entities.Service) error {
	return u.serviceRepo.Create(service)
}
