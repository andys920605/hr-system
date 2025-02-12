package main

import (
	"time"

	"github.com/gin-contrib/cors"

	"github.com/andys920605/hr-system/cmd/api/router"
	"github.com/andys920605/hr-system/cmd/injection"
	"github.com/andys920605/hr-system/internal/north/remote/source/handler"
	"github.com/andys920605/hr-system/pkg/http"
	"github.com/andys920605/hr-system/pkg/http/middleware"
)

func main() {
	i := injection.New()

	// handler
	healthHandler := handler.NewHealthHandler()
	EmployeeHandler := handler.NewEmployeeHandler(i.EmployeeAppService)

	// api server
	server := http.NewServer(i.Logger, i.Config.Server.Name)

	// middleware
	interceptorHandler := middleware.NewInterceptor().Handler()

	server.RegisterCORS(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})

	server.RegisterRouter(router.NewRouter(
		interceptorHandler,
		healthHandler,
		EmployeeHandler,
	))
	server.Run(i.Config.Server.Port)
}
