package main

import (
	"github.com/EricMesquita/tcc-golang/internal/app"
	"github.com/EricMesquita/tcc-golang/internal/infra/logger"
	"github.com/EricMesquita/tcc-golang/internal/infra/variables"
	"github.com/KimMachineGun/automemlimit/memlimit"
)

func main() {
	logger.Init(&logger.Option{
		ServiceName:    variables.ServiceName(),
		ServiceVersion: variables.ServiceVersion(),
		Environment:    variables.Environment(),
		LogLevel:       variables.LogLevel(),
	})

	defer logger.Sync()

	memlimit.SetGoMemLimitWithProvider(memlimit.Limit(2048*2028), 1000)

	application := app.Instance()
	application.Start(false)
}
