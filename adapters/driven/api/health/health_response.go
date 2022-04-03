package health

type HealthResponse struct {
	Health     string
	Components []Health
}

type Health struct {
	Status  string
	Name    string
	Details map[string]any
}
