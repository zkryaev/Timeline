package config

import (
	"log"
	"os"
	"time"
	"timeline/internal/libs/envars"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App   Application `yaml:"app"`
	DB    Database
	Mail  Mail
	Token Token `yaml:"token"`
}

type Application struct {
	Env        string `yaml:"env" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Host        string        `yaml:"host" env-default:"localhost"`
	Port        string        `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"5m"`
}

type Database struct {
	Protocol string `env:"DB" env-required:"true"`
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     string `env:"DB_PORT" env-required:"true"`
	Name     string `env:"DB_NAME" env-required:"true"`
	User     string `env:"DB_USER" env-required:"true"`
	Password string `env:"DB_PASSWD" env-required:"true"`
	SSLmode  string `env:"DB_SSLMODE" env-required:"true"`
}

type Mail struct {
	Host     string `env:"MAIL_HOST" env-required:"true"`
	Port     int    `env:"MAIL_PORT" env-required:"true"`
	User     string `env:"MAIL_USER" env-required:"true"`
	Password string `env:"MAIL_PASSWD" env-required:"true"`
}

type Token struct {
	AccessTTL  time.Duration `yaml:"access_ttl" env-default:"1m"`
	RefreshTTL time.Duration `yaml:"refresh_ttl" env-default:"5m"`
}

func MustLoad() Config {
	configPath := envars.GetPath("CONFIG_PATH")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("the cfg file doesn't exist at the path: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed with reading config: %s", err)
	}
	return cfg
}
