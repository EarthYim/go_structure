package gin

import (
	"fmt"
	"go_stucture/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginContext struct {
	*gin.Context
}

func NewGinContext(c *gin.Context) *ginContext {
	return &ginContext{
		Context: c,
	}
}

func (c *ginContext) StatusOK() {
	c.Context.Status(http.StatusOK)
}

type ginRouter struct {
	gin *gin.Engine
}

func NewGinRouter() *ginRouter {
	return &ginRouter{
		gin: gin.Default(),
	}
}

func (r *ginRouter) Run(port string) {
	r.gin.Run(fmt.Sprintf(":%s", port))
}

//handler wrapper

func (r *ginRouter) GET(path string, handler func(core.Context)) {
	r.gin.GET(path, func(c *gin.Context) {
		handler(NewGinContext(c))
	})
}
