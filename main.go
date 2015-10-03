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

	// Analyse self-contained packages.
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

	// Compute scoped fan-out of each package.
	fanOut := make(map[string]int, 1)
	for p, imp := range imports {
		fo := len(imp)
		fanOut[p] = fo
		fmt.Printf("%s has fan-out %d\n", p, fo)
	}

	// Compute scoped fan-in of each package.
	fanIn := make(map[string]int, 1)
	for _, imp := range imports {
		for _, i := range imp {
			fanIn[i] = fanIn[i] + 1
		}
	}

	// Compute scoped instability of each package.
	instab := make(map[string]float32, 1)
	for p, _ := range imports {
		instability := float32(fanOut[p]) / float32((fanIn[p] + fanOut[p]))
		instab[p] = instability
		fmt.Printf("%s has fan-in %d and instability %.2f\n", p, fanIn[p], instability)
	}

	// Check stable dependencies principle violations.
	for p, imp := range imports {
		for _, q := range imp {
			if instab[q] > instab[p] {
				fmt.Printf("%s depends on %s but the former is more stable than the latter\n", p, q)
			}
		}
	}

	// Detect direct dependency cycles.
	for p, imp := range imports {
		for _, i := range imp {
			for _, q := range imports[i] {
				if q == p {
					fmt.Printf("Direct dependency cycle between %s and %s\n", p, i)
				}
			}
		}
	}
}
