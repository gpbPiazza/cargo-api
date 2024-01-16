package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gpbPiazza/cargo-api/src/infrastructure/envs"
)

var globalDB *sql.DB

func New() *sql.DB {
	if globalDB != nil {
		return globalDB
	}

	db, err := sql.Open("postgres", dataSourceName())
	if err != nil {
		log.Fatal(fmt.Errorf("error on connect to database: %s", err.Error()))
	}

	dbEnvs := envs.New().Database
	db.SetConnMaxIdleTime(time.Duration(dbEnvs.MaxIdleConns))
	db.SetMaxOpenConns(dbEnvs.MaxOpenConns)

	globalDB = db

	return globalDB
}

func dataSourceName() string {
	dbEnvs := envs.New().Database
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnvs.Username, dbEnvs.Password, dbEnvs.Host, dbEnvs.Port, dbEnvs.Name)
}
