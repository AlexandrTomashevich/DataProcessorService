package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Database struct {
	User     string `yaml:"username" env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Host     string `yaml:"host" env:"POSTGRES_HOST"`
	Port     string `yaml:"port" env:"POSTGRES_PORT"`
	Dbname   string `yaml:"dbname" env:"POSTGRES_DB"`
	Sslmode  string `yaml:"sslmode" env:"POSTGRES_SSLMODE"`
	Driver   string `yaml:"driver" env:"POSTGRES_DRIVER"`
}

type ServerHTTP struct {
	Address string `yaml:"address" env:"SERVER_ADDRESS"`
}

type Config struct {
	ServerHTTP ServerHTTP `yaml:"server_http"`
	Database   Database   `yaml:"database"`
}

func (db *Database) ConnectionString() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		db.User, db.Password, db.Host, db.Port, db.Dbname, db.Sslmode,
	)
}

const ConfigPath = "configs/config.yaml"

var (
	once       sync.Once
	instance   Config
	configPath string
)

func InitConfig() *Config {
	once.Do(func() {

		instance = Config{}

		if err := cleanenv.ReadConfig(ConfigPath, &instance); err != nil {
			log.Fatalf("failed to read config: %s", err)
		}
		if password := os.Getenv("DB_PASSWORD"); password != "" {
			instance.Database.Password = password
		}
	})

	slog.Info("config successfully initialized", instance)

	return &instance
}
