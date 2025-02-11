// nolint: lll
package router

import (
	"github.com/gin-gonic/gin"

	"github.com/andys920605/hr-system/internal/north/remote/source/handler"
	"github.com/andys920605/hr-system/pkg/errors"
)

type Router struct {
	// middleware
	interceptorHandler gin.HandlerFunc

	// handler
	healthHandler *handler.HealthHandler
}

func NewRouter(
	interceptorHandler gin.HandlerFunc,
	healthHandler *handler.HealthHandler,
) *Router {
	return &Router{
		interceptorHandler: interceptorHandler,
		healthHandler:      healthHandler,
	}
}

func (r *Router) Register(engine *gin.Engine) {
	engine.NoRoute(func(c *gin.Context) {
		_ = c.Error(errors.RouteNotFound)
	})
	engine.GET("/healthz", r.healthHandler.Check)

	// middleware
	normal := engine.Group("/srv", r.interceptorHandler)

	normal.GET("/healthz", r.healthHandler.Check)

}
