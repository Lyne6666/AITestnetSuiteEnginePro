// cmd/aitestnetsuiteenginepro/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"aitestnetsuiteenginepro/internal/aitestnetsuiteenginepro"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := aitestnetsuiteenginepro.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
