package integration_test

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"testing"

	"{{ .Scaffold.gomod }}/cmd/api/web"
)

func newTestServer(t *testing.T) (base string, shutdown func()) {
	randomPort := rand.Intn(60000) + 8000

	config := web.Config{
		Mode:     web.ModeProduction,
		LogLevel: "debug",
		Web: web.WebConfig{
			Host:           "localhost",
			Port:           fmt.Sprintf("%d", randomPort),
			AllowedOrigins: "*",
			IdleTimeout:    30,
			ReadTimeout:    10,
			WriteTimeout:   10,
		},
	}

	args := &web.WebArgs{
		Conf:  &config,
		Build: "test",
	}

	go func() {
		site := web.New(args)
		t.Cleanup(func() {
			_ = site.Shutdown("test finished")
		})

		_ = site.Start()
	}()

	// wait for server to start by trying to ping /ping
	for {
		resp, err := http.Get(fmt.Sprintf("http://localhost:%d/status", randomPort))
		if err == nil && resp.StatusCode == 200 {
			_ = resp.Body.Close()
			break
		}
	}

	return fmt.Sprintf("http://localhost:%d", randomPort), func() {}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
