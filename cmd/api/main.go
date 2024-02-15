package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gpbPiazza/cargo-api/src/infrastructure/databases/sql"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/envs"
	"github.com/gpbPiazza/cargo-api/src/presentation/http_api"
)

func main() {
	ctx := context.Background()
	envs := envs.New()
	sql.ConnectDB(ctx)
	app := http_api.NewApp()
	log.Fatal(app.Listen(fmt.Sprintf(":%s", envs.ServerPort)))
}
