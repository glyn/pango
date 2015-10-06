package disc

import (
	"fmt"
	"go/parser"
	"path/filepath"

	"github.com/glyn/pango/pkg"

	"os"

	"go/token"

	"strings"
)

type disc struct {
	scope  pkg.Pkg
	fset   *token.FileSet
	gopath string
}

func New(scope string) *disc {
	return &disc{
		scope:  pkg.Pkg(scope),
		fset:   token.NewFileSet(),
		gopath: os.Getenv("GOPATH"),
	}
}

func (d *disc) Discover(root pkg.Pkg) pkg.PGraph {
	imports := pkg.NewPGraph()
	d.Walk(root, func(p pkg.Pkg, imp pkg.PSet) {
		imports.AddImports(p, imp)
		fmt.Printf("Package %s imports: %v.\n", d.ShortName(p), d.ShortNames(imp))
	})
	return imports
}

func (d *disc) DiscoverAll() pkg.PGraph {
	allImports := make(pkg.PGraph, 1)
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

func (d *disc) Walk(p pkg.Pkg, visit func(pkg.Pkg, pkg.PSet)) {
	d.traverse(p, pkg.NewPSet(), visit)
}

func (d *disc) traverse(p pkg.Pkg, visited pkg.PSet, visit func(pkg.Pkg, pkg.PSet)) {
	if visited.Contains(p) {
		return
	}
	visited.Add(p)

	imports := d.imports(p)
	visit(p, imports)
	imports.Walk(func(imp pkg.Pkg) {
		if d.scope.HasSubpackage(imp) {
			d.traverse(imp, visited, visit)
		}
	})
}

func (d *disc) ShortName(p pkg.Pkg) string {
	if !d.scope.HasSubpackage(p) {
		panic(fmt.Sprintf("No shortname for %s in scope %s %v", p, d.scope, d.scope.HasSubpackage(p)))
	}
	return string(p[len(d.scope):])
}

func (d *disc) ShortNames(p pkg.PSet) []string {
	result := []string{}
	p.Walk(func(q pkg.Pkg) {
		s := d.ShortName(q)
		result = append(result, s)
	})
	return result
}

func (d *disc) imports(p pkg.Pkg) pkg.PSet {
	i := pkg.NewPSet()
	pkgDir := d.pkgDir(p)
	pAsts, err := parser.ParseDir(d.fset, pkgDir, nil, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}
	for _, ast := range pAsts {
		for _, f := range ast.Files {
			for _, is := range f.Imports {
				q := pkg.Pkg(strings.Trim(is.Path.Value, `"`))
				if q != p && d.scope.HasSubpackage(q) {
					i.Add(q)
				}
			}
		}
	}
	return i
}

func (d *disc) pkgDir(p pkg.Pkg) string {
	return filepath.Join(d.gopath, "src", string(p))
}

func (d *disc) dirPkg(dir string) pkg.Pkg {
	src := filepath.Join(d.gopath, "src")
	if !strings.HasPrefix(dir, src) {
		panic(dir)
	}
	return pkg.Pkg(dir[len(src)+1:])
}
