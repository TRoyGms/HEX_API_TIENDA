package domain

import "demo/src/services/domain/entities"

type IServiceRepository interface {
	Save(service *entities.Service) (int, error)
	GetAll() ([]*entities.Service, error)
	GetByID(id int) (*entities.Service, error)
	DeleteById(id int) error
	EditById(id int, updatedService *entities.Service) error
}

