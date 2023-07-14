package main

import (
	"fmt"
	"github/cam-per/nst"

	"github.com/urfave/cli"
)

func search(ctx *cli.Context) error {
	title := ctx.Args().First()
	headers := nst.Search(title)

	for _, book := range headers {
		fmt.Println(book.Name)
	}

	return nil
}
