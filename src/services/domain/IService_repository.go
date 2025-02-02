package domain

import "demo/src/services/domain/entities"

type IServiceRepository interface {
	Create(service *entities.Service) error
	Delete(serviceID int) error
	Update(service *entities.Service) error
	GetByID(serviceID int) (*entities.Service, error)
	GetAll() ([]*entities.Service, error)
}
