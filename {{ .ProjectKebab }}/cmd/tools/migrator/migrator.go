package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"{{ .Scaffold.gomod }}/internal/data/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("internal/data/migrations/sql")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Prompt user for migration name.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Migration Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed reading migration name: %v", err)
	}

	name = strings.TrimSpace(name)

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	err = migrate.NamedDiff(ctx, "postgresql://user:pass@127.0.0.1:55102/test?sslmode=disable", name, opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}

	fmt.Println("Migration file generated successfully.")
}
