package config

import (
	"fmt"
	"github.com/spf13/viper"
	"scibat/config/messaging/rabbitmq"
	"scibat/config/persistance/postgresql"
	"scibat/config/persistance/redis"
)

var Conf *Config

type Config struct {
	Postgres *postgresql.Config
	Redis    *redis.Config
	RabbitMq *rabbitmq.Config
}

func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var conf Config
	conf.Postgres = postgresql.New()
	fmt.Println(conf.Postgres.Port)
	conf.Redis = redis.New()
	conf.RabbitMq = rabbitmq.New()
	Conf = &conf
}

func Get() *Config {
	return Conf
}
