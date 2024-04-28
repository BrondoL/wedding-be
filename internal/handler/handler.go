package handler

import (
	"github.com/BrondoL/wedding-be/config"
	s "github.com/BrondoL/wedding-be/internal/service"
	"github.com/BrondoL/wedding-be/pkg/logger"
)

type Handler struct {
	cfg               config.Config
	attendanceService s.AttendanceService
	logger            logger.AppLogger
}

type HandlerConfig struct {
	Cfg               config.Config
	AttendanceService s.AttendanceService
	Logger            logger.AppLogger
}

func NewHandler(c *HandlerConfig) *Handler {
	return &Handler{
		cfg:               c.Cfg,
		attendanceService: c.AttendanceService,
		logger:            c.Logger,
	}
}
