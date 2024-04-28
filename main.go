package main

import (
	"fmt"

	"github.com/BrondoL/wedding-be/config"
	"github.com/BrondoL/wedding-be/internal/constant"
	"github.com/BrondoL/wedding-be/internal/handler"
	r "github.com/BrondoL/wedding-be/internal/repository"
	"github.com/BrondoL/wedding-be/internal/router"
	"github.com/BrondoL/wedding-be/internal/service"
	"github.com/BrondoL/wedding-be/pkg/cache"
	"github.com/BrondoL/wedding-be/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	config.LoadEnv()
	cfg := config.GetEnv()

	redisCache := cache.NewRedisCache(cfg)
	db := config.GetConn(cfg)
	loggerInstance := logrus.New()
	loggerInstance.SetLevel(logrus.DebugLevel)
	loggerInstance.SetFormatter(&logrus.JSONFormatter{})
	log := logger.NewLogger(loggerInstance)

	attendanceRepository := r.NewAttendanceRepository(&r.ARConfig{DB: db})

	attendanceService := service.NewAttendanceService(&service.ASConfig{
		Cfg:                  cfg,
		Cache:                redisCache,
		AttendanceRepository: attendanceRepository,
	})

	handler := handler.NewHandler(&handler.HandlerConfig{
		Cfg:               cfg,
		AttendanceService: attendanceService,
		Logger:            log,
	})

	if cfg.APP_ENV == constant.EnvironmentProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// // handle CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{cfg.CLIENT_URL, "http://localhost:8000"}
	config.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-Max"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// handle No Route
	r.NoRoute(handler.NoRoute)
	r.GET("/", handler.DefaultHandler)
	r.Use(gin.Recovery())

	version := fmt.Sprintf("/api/%s", cfg.API_VERSION)
	api := r.Group(version)

	router.AttendanceRoute(api, handler)

	port := fmt.Sprintf(":%s", cfg.APP_PORT)
	r.Run(port)
}
