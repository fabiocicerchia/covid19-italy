package main

import (
	"log"
	"os"

	clisetup "covid19.fabiocicerchia.it/cli/pkg/factory/clisetup"
)

func main() {
	app := clisetup.Config()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
