package packages

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
