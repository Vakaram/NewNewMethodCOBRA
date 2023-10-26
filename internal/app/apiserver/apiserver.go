package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"net/http"
	"testdbNewMetods/internal/app/modelEmployee"
	"testdbNewMetods/storeHranilishe"
)

type APIServer struct {
	config       *ConfigForApiserver                 //
	logger       *logrus.Logger                      //добавляем логгер в apiserver
	router       *mux.Router                         // то что слушаем и реагируем типа /hello
	employeeRepo *storeHranilishe.EmployeeRepository //тест вынос логики store
}

// Вроде это вообще типо набор по умолчанию типа мы конфиг наш передадим чтобы вызывать старт метод
// но логеры и роутеры они вызываются из своей либы и нам фунецию писать не нужно
func New(config *ConfigForApiserver, store *storeHranilishe.Store) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),    // logrus.NewStore() - это встроенно в логрус а не нами написано
		router: mux.NewRouter(), // возвращает новый экземпляр маршрутизатора
		//store:        store,
		employeeRepo: storeHranilishe.NewEmployeeRepository(store),
	}
}

// тест
// чтобы был доступ ко всем структурам в APIServer его тут надо проинициализировать в стерте
func (s *APIServer) Start() error {
	//логеру добавили уровень логирования и еще можно приписывать настройке в функции  configurateLogger
	if err := s.configurateLogger(); err != nil {
		return err
	}
	s.logger.Info("Запущен метод START")
	//Сказали роутеру что слушать
	s.configRouter()
	//s.configStore()

	return http.ListenAndServe(s.config.HTTPServer.Address, s.router)

}

//тут еще конфик надо сделаь емое
//func (s *APIServer) configStore()  {
//	s.store.
//
//}

// функция для определния поведения логера *конфиг логер типо куда что сохраняем уровень логирование и  тд
func (s *APIServer) configurateLogger() error {
	leverl, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	//взяли обратились к сткрутре у которой есть функция в которой левл записан как
	//debug и потом уже достали это значение и передали его s.logger.SetLevel(leverl)
	s.logger.SetLevel(leverl)
	return nil
}

// запросы будут обрабатываться тут
// а искать программа их пойдет в serviceOtvetAPI
// и еще поставить нужно их в интерфейс конфига в servis styrct
func (s *APIServer) configRouter() {
	s.router.HandleFunc("/", s.HendleMain)       // сюда придем посмотрим а уж потом вызовим функцию которая ниже =)
	s.router.HandleFunc("/hello", s.HandleHello) // сюда придем посмотрим а уж потом вызовим функцию которая ниже =)
	s.router.HandleFunc("/addemploye", s.AddEmployee)
	//s.router.HandleFunc("/createemployee", s.CreateEmployee)
}

func (s *APIServer) HandleHello(w http.ResponseWriter, r *http.Request) {
	//service := epmloyeeService.NewMyService() // определили сервис
	response := s.config.ServisEmployee.SayHello() // определили нужный метод
	fmt.Fprintf(w, response)                       // передали нужный метод
}

func (s *APIServer) HendleMain(w http.ResponseWriter, r *http.Request) {
	//service := epmloyeeService.NewMyService() // определили сервис
	response := s.config.ServisEmployee.HendleMain() // определили нужный метод
	fmt.Fprintf(w, response)                         // передали нужный метод
}

func (s *APIServer) AddEmployee(w http.ResponseWriter, r *http.Request) {
	modelEmployee := modelEmployee.EmployeeModel{
		Login:    "Testlog",
		Password: "TestPass",
	}
	//cxt := s.employeeRepo.Store.CXT
	id, err := s.employeeRepo.CreateEmployee(modelEmployee)
	if err != nil {
		response := s.config.ServisEmployee.AddEmployee(err.Error()) // определили нужный метод
		fmt.Fprintf(w, response)
		return
	}
	response := s.config.ServisEmployee.AddEmployee(id) // определили нужный метод
	fmt.Fprintf(w, response)                            // передали нужный метод
}

//func (s *APIServer) CreateEmployeeINStoreINEmployeerepository() (string, error) {
//	var id string
//	if err := s.store.DB.QueryRow(
//		"INSERT INTO turnixSchem.employees (login,password) VALUES ($1,$2) RETURNING id ",
//		"la la", //
//		"la la la 3 ",
//	).Scan(&id); err != nil {
//		return "err", err
//	}
//
//	return id, nil
//
//}
