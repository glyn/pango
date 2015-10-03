package disc

import (
	"fmt"
	"go/parser"
	"path/filepath"

	"os"

	"go/token"

	"strings"
)

type disc struct {
	scope, root string
	fset        *token.FileSet
	gopath      string
}

func New(scope, root string) *disc {
	return &disc{
		scope:  scope,
		root:   root,
		fset:   token.NewFileSet(),
		gopath: os.Getenv("GOPATH"),
	}
}

func (d *disc) Discover() {
	d.walk(d.root, func(p string, imp []string) {
		fmt.Printf("Package %s imports: %v.\n", p, imp)
	})
}

func (d *disc) walk(p string, visit func(string, []string)) {
	d.traverse(p, make(map[string]bool, 1), visit)
}

func (d *disc) traverse(p string, visited map[string]bool, visit func(string, []string)) {
	if visited[p] {
		return
	}
	visited[p] = true

	imports := d.imports(p)
	shortNames := d.shortNames(imports)
	visit(d.shortName(p), shortNames)
	for _, imp := range imports {
		if subpackage(imp, d.scope) {
			d.traverse(imp, visited, visit)
		}
	}
}

// returns true if and only if p is a subpackage of q
func subpackage(p, q string) bool {
	return strings.HasPrefix(p, q)
}

func (d *disc) shortName(p string) string {
	if !subpackage(p, d.scope) {
		panic("failed")
	}
	return p[len(d.scope)+1:]
}

func (d *disc) shortNames(p []string) []string {
	result := []string{}
	for _, q := range p {
		s := d.shortName(q)
		result = append(result, s)
	}
	return result
}

func (d *disc) imports(pkg string) []string {
	i := make(map[string]bool, 1)
	pkgDir := filepath.Join(d.gopath, "src", pkg)
	pAsts, err := parser.ParseDir(d.fset, pkgDir, nil, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}
	for _, ast := range pAsts {
		for _, f := range ast.Files {
			for _, is := range f.Imports {
				q := strings.Trim(is.Path.Value, `"`)
				if q != pkg && subpackage(q, d.scope) {
					i[q] = true
				}
			}
		}
	}
	result := []string{}
	for p, _ := range i {
		result = append(result, p)
	}
	return result
}
