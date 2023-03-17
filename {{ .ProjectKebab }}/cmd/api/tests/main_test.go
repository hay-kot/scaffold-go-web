package integration_test

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"testing"

	"{{ .Scaffold.gomod }}/cmd/api/web"
	"{{ .Scaffold.gomod }}/internal/sys/config"
	{{- if .Scaffold.use_database }}
	"{{ .Scaffold.gomod }}/internal/data/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
	{{ end -}}
)

func newTestServer(t *testing.T) (base string, shutdown func()) {
	randomPort := rand.Intn(60000) + 8000

	{{- if .Scaffold.use_database -}}
	// =========================================================================
	// Database Migrations
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer func() {
		_ = client.Close()
	}()
	{{- end }}

	config := config.Config{
		Mode:     config.Mode("production"),
		LogLevel: "debug",
		Web: config.WebConfig{
			Host:           "127.0.0.1",
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
		{{- if .Scaffold.use_database -}}
		Client: client,
		{{ end -}}
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
		resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/status", randomPort))
		if err == nil && resp.StatusCode == 200 {
			_ = resp.Body.Close()
			break
		}
	}

	return fmt.Sprintf("http://127.0.0.1:%d", randomPort), func() {}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
