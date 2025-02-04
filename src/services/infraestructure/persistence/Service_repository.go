package repositories

import (
	"database/sql"
	"demo/src/services/domain/entities"
	"errors"
	"log"
)

type ServiceRepository struct {
	DB *sql.DB
}

func NewServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{DB: db}
}

func (repo *ServiceRepository) Save(service *entities.Service) (int, error) {
	query := "INSERT INTO services (name, price) VALUES (?, ?)"
	result, err := repo.DB.Exec(query, service.Name, service.Price)
	if err != nil {
		log.Printf("[ServiceRepository.Save] Error inserting service: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("[ServiceRepository.Save] Error retrieving last insert ID: %v", err)
		return 0, err
	}

	log.Println("[ServiceRepository.Save] Service inserted successfully with ID:", id)
	return int(id), nil
}

// Obtener todos los servicios
func (repo *ServiceRepository) GetAll() ([]*entities.Service, error) {
	query := "SELECT id, name, price FROM services"
	rows, err := repo.DB.Query(query)
	if err != nil {
		log.Printf("[ServiceRepository.GetAll] Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var services []*entities.Service // Change to a slice of pointers
	for rows.Next() {
		var service entities.Service
		if err := rows.Scan(&service.ID, &service.Name, &service.Price); err != nil {
			log.Printf("[ServiceRepository.GetAll] Error scanning row: %v", err)
			return nil, err
		}
		// Append a pointer to the service
		services = append(services, &service)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[ServiceRepository.GetAll] Row iteration error: %v", err)
		return nil, err
	}

	return services, nil
}

// Obtener un servicio por ID
func (repo *ServiceRepository) GetById(id int) (*entities.Service, error) {
	query := "SELECT id, name, price FROM services WHERE id = ?"
	row := repo.DB.QueryRow(query, id)

	var service entities.Service
	if err := row.Scan(&service.ID, &service.Name, &service.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("[ServiceRepository.GetById] Service not found with ID: %d", id)
			return nil, nil
		}
		log.Printf("[ServiceRepository.GetById] Error scanning row: %v", err)
		return nil, err
	}

	return &service, nil
}

// Eliminar un servicio por ID
func (repo *ServiceRepository) DeleteById(id int) error {
	query := "DELETE FROM services WHERE id = ?"
	result, err := repo.DB.Exec(query, id)
	if err != nil {
		log.Printf("[ServiceRepository.DeleteById] Error executing query: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[ServiceRepository.DeleteById] Error retrieving affected rows: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("[ServiceRepository.DeleteById] No service found with ID: %d", id)
		return errors.New("no service found to delete")
	}

	return nil
}

// Editar un servicio por ID
func (repo *ServiceRepository) EditById(id int, updatedService *entities.Service) error {
	query := "UPDATE services SET name = ?, price = ? WHERE id = ?"
	result, err := repo.DB.Exec(query, updatedService.Name, updatedService.Price, id)
	if err != nil {
		log.Printf("[ServiceRepository.EditById] Error executing query: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[ServiceRepository.EditById] Error retrieving affected rows: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("[ServiceRepository.EditById] No service found with ID: %d", id)
		return errors.New("no service found to update")
	}

	return nil
}

func (repo *ServiceRepository) GetByID(id int) (*entities.Service, error) {
	query := "SELECT id, name, price FROM services WHERE id = ?"
	row := repo.DB.QueryRow(query, id)

	var service entities.Service
	if err := row.Scan(&service.ID, &service.Name, &service.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("[ServiceRepository.GetByID] Error scanning row: %v", err)
		return nil, err
	}

	return &service, nil
}
