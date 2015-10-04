package main

import (
	"fmt"
	"os"

	"strings"

	"github.com/codegangsta/cli"
	"github.com/glyn/sago/disc"
)

func main() {
	app := cli.NewApp()
	app.Name = "sago"
	app.Usage = "Static Analysis for Go modules"
	app.Action = func(c *cli.Context) {
		println("Try sago help")
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
		fanOut[p] = len(imp)
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
		instab[p] = instability(fanIn[p], fanOut[p])
		fmt.Printf("%s has fan-in %d, fan-out %d, and instability %.2f\n", d.ShortName(p), fanIn[p], fanOut[p], instab[p])
	}

	// Check stable dependencies principle violations.
	for p, imp := range imports {
		for _, q := range imp {
			if instab[q] > instab[p] {
				fmt.Printf("%s depends on the less stable %s\n", d.ShortName(p), d.ShortName(q))
			}
		}
	}

	// Detect direct dependency cycles.
	for p, imp := range imports {
		for _, i := range imp {
			for _, q := range imports[i] {
				if q == p {
					fmt.Printf("Direct dependency cycle between %s and %s\n", d.ShortName(p), d.ShortName(i))
				}
			}
		}
	}
}

func instability(fanIn, fanOut int) float32 {
	if fan := fanIn + fanOut; fan > 0 {
		return float32(fanOut) / float32(fan)
	}
	return 0
}
