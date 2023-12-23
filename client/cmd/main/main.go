package main

import (
	"log"

	"github.com/BelyaevEI/GophKeeper/client/internal/app"
)

func main() {

	// Init application
	app, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.RunServer(); err != nil {
		log.Fatal(err)
	}

}
