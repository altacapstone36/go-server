package middleware

import (
	"go-hospital-server/internal/utils/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)


func Logging(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${status} ${method} ${host}${path} ${latency_human}`,
		Output: &logger.LogDriver,
	}))
	e.Logger.SetLevel(log.DEBUG)
}

