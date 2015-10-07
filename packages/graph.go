package packages

type PGraph map[Pkg]PSet

func NewPGraph() PGraph {
	return make(PGraph, 1)
}

func (g PGraph) AddImports(p Pkg, extra PSet) PGraph {
	if _, ok := g[p]; !ok {
		g[p] = NewPSet()
	}
	g[p].AddAll(extra)
	return g
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
