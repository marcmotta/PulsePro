// cmd/pulsepro/main.go
package main

import (
	"flag"
	"log"
	"os"

	"pulsepro/internal/pulsepro"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := pulsepro.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
