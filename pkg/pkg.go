package pkg

import "strings"

type Pkg string

// returns true if and only if q is a subpackage of p
func (p Pkg) HasSubpackage(q Pkg) bool {
	return strings.HasPrefix(string(q), string(p))
}

type PGraph map[Pkg]PSet

func NewPGraph() PGraph {
	return make(PGraph, 1)
}

func (g PGraph) AddImports(p Pkg, extra PSet) {
	if _, ok := g[p]; !ok {
		g[p] = NewPSet()
	}
	g[p].AddAll(extra)
}

func (g PGraph) Imports(p Pkg) (PSet, bool) {
	ps, ok := g[p]
	return ps, ok
}

func (g PGraph) Packages() PSet {
	ps := NewPSet()
	for p, _ := range g {
		ps.Add(p)
	}
	return ps
}

func (g PGraph) Walk(visit func(Pkg, PSet)) {
	for p, imp := range g {
		visit(p, imp)
	}
}

type PSet map[Pkg]bool

func NewPSet(ps ...Pkg) PSet {
	s := make(PSet, 1)
	for _, p := range ps {
		s.Add(p)
	}
	return s
}

func (s PSet) AddAll(t PSet) {
	for p, _ := range t {
		s.Add(p)
	}
}

func (s PSet) Add(p Pkg) {
	s[p] = true
}

func (s PSet) Contains(p Pkg) bool {
	_, ok := s[p]
	return ok
}

func (s PSet) Packages() []Pkg {
	pkgs := []Pkg{}
	for p, _ := range s {
		pkgs = append(pkgs, p)
	}
	return pkgs
}

func (s PSet) Size() int {
	return len(s)
}

func (s PSet) Walk(visit func(p Pkg)) {
	for p, _ := range s {
		visit(p)
	}
}
