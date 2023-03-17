package main

import (
	"fmt"

	"{{ .Scaffold.gomod }}/cmd/api/web"
	"{{ .Scaffold.gomod }}/internal/sys/config"

	{{- if .Scaffold.use_database }}
	"context"
	"log"
	"os"
	"path/filepath"

	atlas "ariga.io/atlas/sql/migrate"

	"{{ .Scaffold.gomod }}/internal/data/ent"
	"{{ .Scaffold.gomod }}/internal/data/migrations"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
	{{- end }}
)

var (
	version   = "nightly"
	commit    = "HEAD"
	buildTime = "now"
)

func build() string {
	short := commit
	if len(short) > 7 {
		short = short[:7]
	}

	return fmt.Sprintf("%s, commit %s, built at %s", version, short, buildTime)
}

func main() {
	conf, err := config.NewFromCLI()
	if err != nil {
		panic(err)
	}

	{{- if .Scaffold.use_database -}}
	// =========================================================================
	// Initialize DB Connection

	dbString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		conf.Postgres.User,
		conf.Postgres.Password,
		conf.Postgres.Name,
		conf.Postgres.SSLMode.String(),
		conf.Postgres.Host,
		conf.Postgres.Port,
	)

	db, err := ent.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}

	// =========================================================================
	// Database Migrations

	temp := filepath.Join(os.TempDir(), "migrations")

	err = migrations.Write(temp)
	if err != nil {
		log.Fatal(err)
	}

	dir, err := atlas.NewLocalDir(temp)
	if err != nil {
		log.Fatal(err)
	}

	options := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	}

	// !WARNING: This runs all migrations automatically.
	err = db.Schema.Create(context.Background(), options...)
	if err != nil {
		log.Fatal(err)
	}

	err = os.RemoveAll(temp)
	if err != nil {
		log.Fatal(err)
	}

	defer func(c *ent.Client) {
		err := c.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	{{- end -}}

	// =========================================================================
	// Run API Server

	args := &web.WebArgs{
		Conf:   conf,
		Build:  build(),
		{{ if .Scaffold.use_database -}}
		Client: db,
		{{- end }}
	}

	if err := web.New(args).Start(); err != nil {
		panic(err)
	}
}
