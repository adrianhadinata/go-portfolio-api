package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
}

type DBConfig struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
	DbDriver string
}

type Config struct {
	ApiConfig
	DBConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DBConfig = DBConfig{
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("DB_USERNAME"),
		DbPass: os.Getenv("DB_PASSWORD"),
		DbName: os.Getenv("DB_NAME"),
		DbDriver: os.Getenv("DB_DRIVER"),
	}

	fmt.Println(c.DBConfig.DbPort)

	if c.ApiConfig.ApiPort == "" {
		return errors.New("apiport cannot be empty")
	}

	if c.DBConfig.DbHost == "" {
		return errors.New("database host cannot be empty")
	}

	if c.DBConfig.DbPort == "" {
		return errors.New("database port cannot be empty")
	}

	if c.DBConfig.DbUser == "" {
		return errors.New("database username cannot be empty")
	}

	if c.DBConfig.DbPass == "" {
		return errors.New("database password cannot be empty")
	}

	if c.DBConfig.DbName == "" {
		return errors.New("database name cannot be empty")
	}

	if c.DBConfig.DbDriver == "" {
		return errors.New("database driver cannot be empty")
	}

	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}
	if err := config.readConfig(); err != nil {
		return nil, err
	}
	return config, nil
}