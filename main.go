package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/glyn/pango/disc"
	. "github.com/glyn/pango/pkg"
)

func main() {
	app := cli.NewApp()
	app.Name = "pango"
	app.Usage = "Package Analysis for Go"
	app.Action = func(c *cli.Context) {
		println("Try pango help")
	}

	app.Commands = []cli.Command{
		{
			Name:    "analyse",
			Aliases: []string{"a"},
			Usage:   "Analyses a Go package and its dependencies",
			Action:  analyse,
		},
	}

	app.Run(os.Args)
}

func analyse(c *cli.Context) {
	d := disc.New(c.Args().Get(0))

	root := Pkg(c.Args().Get(1))
	var imports PGraph
	if root != "" {
		imports = d.Discover(root)
	} else {
		imports = d.DiscoverAll()
	}

	// Analyse self-contained packages.
	imports.Packages().Walk(func(p Pkg) {
		escape := false
		d.Walk(p, func(q Pkg, qi PSet) {
			qi.Walk(func(i Pkg) {
				if !p.HasSubpackage(i) {
					escape = true
				}
			})
		})
		if !escape {
			fmt.Printf("%s is self-contained\n", d.ShortName(p))
		}
	})

	// Compute scoped fan-out of each package.
	fanOut := make(map[Pkg]int, 1)
	imports.Walk(func(p Pkg, i PSet) {
		fanOut[p] = i.Size()
	})

	// Compute scoped fan-in of each package.
	fanIn := make(map[Pkg]int, 1)
	imports.Walk(func(_ Pkg, i PSet) {
		i.Walk(func(q Pkg) {
			fanIn[q] = fanIn[q] + 1
		})
	})

	// Compute scoped instability of each package.
	instab := make(map[Pkg]float32, 1)
	imports.Walk(func(p Pkg, i PSet) {
		instab[p] = instability(fanIn[p], fanOut[p])
		fmt.Printf("%s has fan-in %d, fan-out %d, and instability %.2f\n", d.ShortName(p), fanIn[p], fanOut[p], instab[p])
	})

	// Check stable dependencies principle violations.
	imports.Walk(func(p Pkg, i PSet) {
		i.Walk(func(q Pkg) {
			if instab[q] > instab[p] {
				fmt.Printf("%s depends on the less stable %s\n", d.ShortName(p), d.ShortName(q))
			}
		})
	})

	// Detect direct dependency cycles.
	imports.Walk(func(p Pkg, imp PSet) {
		imp.Walk(func(i Pkg) {
			if q, ok := imports.Imports(i); ok {
				if q.Contains(p) {
					fmt.Printf("Direct dependency cycle between %s and %s\n", d.ShortName(p), d.ShortName(i))
				}
			}
		})
	})
}

func instability(fanIn, fanOut int) float32 {
	if fan := fanIn + fanOut; fan > 0 {
		return float32(fanOut) / float32(fan)
	}
	return 0
}
