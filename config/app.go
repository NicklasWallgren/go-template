package config

import (
	"fmt"
	"github.com/NicklasWallgren/go-template/infrastructure/env"
)

type Database struct {
	Name               string
	User               string
	Password           string
	Host               string
	Port               string
	MigrationDirectory string
	Driver             string // TODO, validate supported dialects
}

func NewDatabase(name string, user string, password string, host string, port string, migrationDirectory string, dialect string) *Database {
	return &Database{Name: name, User: user, Password: password, Host: host, Port: port, MigrationDirectory: migrationDirectory, Driver: dialect}
}

type Log struct {
	level string
}

func NewLog(level string) *Log {
	return &Log{level: level}
}

type HttpServer struct {
	Port     string
	JwtToken string // TODO, belongs to another struct?
}

func NewHttpServer(port string, jwtToken string) *HttpServer {
	return &HttpServer{Port: port, JwtToken: jwtToken}
}

type RabbitMQ struct {
	Name     string
	Password string
	Host     string
}

func NewRabbitMQ(name string, password string, host string) *RabbitMQ {
	return &RabbitMQ{Name: name, Password: password, Host: host}
}

func (r RabbitMQ) ToDsn() string {
	return fmt.Sprintf("amqp://%s:%s@%s", r.Name, r.Password, r.Host)
}

type AppConfig struct {
	Assets     *Assets
	Database   *Database
	Log        Log
	HttpServer HttpServer
	RabbitMQ   *RabbitMQ
}

func NewAppConfig(assets *Assets, env env.Env) *AppConfig {
	return &AppConfig{
		Assets:     assets,
		Database:   NewDatabase(env.DBName, env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBMigrationsDirectory, env.DBDriver),
		Log:        *NewLog(""),
		HttpServer: *NewHttpServer(env.ServerPort, env.JWTSecret),
		RabbitMQ:   NewRabbitMQ(env.RabbitMQUser, env.RabbitMQPassword, env.RabbitMQHost),
	}
}
