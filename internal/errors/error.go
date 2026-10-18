package custom_errors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrBadRequest   = errors.New("bad request")
	ErrExternal     = errors.New("external error")
	ErrInternal     = errors.New("internal error")
	ErrConflict     = errors.New("conflict error")
	ErrNotSupported = errors.New("only github repository URLs are supported")
	ErrInvalidUrl   = errors.New("invalid url")
)

func HTTPStatus(err error) int {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrBadRequest):
		return http.StatusBadRequest
	case errors.Is(err, ErrExternal):
		return http.StatusBadGateway
	case errors.Is(err, ErrInternal):
		return http.StatusInternalServerError
	case errors.Is(err, ErrConflict):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func Respond(c *gin.Context, err error) {
	c.JSON(HTTPStatus(err), gin.H{
		"error": err.Error(),
	})
}
