package resultJobsMyApi

// а сервис это как сервис отдельный для части кода сервис будет я думаю и для сотрудников и тд

type MyService struct {
}

// Конструктор сервиса вроде так правильно
func NewMyServiceEmployee() *MyService {
	return &MyService{}
}

// Метод над myService
func (s *MyService) SayHello() string {
	return "Hello Vakaram"
}

// Метод над myService
func (s *MyService) HendleMain() string {
	return "Вы на стартовой странице урааааа"
}

// Метод над myService
func (s *MyService) ADDTest(idErr string) string {
	return "Проверь нового пользователя в бд, эх как бы сюда еще что нибудь передать хо хо хо " + idErr
}
