package storeHranilishe

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

// у конфига по сути только строка подключения
type Config struct {
	DatabaseURL string `yaml:"database_url"` // для сервера
}

// функция нужна чтобы вернуть config те распарсить доку и вернуть строку подключения
func NewConfigReturnStrokaPodkl() *Config {
	configPath := "internal/config/local.yaml" // os.Getenv("CONFIG_PATH")
	// check if file exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	// проверяю что удалось считать строку из файла
	logrus.Info("Удалось выпарсить вот такоую строку " + cfg.DatabaseURL)

	return &cfg

}
