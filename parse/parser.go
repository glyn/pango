package parse

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	. "github.com/glyn/pango/packages"
)

type disc struct {
	scope  Pkg
	fset   *token.FileSet
	gopath string
}

func New(scope string) *disc {
	return &disc{
		scope:  Pkg(scope),
		fset:   token.NewFileSet(),
		gopath: os.Getenv("GOPATH"),
	}
}

func (d *disc) Discover(root Pkg) PGraph {
	imports := NewPGraph()
	d.Walk(root, func(p Pkg, imp PSet) {
		imports.AddImports(p, imp)
		fmt.Printf("Package %s imports: %v.\n", d.ShortName(p), d.ShortNames(imp))
	})
	return imports
}

func (d *disc) DiscoverAll() PGraph {
	allImports := make(PGraph, 1)
	scopeDir := d.pkgDir(d.scope)
	filepath.Walk(scopeDir, func(pathString string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return err
		}
		if strings.HasSuffix(pathString, ".git") {
			return filepath.SkipDir
		}
		pkg := d.dirPkg(pathString)
		if _, ok := allImports.Imports(pkg); !ok {
			i := d.Discover(pkg)
			for p, imp := range i {
				if _, ok := allImports[p]; !ok {
					allImports[p] = imp
				}
			}
		}
		return err
	})
	return allImports
}

func (d *disc) Walk(p Pkg, visit func(Pkg, PSet)) {
	d.traverse(p, NewPSet(), visit)
}

func (d *disc) traverse(p Pkg, visited PSet, visit func(Pkg, PSet)) {
	if visited.Contains(p) {
		return
	}
	visited.Add(p)

	imports := d.imports(p)
	visit(p, imports)
	imports.Walk(func(imp Pkg) {
		if d.scope.HasSubpackage(imp) {
			d.traverse(imp, visited, visit)
		}
	})
}

func (d *disc) ShortName(p Pkg) string {
	if !d.scope.HasSubpackage(p) {
		panic(fmt.Sprintf("No shortname for %s in scope %s %v", p, d.scope, d.scope.HasSubpackage(p)))
	}
	return string(p[len(d.scope):])
}

func (d *disc) ShortNames(p PSet) []string {
	result := []string{}
	p.Walk(func(q Pkg) {
		s := d.ShortName(q)
		result = append(result, s)
	})
	return result
}

func (d *disc) imports(p Pkg) PSet {
	i := NewPSet()
	pkgDir := d.pkgDir(p)
	pAsts, err := parser.ParseDir(d.fset, pkgDir, nil, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}
	for _, ast := range pAsts {
		for _, f := range ast.Files {
			for _, is := range f.Imports {
				q := Pkg(strings.Trim(is.Path.Value, `"`))
				if q != p && d.scope.HasSubpackage(q) {
					i.Add(q)
				}
			}
		}
	}
	return i
}

func (d *disc) pkgDir(p Pkg) string {
	return filepath.Join(d.gopath, "src", string(p))
}

func (d *disc) dirPkg(dir string) Pkg {
	src := filepath.Join(d.gopath, "src")
	if !strings.HasPrefix(dir, src) {
		panic(dir)
	}
	return Pkg(dir[len(src)+1:])
}
