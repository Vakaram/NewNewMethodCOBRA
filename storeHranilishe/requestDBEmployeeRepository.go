package storeHranilishe

type EmployeeRepository struct {
	Store *Store
}

func NewEmployeeRepository(store *Store) *EmployeeRepository {
	return &EmployeeRepository{
		Store: store,
	}
}

// тестовая функция проверить работу db
func (r *EmployeeRepository) CreateEmployee() (string, error) {
	var id string
	if err := r.Store.DB.QueryRow(
		"INSERT INTO turnixSchem.employees (login,password) VALUES ($1,$2) RETURNING id ",
		"la la",
		"la la la 3 ").Scan(&id); err != nil {
		return "err", err
	}
	return id, nil
}
