package server

import (
	"fmt"
	"github.com/EricMesquita/tcc-golang/internal/infra/server/middleware"
	"github.com/EricMesquita/tcc-golang/internal/infra/variables"
	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func New() (e *echo.Echo) {
	e = echo.New()

	// Configure request
	e.Use(middleware.ConfigRequest())

	// Configure cors
	e.Use(middleware.ConfigCors())

	// Configure Timeout
	e.Use(middleware.ConfigTimeout())

	// Configure Recover Timeout
	e.Use(echoMiddleware.Recover())

	e.Use(echoPrometheus.MetricsMiddleware())
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.Server.Addr = fmt.Sprintf("%s:%d", variables.ServerHost(), variables.ServerPort())

	return e
}
