package util

import (
	"errors"
	"net/http"

	"github.com/BrondoL/wedding-be/internal/dto/response"
)

func ResponseSuccess(data interface{}, msg string, code ...int) response.JsonResponse {
	resCode := http.StatusOK
	if len(code) > 0 {
		resCode = code[0]
	}

	return responseJSON(data, msg, nil, resCode)
}

func ResponseError(err error) response.JsonResponse {
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
			if ew.Err != nil {
				return responseJSON(nil, ew.Err.Error(), ew.Validation, resCode)
			}
			return responseJSON(nil, ew.Message, ew.Validation, resCode)
		}

		return responseJSON(nil, MsgServerError, nil, resCode)
	}

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
