package util

import (
	"errors"
	"net/http"

	"github.com/BrondoL/wedding-be/internal/dto/response"
	"github.com/BrondoL/wedding-be/pkg/logger"
	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, logger logger.AppLogger, data interface{}, msg string, code ...int) response.JsonResponse {
	resCode := http.StatusOK
	if len(code) > 0 {
		resCode = code[0]
	}

	logger.Info(msg, map[string]interface{}{
		"code":   resCode,
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	})

	return responseJSON(data, msg, nil, resCode)
}

func ResponseError(c *gin.Context, err error, logger logger.AppLogger) response.JsonResponse {
	resCode := http.StatusInternalServerError
	var ew *ErrorWrapper
	if errors.As(err, &ew) {
		switch ew.Code {
		case CodeClientError:
			resCode = http.StatusBadRequest
		case CodeNotFoundError:
			resCode = http.StatusNotFound
		case CodeConflictError:
			resCode = http.StatusConflict
		case CodeClientUnauthorized:
			resCode = http.StatusUnauthorized
		case CodeClientForbidden:
			resCode = http.StatusForbidden
		case CodeUnprocessableEntity:
			resCode = http.StatusUnprocessableEntity
		}

		if resCode != http.StatusInternalServerError {
			LogAppErr(c, ew, logger, resCode)

			if ew.Err != nil {
				return responseJSON(nil, ew.Err.Error(), ew.Validation, resCode)
			}
			return responseJSON(nil, ew.Message, ew.Validation, resCode)
		}

		LogAppErr(c, ew, logger, resCode)

		return responseJSON(nil, MsgServerError, nil, resCode)
	}

	logger.Error(err.Error(), map[string]interface{}{
		"code":   resCode,
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	})

	return responseJSON(nil, MsgServerError, nil, resCode)
}

func responseJSON(data interface{}, msg string, errors interface{}, code int) response.JsonResponse {
	return response.JsonResponse{
		Code:    code,
		Message: msg,
		Errors:  errors,
		Data:    data,
	}
}
