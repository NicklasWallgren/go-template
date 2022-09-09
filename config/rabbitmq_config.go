package config

import "fmt"

type RabbitMQConfig struct {
	Name     string
	Password string
	Host     string
}

func NewRabbitMQConfig(name string, password string, host string) *RabbitMQConfig {
	return &RabbitMQConfig{Name: name, Password: password, Host: host}
}

func (r RabbitMQConfig) ToDsn() string {
	return fmt.Sprintf("amqp://%s:%s@%s", r.Name, r.Password, r.Host)
}
