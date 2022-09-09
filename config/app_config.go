package config

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/env"
)

type AppConfig struct {
	Assets     *AssetsConfig
	Database   *DatabaseConfig
	HTTPServer *HTTPServerConfig
	RabbitMQ   *RabbitMQConfig
}

func NewAppConfig(assets *AssetsConfig, env env.Env) *AppConfig {
	// TODO, inject via FX?
	db := NewDatabaseConfig(
		env.DBName, env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBMigrationsDirectory, env.DBDriver)

	return &AppConfig{
		Assets:     assets,
		Database:   db,
		HTTPServer: NewHTTPServerConfig(env.ServerPort),
		RabbitMQ:   NewRabbitMQConfig(env.RabbitMQUser, env.RabbitMQPassword, env.RabbitMQHost),
	}
}
