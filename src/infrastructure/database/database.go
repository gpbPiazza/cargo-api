package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/gpbPiazza/cargo-api/src/infrastructure/env"
)

// https://pkg.go.dev/database/sql
// https://github.com/golang/go/wiki/SQLDrivers

var globalDB *sql.DB

func Init() {
	envs := env.Envs()
	driverName := fmt.Sprintf("driver-%s", envs.Database.Name)

	db, err := sql.Open(driverName, envs.Database.Name)
	if err != nil {
		log.Fatal(err)
	}

	globalDB = db
}

func Ping(ctx context.Context) error {
	// ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()

	if err := globalDB.PingContext(ctx); err != nil {
		return err
	}

	return nil
}
