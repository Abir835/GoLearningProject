package db

import (
	"github.com/jmoiron/sqlx"
)

var readDB *sqlx.DB
var writeDB *sqlx.DB

func GetReadDB() *sqlx.DB {
	return readDB
}

func GetWriteDB() *sqlx.DB {
	return writeDB
}
