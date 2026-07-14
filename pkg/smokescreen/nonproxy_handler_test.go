//go:build !nounit
// +build !nounit

package smokescreen

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildProxyUsesConfiguredNonproxyHandler(t *testing.T) {
	conf := NewConfig()
	conf.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	recorder := httptest.NewRecorder()
	BuildProxy(conf).ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/", nil))

	assert.Equal(t, http.StatusNoContent, recorder.Code)
}
