package config

import (
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DBHOST"`
	DBName     string `mapstructure:"DBNAME"`
	DBUser     string `mapstructure:"DBUSER"`
	DBPort     string `mapstructure:"DBPORT"`
	DBPassword string `mapstructure:"DBPASSWORD"`
}

var envs = []string{
	"DBHOST", "DBNAME", "DBUSER", "DBPORT", "DBPASSWORD",
}

func LoadConfig() (Config, error) {
	var cfg Config
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return cfg, err
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	if err := validator.New().Struct(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
