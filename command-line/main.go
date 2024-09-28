package main

import (
	"command-line/app"
	"log"
	"os"
)

func main() {
	application := app.App()

	err := application.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}