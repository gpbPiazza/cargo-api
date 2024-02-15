package sql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gpbPiazza/cargo-api/src/infrastructure/envs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var gPool *pgxpool.Pool

func ConnectDB(ctx context.Context) {
	if gPool != nil {
		return
	}

	poolConfig, err := pgxpool.ParseConfig(connString())
	if err != nil {
		log.Fatalln("unable to parse connString:", err)
	}

	setPoolConfig(poolConfig)

	connPool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalln("unable to create connection pool:", err)
	}

	if err := connPool.Ping(ctx); err != nil {
		log.Fatalln("unable to ping database:", err)
	}

	gPool = connPool
}

func connString() string {
	dbEnvs := envs.New().Database
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnvs.Username, dbEnvs.Password, dbEnvs.Host, dbEnvs.Port, dbEnvs.Name)
}

func setPoolConfig(poolConfig *pgxpool.Config) {
	dbEnvs := envs.New().Database
	poolConfig.MaxConnIdleTime = time.Duration(dbEnvs.MaxIdleConns)
	poolConfig.MaxConns = int32(dbEnvs.MaxOpenConns)
}
