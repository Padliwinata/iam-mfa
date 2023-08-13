package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Padliwinata/iam-mfa/bin/pkg/errors"
	"github.com/Padliwinata/iam-mfa/bin/pkg/log"

	"github.com/labstack/echo/v4"
)

type response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func getErrorStatusCode(err error) int {
	errString, ok := err.(*errors.ErrorString)
	if ok {
		return errString.Code()
	}

	// default http status code
	return http.StatusInternalServerError
}

type Meta struct {
	Method        string    `json:"method"`
	Url           string    `json:"url"`
	Code          string    `json:"code"`
	ContentLength int64     `json:"content_length"`
	Date          time.Time `json:"date"`
	Ip            string    `json:"ip"`
}

func RespSuccess(c echo.Context, data interface{}, message string) error {
	meta := Meta{
		Date:          time.Now(),
		Url:           c.Path(),
		Method:        c.Request().Method,
		Code:          fmt.Sprintf("%v", http.StatusOK),
		ContentLength: c.Request().ContentLength,
		Ip:            c.RealIP(),
	}
	byteMeta, _ := json.Marshal(meta)
	log.GetLogger().Info("service-info", "Logging service...", "audit-log", string(byteMeta))
	return c.JSON(http.StatusOK, response{
		Message: message,
		Data:    data,
		Code:    http.StatusOK,
		Success: true,
	})
}

func RespError(c echo.Context, err error) error {
	meta := Meta{
		Date:          time.Now(),
		Url:           c.Path(),
		Method:        c.Request().Method,
		Code:          fmt.Sprintf("%v", getErrorStatusCode(err)),
		Ip:            c.RealIP(),
		ContentLength: c.Request().ContentLength,
	}
	byteMeta, _ := json.Marshal(meta)

	log.GetLogger().Error("service-error", "Logging service...", "audit-log", string(byteMeta))

	return c.JSON(getErrorStatusCode(err), response{
		Data:    nil,
		Message: err.Error(),
		Code:    getErrorStatusCode(err),
		Success: false,
	})
}
