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

// ConnectDB star the connection with relational data base and configuration of;
// it returns a new value of the passed ctx with the connections pool seted,
// use SetConnCtx and Conn functions to use the pool conection with the ctx.
func ConnectDB(ctx context.Context) *pgxpool.Pool {
	if gPool != nil {
		return gPool
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

	return gPool
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
