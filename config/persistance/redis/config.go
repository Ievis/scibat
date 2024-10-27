package redis

import "github.com/spf13/viper"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBNumber int32
}

func New() *Config {
	setDefaults()
	return &Config{
		Host:     viper.GetString("REDIS_HOST"),
		Port:     viper.GetString("REDIS_PORT"),
		Username: viper.GetString("REDIS_USER"),
		Password: viper.GetString("REDIS_PASSWORD"),
		DBNumber: viper.GetInt32("REDIS_DB"),
	}
}

func setDefaults() {
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", "5432")
	viper.SetDefault("REDIS_USER", "postgres")
	viper.SetDefault("REDIS_PASSWORD", "postgres")
	viper.SetDefault("REDIS_DB", 0)
}
