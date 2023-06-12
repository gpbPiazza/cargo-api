package main

import "github.com/gpbPiazza/cargo-api/src/presentation"

func main() {
	app := presentation.NewApp()
	app.Start()
}
