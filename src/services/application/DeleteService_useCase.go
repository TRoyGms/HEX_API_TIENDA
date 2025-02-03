package application

import "demo/src/services/domain"

type DeleteServiceUseCase struct {
	serviceRepo domain.IServiceRepository
}

func NewDeleteServiceUseCase(serviceRepo domain.IServiceRepository) *DeleteServiceUseCase {
	return &DeleteServiceUseCase{serviceRepo: serviceRepo}
}

func (u *DeleteServiceUseCase) Execute(serviceID int) error {
	return u.serviceRepo.DeleteById(serviceID)
}
