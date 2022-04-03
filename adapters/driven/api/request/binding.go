package request

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// ShouldBindUri binds the passed struct pointer using the specified binding engine.
func ShouldBindUri(c *gin.Context, obj interface{}) error {
	m := make(map[string][]string)
	for _, v := range c.Params {
		m[v.Key] = []string{v.Value}
	}

	return binding.Uri.BindUri(m, obj)
}
