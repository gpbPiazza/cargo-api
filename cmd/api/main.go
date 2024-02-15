package main

import "github.com/gpbPiazza/cargo-api/src/presentation/http_api"

func main() {
	app := http_api.NewApp()
	app.Start()
}
