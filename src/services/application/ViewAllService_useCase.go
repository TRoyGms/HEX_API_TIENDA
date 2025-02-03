package application

import (
	"demo/src/services/domain"
	"demo/src/services/domain/entities"
)

type ViewAllServicesUseCase struct {
	serviceRepository domain.IServiceRepository
}

func NewViewAllServicesUseCase(serviceRepository domain.IServiceRepository) *ViewAllServicesUseCase {
	return &ViewAllServicesUseCase{serviceRepository: serviceRepository}
}

func (u *ViewAllServicesUseCase) Execute() ([]entities.Service, error) {
	services, err := u.serviceRepository.GetAll()
	if err != nil {
		return nil, err
	}

	// Convertir []*entities.Service a []entities.Service
	var result []entities.Service
	for _, service := range services {
		result = append(result, *service)
	}

	return result, nil
}
