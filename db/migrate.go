package db

import (
	_ "github.com/lib/pq"
	"go-learning-project/config"
	"log/slog"
	"os"

	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB() {
	conf := config.GetConfig()

	// try migrating tables
	migrations := &migrate.FileMigrationSource{
		Dir: conf.MigrationSource,
	}

	_, err := migrate.Exec(writeDB.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("Successfully migrated database")
}
