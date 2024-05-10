package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	KafkaApiKey    string `mapstructure:"KAFKA_API_KEY"`
	KafkaApiSecret string `mapstructure:"KAFKA_API_SECRET"`
	KafkaBrokers   string `mapstructure:"KAFKA_BROKERS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err

}
