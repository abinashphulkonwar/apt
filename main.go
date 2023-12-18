package main

import (
	"log"
	"os"

	"github.com/abinashphulkonwar/apt/services"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "apt",
		Usage:  "fetch data",
		Action: services.HandlerRoot,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "get",
				Usage:   "path for fetch data",
				Aliases: []string{"g"},
				Action:  services.Handler,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
