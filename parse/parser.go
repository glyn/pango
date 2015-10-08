package parse

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	. "github.com/glyn/pango/packages"
)

type parseContext struct {
	scope  Pkg
	fset   *token.FileSet
	gopath string
}

func New(scope, gopath string) *parseContext {
	return &parseContext{
		scope:  Pkg(scope),
		fset:   token.NewFileSet(),
		gopath: gopath,
	}
}

func (pc *parseContext) Parse() (PGraph, error) {
	allImports := NewPGraph()
	scopeDir := pc.pkgDir(pc.scope)
	err := filepath.Walk(scopeDir, func(pathString string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return err
		}
		if strings.HasSuffix(pathString, ".git") {
			return filepath.SkipDir
		}
		pkg := pc.dirPkg(pathString)
		if _, ok := allImports.Imports(pkg); !ok {
			i := pc.parse(pkg)
			for p, imp := range i {
				if _, ok := allImports[p]; !ok {
					allImports[p] = imp
				}
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return allImports, nil
}

func (pc *parseContext) parse(root Pkg) PGraph {
	imports := NewPGraph()
	pc.Walk(root, func(p Pkg, imp PSet) {
		imports.AddImports(p, imp)
	})
	return imports
}

func (pc *parseContext) Walk(p Pkg, visit func(Pkg, PSet)) {
	pc.traverse(p, NewPSet(), visit)
}

func (pc *parseContext) traverse(p Pkg, visited PSet, visit func(Pkg, PSet)) {
	if visited.Contains(p) {
		return
	}
	visited.Add(p)

	imports := pc.imports(p)
	visit(p, imports)
	imports.Walk(func(imp Pkg) {
		if pc.scope.HasSubpackage(imp) {
			pc.traverse(imp, visited, visit)
		}
	})
}

func (pc *parseContext) ShortName(p Pkg) string {
	if !pc.scope.HasSubpackage(p) {
		return string(p)
	}
	return "." + string(p[len(pc.scope):])
}

func (pc *parseContext) ShortNames(p PSet) []string {
	result := []string{}
	p.Walk(func(q Pkg) {
		s := pc.ShortName(q)
		result = append(result, s)
	})
	return result
}

func (pc *parseContext) imports(p Pkg) PSet {
	i := NewPSet()
	pkgDir := pc.pkgDir(p)
	pAsts, err := parser.ParseDir(pc.fset, pkgDir, nil, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}
	for _, ast := range pAsts {
		for _, f := range ast.Files {
			for _, is := range f.Imports {
				q := Pkg(strings.Trim(is.Path.Value, `"`))
				if q != p && Pkg("github.com").HasSubpackage(q) {
					i.Add(q)
				}
			}
		}
	}
	return i
}

func (pc *parseContext) pkgDir(p Pkg) string {
	return filepath.Join(pc.gopath, "src", string(p))
}

func (pc *parseContext) dirPkg(dir string) Pkg {
	src := filepath.Join(pc.gopath, "src")
	if !strings.HasPrefix(dir, src) {
		panic(dir)
	}
	return Pkg(dir[len(src)+1:])
}
