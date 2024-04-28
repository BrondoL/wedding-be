package util

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/BrondoL/wedding-be/pkg/logger"
	"github.com/gin-gonic/gin"
)

/*
Logs error from outermost to innermost
Still could be improved, made just to make debugging this app easier
*/
func LogAppErr(c *gin.Context, ew *ErrorWrapper, logger logger.AppLogger, code int) {
	msg := make(map[string]interface{})

	if ew == nil {
		return
	}

	setMessage(msg, ew)

	if ew.Err != nil && !errors.As(ew.Err, &ew) {
		msg["error"] = ew.Err.Error()
	}

	msg["code"] = code
	msg["method"] = c.Request.Method
	msg["path"] = c.Request.URL.Path

	text, _ := msg["message"].(string)
	delete(msg, "message")
	if code != http.StatusInternalServerError {
		logger.Warn(text, msg)
	} else {
		logger.Error(text, msg)
	}
}

func setMessage(msg map[string]interface{}, ew *ErrorWrapper) {
	msg["message"] = ew.Message
	msg["filename"] = ew.Filename
	msg["line"] = strconv.Itoa(ew.LineNumber)
}
