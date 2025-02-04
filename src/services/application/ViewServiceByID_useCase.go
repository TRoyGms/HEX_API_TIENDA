package application

import (
	"demo/src/services/domain"
	"demo/src/services/domain/entities"
	"errors"
)

type ViewServiceByIDUseCase struct {
	serviceRepository domain.IServiceRepository
}

func NewViewServiceByIDUseCase(serviceRepository domain.IServiceRepository) *ViewServiceByIDUseCase {
	return &ViewServiceByIDUseCase{serviceRepository: serviceRepository}
}

func (u *ViewServiceByIDUseCase) Execute(serviceID int) (*entities.Service, error) {
	service, err := u.serviceRepository.GetByID(serviceID)
	if err != nil {
		return nil, err
	}
	if service == nil {
		return nil, errors.New("service not found")
	}
	return service, nil
}
