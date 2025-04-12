package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server serverConfig
	Logger loggerConfig
}

type serverConfig struct {
	Port      int
	LocalHost string
	ApiKey    string
}

type loggerConfig struct {
	Path     string
	LogLevel string
}

func GetConfig() *Config {
	path := getPath(os.Getenv("APP_ENV"))
	v := parseConfig(path, "yml")
	cfg := loadConfig(v)
	return cfg
}

func getPath(env string) string {
	if env == "docker" {
		return "../config/docker"
	} else if env == "production" {
		return "../config/production"
	} else {
		return "../config/development"
	}
}

func parseConfig(filename string, filetype string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType(filetype)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return v
}

func loadConfig(v *viper.Viper) *Config {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}