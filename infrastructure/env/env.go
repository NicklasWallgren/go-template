package env

import (
	"fmt"
	"log"

	"github.com/iamolegga/enviper"
	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	ServerPort            string `mapstructure:"SERVER_PORT"`
	Environment           string `mapstructure:"ENV"`
	LogLevel              string `mapstructure:"LOG_LEVEL"`
	DBUsername            string `mapstructure:"DB_USER"`
	DBPassword            string `mapstructure:"DB_PASS"`
	DBHost                string `mapstructure:"DB_HOST"`
	DBPort                string `mapstructure:"DB_PORT"`
	DBName                string `mapstructure:"DB_NAME"`
	DBMigrationsDirectory string `mapstructure:"DB_MIGRATION_DIR"`
	DBDialect             string `mapstructure:"DB_DIALECT"` // TODO, rename to DBDriver?
	JWTSecret             string `mapstructure:"JWT_SECRET"`
	RabbitMQUser          string `mapstructure:"RABBITMQ_USER"`
	RabbitMQPassword      string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQHost          string `mapstructure:"RABBITMQ_HOST"`
}

// NewEnv creates a new environment
func NewEnv() Env {
	return NewEnvWithPath(".env")
}

// NewEnvWithPath creates a new environment
func NewEnvWithPath(configFile string) Env {
	v := enviper.New(viper.New())
	v.SetConfigFile(configFile)

	// The environment variables has higher priority than the values defined in .env file
	env := Env{}
	err := v.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	return env
}

func (e Env) ToDatabaseDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", e.DBUsername, e.DBPassword, e.DBHost, e.DBPort, e.DBName)
}
