package config

type HTTPServerConfig struct {
	Port string
}

func NewHTTPServerConfig(port string) *HTTPServerConfig {
	return &HTTPServerConfig{Port: port}
}
