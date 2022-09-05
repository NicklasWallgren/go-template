package common

import (
	"time"

	"github.com/NicklasWallgren/go-template/adapters/driven/logger"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

// RequestHandler function.
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler.
func NewRequestHandler(logger logger.Logger) RequestHandler {
	engine := gin.New()
	engine.Use(ginzap.Ginzap(logger.GetZapLogger(), time.RFC3339, true))

	return RequestHandler{Gin: engine}
}
