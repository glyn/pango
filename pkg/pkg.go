package pkg

type pkg string

type pgraph map[pkg]pset

func NewPGraph() pgraph {
	return make(pgraph, 1)
}

func (g pgraph) AddImports(p pkg, i ...pkg) {
	if imp, ok := g[p]; ok {
		g[p] = imp.AddAll(i)
	} else {
		g[p] = NewPSet(i)
	}
}

type pset map[pkg]bool

func NewPSet(ps ...pkg) pset {
	var s pset
	for _, p := range ps {
		s.Add(p)
	}
	return s
}

func (s pset) AddAll(t pset) {
	for p, _ := range t {
		s.Add(p)
	}
}

func (s pset) Add(p pkg) {
	s[p] = true
}
