package rabbitmq

import "github.com/spf13/viper"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
}

func New() *Config {
	setDefaults()
	return &Config{
		Host:     viper.GetString("RABBITMQ_HOST"),
		Port:     viper.GetString("RABBITMQ_PORT"),
		Username: viper.GetString("RABBITMQ_USER"),
		Password: viper.GetString("RABBITMQ_PASSWORD"),
	}
}

func setDefaults() {
	viper.SetDefault("RABBITMQ_HOST", "localhost")
	viper.SetDefault("RABBITMQ_PORT", "5432")
	viper.SetDefault("RABBITMQ_USER", "postgres")
	viper.SetDefault("RABBITMQ_PASSWORD", "postgres")
}
