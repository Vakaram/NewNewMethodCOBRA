package apiserver

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type ConfigForApiserver struct {
	Env            string `yaml:"env"    env-default:"local"` // для сервера
	HTTPServer     `yaml:"http_server"`
	LogLevel       string `yaml:"log-level"    env-default:"debug"` // для сервера
	ServisEmployee servis
}

type HTTPServer struct {
	Address     string        `yaml:"address"      env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout"      env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type servis interface {
	SayHello() string
	HendleMain() string
	AddEmployee(id string) string
}

// MustNewConfigApiSrever: заполнять будет нашу стрктуру конфига ConfigForApiserver
func MustNewConfigApiSrever() *ConfigForApiserver {
	configPath := "internal/config/local.yaml" // os.Getenv("CONFIG_PATH")

	// check if file exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var configReturnFilled ConfigForApiserver

	if err := cleanenv.ReadConfig(configPath, &configReturnFilled); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	fmt.Printf("Собрали новый конфиг для APIServera : %v", configReturnFilled)
	//configReturnFilled.LogLevel = "debug"
	return &configReturnFilled
}
