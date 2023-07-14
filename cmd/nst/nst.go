package main

import (
	"github.com/urfave/cli"
)

var app *cli.App

func configure() {
	app = &cli.App{
		Name:        "nst",
		Description: "Novel Saving Tool ",

		Commands: []cli.Command{
			{
				Name:    "search",
				Aliases: []string{"s", "search"},
				Action:  search,
			},
		},
	}
	app.UseShortOptionHandling = true
}
