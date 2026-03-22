package helper

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imanudd/inventorySvc-clean-architecture/constant"
	"github.com/imanudd/inventorySvc-clean-architecture/pkg/validator"
)

type JSONResponse struct {
	Code       int         `json:"code,omitempty"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, code int, data ...interface{}) {
	hte := JSONResponse{
		StatusCode: code,
		Message:    http.StatusText(code),
	}

	if len(data) > 0 {
		hte.Data = data[0]
	}

	c.JSON(code, hte)
}

func Error(c *gin.Context, code int, message constant.ErrorMessage) {
	hte := JSONResponse{
		Code:       code,
		StatusCode: code,
		Message:    string(message),
	}

	c.JSON(code, hte)
}

func InternalError(c *gin.Context, err error) {
	response := JSONResponse{
		Code:       http.StatusInternalServerError,
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	}

	if errors.Is(err, &validator.ValidationErrors{}) {
		response.Code = http.StatusBadRequest
		response.StatusCode = http.StatusBadRequest
		response.Message = err.Error()
	}

	c.JSON(response.Code, response)
}
