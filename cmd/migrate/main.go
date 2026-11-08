package main

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/zrajesh/olx-backend/internal/config"
)

func main () {
	if (len(os.Args) < 2) {
		log.Fatal("Usage: make migrate <up|down>")
		return
	}

	cfg := config.MustLoad()

	m, err := migrate.New(
        "file://migrations",
        cfg.DatabaseURL)

	if err != nil {
		log.Fatalf("migration.new: %v", err)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Usage: make migrate-force v=<version> (use -1 to reset to no version)")
		}
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("invalid version %q: %v", os.Args[2], err)
		}
		if err := m.Force(version); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}
}