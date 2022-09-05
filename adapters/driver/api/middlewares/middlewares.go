package middlewares

// Middleware interface.
type Middleware interface {
	Setup()
}

// Middlewares contains multiple middleware.
type Middlewares []Middleware

// NewMiddlewares creates new middlewares
// nolint: interfacer
func NewMiddlewares(middlewares []Middleware) Middlewares {
	return middlewares
}

// Setup sets up middlewares.
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
