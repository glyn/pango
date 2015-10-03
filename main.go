package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/glyn/gomod/disc"
)

func main() {
	app := cli.NewApp()
	app.Name = "gomod"
	app.Usage = "discover and analyse Go modules"
	app.Action = func(c *cli.Context) {
		println("Try gomod help")
	}

	app.Commands = []cli.Command{
		{
			Name:    "discover",
			Aliases: []string{"d"},
			Usage:   "discover Go modules",
			Action:  disc.New().Discover,
		},
	}

	app.Run(os.Args)
}
