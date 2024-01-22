package sql

import (
	"context"
	"database/sql"
	"log"
)

type ctxKey struct{}

func SetConnToCtx(ctx context.Context, conn *sql.Conn) context.Context {
	if conn != nil {
		return context.WithValue(ctx, ctxKey{}, conn)
	}

	gConn, err := globalDB.Conn(ctx)
	if err != nil {
		log.Fatal("error on get Conn From globalDB")
	}

	return context.WithValue(ctx, ctxKey{}, gConn)
}

func ConnFromCtx(ctx context.Context) *sql.Conn {
	conn, ok := ctx.Value(ctxKey{}).(*sql.Conn)

	if !ok {
		log.Fatal("ctx has not connection db")
	}

	return conn
}
