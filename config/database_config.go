package config

type DatabaseConfig struct {
	Name               string
	User               string
	Password           string
	Host               string
	Port               string
	MigrationDirectory string
	Driver             string // TODO, validate supported dialects
}

func NewDatabaseConfig(
	name string,
	user string,
	password string,
	host string,
	port string,
	migrationDirectory string,
	dialect string,
) *DatabaseConfig {
	return &DatabaseConfig{
		Name:               name,
		User:               user,
		Password:           password,
		Host:               host,
		Port:               port,
		MigrationDirectory: migrationDirectory,
		Driver:             dialect,
	}
}
