package sql

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ctxConnKey struct{}

func SetConnCtx(ctx context.Context, cPool *pgxpool.Pool) context.Context {
	if cPool != nil {
		return context.WithValue(ctx, ctxConnKey{}, cPool)
	}

	return context.WithValue(ctx, ctxConnKey{}, gPool)
}

func Conn(ctx context.Context) *pgxpool.Pool {
	cPool, ok := ctx.Value(ctxConnKey{}).(*pgxpool.Pool)

	if !ok {
		log.Fatal("ctx has not connection pool with db")
	}

	return cPool
}
