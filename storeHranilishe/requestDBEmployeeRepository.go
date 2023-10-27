package storeHranilishe

import (
	"fmt"
	"net/http"
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

type InterfaceEmployeeRepository interface {
	AddEmployee(w http.ResponseWriter, r *http.Request)
}

// тестовая функция проверить работу db
//func (r *EmployeeRepository) CreateEmployee(emp modelEmployee.EmployeeModel) (string, error) {
//	id, err := r.Store.CreateEmployee(emp)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
//		return "err", err
//	}
//	return id, nil
//}

func (e *EmployeeRepository) AddEmployee(w http.ResponseWriter, r *http.Request) {
	modelEmployee := modelEmployee.EmployeeModel{
		Login:    "Testlog",
		Password: "TestPass",
	}
	id, err := e.Store.CreateEmployee(modelEmployee)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return
	}

	response := id           // определили нужный метод
	fmt.Fprintf(w, response) // передали нужный метод
}
