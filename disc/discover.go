package disc

import (
	"fmt"
	"go/parser"
	"path/filepath"

	"os"

	"go/token"

	"strings"

	"github.com/codegangsta/cli"
)

type disc struct {
	fset   *token.FileSet
	gopath string
}

func New() *disc {
	return &disc{
		fset:   token.NewFileSet(),
		gopath: os.Getenv("GOPATH"),
	}
}

func (d *disc) Discover(c *cli.Context) {
	rootPkg := c.Args().First()
	d.traverse(rootPkg)
}

func (d *disc) traverse(p string) {
	imports := d.imports(p)
	fmt.Printf("Imports: %v\n", imports)
	for _, imp := range imports {
		if subpackage(imp, p) {
			d.traverse(imp)
		}
	}
}

func subpackage(p, q string) bool {
	return strings.HasPrefix(p, q)
}

func (d *disc) imports(pkg string) []string {
	i := []string{}
	pkgDir := filepath.Join(d.gopath, "src", pkg)
	pAsts, err := parser.ParseDir(d.fset, pkgDir, nil, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}
	for _, ast := range pAsts {
		for _, f := range ast.Files {
			for _, is := range f.Imports {
				i = append(i, strings.Trim(is.Path.Value, `"`))
			}
		}
	}
	return i
}
