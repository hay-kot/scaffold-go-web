package integration_test

import (
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestPing(t *testing.T) {
	baseurl, shutdown := newTestServer(t)
	t.Cleanup(shutdown)

	e := httpexpect.Default(t, baseurl)

	e.GET("/status").
		Expect().
		Status(200).
		JSON().
		Object().
		HasValue("status", "ok").
		HasValue("build", "test")
}
