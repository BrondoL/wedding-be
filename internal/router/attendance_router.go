package router

import (
	"github.com/BrondoL/wedding-be/internal/handler"
	"github.com/gin-gonic/gin"
)

func AttendanceRoute(r *gin.RouterGroup, h *handler.Handler) {
	r.GET("/attendances", h.GetAttendances)
	r.POST("/attendances", h.CreateAttendance)
}
