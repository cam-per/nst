package main

import (
	"log"
	"os"

	"github/cam-per/nst"
	"github/cam-per/nst/config"
	_ "github/cam-per/nst/pkg/sources/ranobes"
)

func main() {
	err := config.Load("")
	if err != nil {
		log.Fatal(err)
	}

	err = nst.InitSelenium()
	if err != nil {
		log.Fatal(err)
	}
	defer nst.SeleniumStop()

	configure()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
