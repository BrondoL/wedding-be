package handler

import (
	"github.com/BrondoL/wedding-be/config"
	s "github.com/BrondoL/wedding-be/internal/service"
)

type Handler struct {
	cfg               config.Config
	attendanceService s.AttendanceService
}

type HandlerConfig struct {
	Cfg               config.Config
	AttendanceService s.AttendanceService
}

func NewHandler(c *HandlerConfig) *Handler {
	return &Handler{
		cfg:               c.Cfg,
		attendanceService: c.AttendanceService,
	}
}
