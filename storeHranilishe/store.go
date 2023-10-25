package storeHranilishe

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// здесь создается сторе и регулируется и настраивается
// структура стора
type Store struct {
	config *Config // по сути строка подключения с портом и паролем сюда она нам вернется при инициализации
	db     *sql.DB
}

// вернет нам стркутуру store и мы сможем использовать методы store например open
// а мы при вызове передадим ему строку из конфига
func NewStore(config *Config) (*Store, error) {
	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logrus.Fatal("Ошибка в пинге NewSTORE ")
	}

	//logrus.Info("Пинг прошел все хорошо")

	logrus.Info("Вижу вызвали NEW Store значит в дб мы уже можем обращаться с запросами =) ")

	return &Store{
		config: config,
		db:     db,
	}, nil
}
