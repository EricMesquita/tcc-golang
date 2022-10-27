package main

import (
	"github.com/EricMesquita/tcc-golang/internal/app"
	"github.com/EricMesquita/tcc-golang/internal/infra/logger"
	"github.com/EricMesquita/tcc-golang/internal/infra/variables"
)

func main() {
	logger.Init(&logger.Option{
		ServiceName:    variables.ServiceName(),
		ServiceVersion: variables.ServiceVersion(),
		Environment:    variables.Environment(),
		LogLevel:       variables.LogLevel(),
	})

	defer logger.Sync()

	application := app.Instance()
	application.Start(false)
}
