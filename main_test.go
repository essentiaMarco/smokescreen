package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeploymentHandler(t *testing.T) {
	tests := []struct {
		path       string
		wantStatus int
		wantBody   string
	}{
		{path: "/", wantStatus: http.StatusOK, wantBody: "OK\n"},
		{path: "/healthcheck", wantStatus: http.StatusOK, wantBody: "OK\n"},
		{path: "/healthz", wantStatus: http.StatusOK, wantBody: "OK\n"},
		{path: "/unknown", wantStatus: http.StatusNotFound, wantBody: "404 page not found\n"},
	}

	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			deploymentHandler(recorder, httptest.NewRequest(http.MethodGet, test.path, nil))

			assert.Equal(t, test.wantStatus, recorder.Code)
			assert.Equal(t, "text/plain; charset=utf-8", recorder.Header().Get("Content-Type"))
			assert.Equal(t, test.wantBody, recorder.Body.String())
		})
	}
}
