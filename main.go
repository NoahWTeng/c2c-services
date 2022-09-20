package main

import (
	app "c2c-services/src"
	"log"
)

func main() {
	application, err := app.CreateNewApp()

	if err != nil {
		log.Fatal(err)
	}

	_ = application.Init()
}
