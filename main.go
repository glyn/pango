package main

import (
	"fmt"
	"os"

	"strings"

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
			Action:  discover,
		},
	}

	app.Run(os.Args)
}

func discover(c *cli.Context) {
	d := disc.New(c.Args().First(), c.Args().Get(1))
	imports := d.Discover()
	for p, _ := range imports {
		escape := false
		d.Walk(p, func(q string, qi []string) {
			for _, i := range qi {
				if !strings.HasPrefix(i, p) {
					escape = true
				}
			}
		})
		if !escape {
			fmt.Printf("%s is self-contained\n", d.ShortName(p))
		}
	}
}
