package custom_errors

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestHTTPStatus(t *testing.T){
	tests := []struct{
		name string
		err error
		expected int
	}{
		{"not found", ErrNotFound, http.StatusNotFound},
		{"bad request", ErrBadRequest, http.StatusBadRequest},
		{"conflict", ErrConflict, http.StatusConflict},
		{"external", ErrExternal, http.StatusBadGateway},
		{"internal", ErrInternal, http.StatusInternalServerError},
		{"unknown", errors.New("random"), http.StatusInternalServerError},
		{
			"wrapped not found",
			fmt.Errorf("%w: repo not found", ErrNotFound),
			http.StatusNotFound,
		},
	}
	for _, v := range tests{
		t.Run(v.name, func(t *testing.T){
			got := HTTPStatus(v.err)
			if got != v.expected{
				t.Errorf("got %d want %d", got, v.expected)
			}
		})
	}
}