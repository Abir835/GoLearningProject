package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-learning-project/config"
	"log/slog"
	"os"
	"time"
)

func connect(dbConfig config.DBConfig) *sqlx.DB {
	dbSource := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	if !dbConfig.EnableSSLMode {
		dbSource += " sslmode=disable"
	}

	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dbCon.SetConnMaxIdleTime(
		time.Duration(dbConfig.MaxIdleTimeInMinute * int(time.Minute)),
	)

	return dbCon
}

func ConnectDB() {
	conf := config.GetConfig()

	readDB = connect(conf.DB.Read)
	slog.Info("Connected to read database")

	writeDB = connect(conf.DB.Write)
	slog.Info("Connected to write database")
}

func CloseDB() {
	if err := readDB.Close(); err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("Disconnected from read database")

	if err := writeDB.Close(); err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("Disconnected from write database")
}
