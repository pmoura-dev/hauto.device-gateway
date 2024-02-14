package config

import (
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type TransactionServiceConfig struct {
	Host string
	Port string
}

func GetTransactionServiceConfig() TransactionServiceConfig {
	return TransactionServiceConfig{
		Host: getEnv("TRANSACTION_SERVICE_HOST", "192.168.1.102"),
		Port: getEnv("TRANSACTION_SERVICE_PORT", "8080"),
	}
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

func GetRabbitMQConfig() RabbitMQConfig {
	return RabbitMQConfig{
		Host:     getEnv("RABBITMQ_HOST", "192.168.1.102"),
		Port:     getEnv("RABBITMQ_PORT", "5672"),
		User:     getEnv("RABBITMQ_USER", "guest"),
		Password: getEnv("RABBITMQ_PASSWORD", "guest"),
	}
}
