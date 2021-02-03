package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type (
	Config struct {
		HTTP  HTTPConfig
		MySQL MySQLConfig
	}

	HTTPConfig struct {
		Host               string
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegabytes"`
	}

	MySQLConfig struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string
	}
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Init(path string) (*Config, error) {

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0]) // folder
	viper.SetConfigName(path[1]) // config file name
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("mysql", &cfg.MySQL); err != nil {
		return err
	}
	return nil
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = os.Getenv("HTTP_HOST")

	cfg.MySQL.User = os.Getenv("DB_USER")
	cfg.MySQL.Password = os.Getenv("DB_PASSWORD")
	cfg.MySQL.Host = os.Getenv("DB_HOST")
}
