package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"runtime"
)

const (
	ProductionEnv = "production"
	DevelopmentEnv = "development"
)


type Configurations struct {
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	Http_port int `env:"HTTP_PORT" envDefault:"8080"`
	Grpc_port int `env:"GRPC_PORT" envDefault:"50051"`
	Db_host string `env:"DB_HOST" envDefault:"localhost"`
	Db_port int `env:"DB_PORT" envDefault:"5432"`
}

var (cfg Configurations)

func LoadConfig() *Configurations {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	err :=godotenv.Load(filepath.Join(currentDir, "config.yaml"))
	if err != nil {
		log.Print("Error loading .env file: ", err)
	}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return &cfg
}

func GetConfig() *Configurations {
	return &cfg
}