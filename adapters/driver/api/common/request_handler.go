package common

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/logger"
	"github.com/gin-gonic/gin"
)

// RequestHandler function.
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler.
func NewRequestHandler(logger logger.Logger) RequestHandler {
	gin.DefaultWriter = logger.GetGinLogger()
	engine := gin.New()

	return RequestHandler{Gin: engine}
}
