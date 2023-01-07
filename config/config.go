package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Logger   LoggerConfig
	Postgres PostgresConfig
}

type LoggerConfig struct {
	Level string
}

type ServerConfig struct {
	Host                        string
	Port                        string
	ShowUnknownErrorsInResponse bool
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	PgDriver string
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	if _, ok := os.LookupEnv("LOCAL"); ok {
		v.AddConfigPath("config")
	} else {
		v.AddConfigPath("/app/config")
	}

	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
