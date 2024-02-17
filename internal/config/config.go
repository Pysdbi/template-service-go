package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		Dsn   `yaml:"database"`
		Minio `yaml:"minio"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true"  yaml:"name"    env:"AppName"`
		Version string `env-required:"false" yaml:"version" env:"AppVersion"`
		Domain  string `env-default:"false" yaml:"domain"  env:"AppDomain"`
		Debug   bool   `env-default:"false"  yaml:"debug"   env:"AppDebug"`
	}

	// HTTP -.
	HTTP struct {
		Host string `env-required:"true" yaml:"host" env:"HttpHost"`
		Port string `env-required:"true" yaml:"port" env:"HttpPort"`
	}

	// Dsn -.
	Dsn struct {
		Database   string `env-required:"true" yaml:"database"   env:"DsnDatabase"`
		Clickhouse string `env-required:"true" yaml:"clickhouse" env:"DsnClickhouse"`
		Amqp       string `env-required:"true" yaml:"amqp"       env:"DsnAmqp"`
	}

	// Minio -.
	Minio struct {
		Host      string `env-required:"true" yaml:"host" env:"MinioHost"`
		AccessKey string `env-required:"true" yaml:"accessKey" env:"MinioAccessKey"`
		SecretKey string `env-required:"true" yaml:"secretKey" env:"MinioSecretKey"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	envPath := ".env"

	_, err := os.Stat(envPath)
	if err == nil {
		if err = godotenv.Load(envPath); err != nil {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("env file (%v) does`t exists. %w", envPath, err)
	}

	if err = cleanenv.ReadConfig("./config.yml", cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	if err = cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
