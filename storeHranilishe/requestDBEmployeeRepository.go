package storeHranilishe

import (
	"fmt"
	"os"
	"testdbNewMetods/internal/app/modelEmployee"
)

type EmployeeRepository struct {
	Store *Store
}

func NewEmployeeRepository(store *Store) *EmployeeRepository {
	return &EmployeeRepository{
		Store: store,
	}
}

// тестовая функция проверить работу db
func (r *EmployeeRepository) CreateEmployee(emp modelEmployee.EmployeeModel) (string, error) {
	id, err := r.Store.CreateEmployee(emp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return "err", err
	}
	return id, nil
}
