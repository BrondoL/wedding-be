package logger

import "github.com/sirupsen/logrus"

type appLogger struct {
	log *logrus.Logger
}

func NewLogger(logger *logrus.Logger) AppLogger {
	return &appLogger{
		log: logger,
	}
}

func (al *appLogger) Debug(message string, fields map[string]interface{}) {
	al.log.WithFields(fields).Debug(message)
}

func (al *appLogger) Info(message string, fields map[string]interface{}) {
	al.log.WithFields(fields).Info(message)
}

func (al *appLogger) Warn(message string, fields map[string]interface{}) {
	al.log.WithFields(fields).Warn(message)
}

func (al *appLogger) Error(message string, fields map[string]interface{}) {
	al.log.WithFields(fields).Error(message)
}

func (al *appLogger) Fatal(message string, fields map[string]interface{}) {
	al.log.WithFields(fields).Fatal(message)
}
