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
	scope  string
	fset   *token.FileSet
	gopath string
}

func New(scope string) *disc {
	return &disc{
		scope:  scope,
		fset:   token.NewFileSet(),
		gopath: os.Getenv("GOPATH"),
	}
}

func (d *disc) Discover(root string) map[string][]string {
	imports := make(map[string][]string, 1)
	d.Walk(root, func(p string, imp []string) {
		imports[p] = imp
		fmt.Printf("Package %s imports: %v.\n", d.ShortName(p), d.ShortNames(imp))
	})
	return imports
}

func (d *disc) DiscoverAll() map[string][]string {
	allImports := make(map[string][]string, 1)
	scopeDir := d.pkgDir(d.scope)
	filepath.Walk(scopeDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return err
		}
		if strings.HasSuffix(path, ".git") {
			return filepath.SkipDir
		}
		pkg := d.dirPkg(path)
		if _, ok := allImports[pkg]; !ok {
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

func (d *disc) Walk(p string, visit func(string, []string)) {
	d.traverse(p, make(map[string]bool, 1), visit)
}

func (d *disc) traverse(p string, visited map[string]bool, visit func(string, []string)) {
	if visited[p] {
		return
	}
	visited[p] = true

	imports := d.imports(p)
	visit(p, imports)
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

func (d *disc) ShortName(p string) string {
	if !subpackage(p, d.scope) {
		panic(fmt.Sprintf("No shortname for %s in scope %s %v", p, d.scope, subpackage(d.scope, p)))
	}
	return p[len(d.scope):]
}

func (d *disc) ShortNames(p []string) []string {
	result := []string{}
	for _, q := range p {
		s := d.ShortName(q)
		result = append(result, s)
	}
	return result
}

func (d *disc) imports(pkg string) []string {
	i := make(map[string]bool, 1)
	pkgDir := d.pkgDir(pkg)
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

func (d *disc) pkgDir(p string) string {
	return filepath.Join(d.gopath, "src", p)
}

func (d *disc) dirPkg(dir string) string {
	src := filepath.Join(d.gopath, "src")
	if !strings.HasPrefix(dir, src) {
		panic(dir)
	}
	return dir[len(src)+1:]
}
