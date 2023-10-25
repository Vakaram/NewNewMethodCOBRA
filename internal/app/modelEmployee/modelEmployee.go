package modelEmployee

// просто моделька для чтения json и передачи в функции как данные
type EmployeeModel struct {
	ID       int
	Name     string
	Login    string
	Password string
	Email    string
	IsAdmin  bool
}
