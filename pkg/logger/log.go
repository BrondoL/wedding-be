package logger

import "github.com/sirupsen/logrus"

type AppLogger interface {
	Debug(message string, fields map[string]interface{})
	Info(message string, fields map[string]interface{})
	Warn(message string, fields map[string]interface{})
	Error(message string, fields map[string]interface{})
	Fatal(message string, fields map[string]interface{})
}

func InitLogger() *logrus.Logger {
	loggerInstance := logrus.New()
	loggerInstance.SetFormatter(&logrus.JSONFormatter{})

	return loggerInstance
}
