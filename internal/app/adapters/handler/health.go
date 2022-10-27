package handler

import (
	"github.com/EricMesquita/tcc-golang/internal/app/domain"
	"github.com/EricMesquita/tcc-golang/internal/app/domain/health"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/presenter"
	"github.com/EricMesquita/tcc-golang/internal/infra/logger"
	"github.com/EricMesquita/tcc-golang/internal/infra/logger/attributes"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type HealthHandler struct {
	services *domain.Services
}

func NewHealthHandler(services *domain.Services) *HealthHandler {
	return &HealthHandler{
		services: services,
	}
}

func (h *HealthHandler) Configure(server *echo.Echo) {
	server.GET("/health", h.health)
}

func (h *HealthHandler) health(c echo.Context) error {
	ctx := c.Request().Context()

	var read, write, redis *health.Health
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	go func() {
		defer waitGroup.Done()

		read = h.services.Health.HealthReadDB(ctx)
		if read.Error != nil {
			logger.Error(ctx, "Health check error on read database", attributes.New().WithError(read.Error))
		}
	}()

	go func() {
		defer waitGroup.Done()

		write = h.services.Health.HealthWriteDB(ctx)
		if write.Error != nil {
			logger.Error(ctx, "Health check error on write database", attributes.New().WithError(write.Error))
		}
	}()

	go func() {
		defer waitGroup.Done()

		redis = h.services.Health.HealthRedisDB(ctx)
		if redis.Error != nil {
			logger.Error(ctx, "Health check error on redis database", attributes.New().WithError(redis.Error))
		}
	}()

	waitGroup.Wait()

	response := presenter.HealthResponse(read, write, redis)
	if response.Healthy {
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusServiceUnavailable, response)
}
