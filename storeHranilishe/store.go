package storeHranilishe

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"testdbNewMetods/internal/app/modelEmployee"
	"time"
)

// здесь создается сторе и регулируется и настраивается
// структура стора
type Store struct {
	config *Config // по сути строка подключения с портом и паролем сюда она нам вернется при инициализации
	DB     *pgxpool.Pool
	CXT    context.Context
}

// вернет нам стркутуру store и мы сможем использовать методы store например open
// а мы при вызове передадим ему строку из конфига
func NewStore(config *Config) (*Store, error) {
	//добавим к контектсу таймаут
	contextParam, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(contextParam, config.DatabaseURL)
	if err != nil {
		logrus.Error("Unable to connect to database: ", err)
		return nil, err
	}
	logrus.Info("Вижу вызвали NEW Store значит в дб мы уже можем обращаться с запросами =) ")

	return &Store{
		config: config,
		DB:     dbpool,
		CXT:    contextParam,
	}, nil
}

func (s *Store) CreateEmployee(emp modelEmployee.EmployeeModel) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	var id string
	err := s.DB.QueryRow(
		ctx,
		"INSERT INTO turnixSchem.employees (login,password) VALUES ($1,$2) RETURNING id ",
		emp.Login, emp.Password).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return "err", err
	}

	return id, nil
}
