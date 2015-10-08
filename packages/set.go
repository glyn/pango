package packages

import (
	"bytes"
	"fmt"
)

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

func (s PSet) Size() int {
	return len(s)
}

func (s PSet) Walk(visit func(p Pkg)) {
	for p, _ := range s {
		visit(p)
	}
}

func (s PSet) Equals(t PSet) bool {
	return s.ContainsAll(t) && t.ContainsAll(s)
}

func (s PSet) ContainsAll(t PSet) bool {
	ca := true
	t.Walk(func(p Pkg) {
		if !s.Contains(p) {
			ca = false
		}
	})
	return ca
}

func (s PSet) String() string {
	buff := bytes.NewBufferString("{")
	first := true
	s.Walk(func(p Pkg) {
		if !first {
			buff.WriteString(",\n")
		}
		first = false
		buff.WriteString(fmt.Sprintf("%q", string(p)))
	})
	buff.WriteString("}")
	return buff.String()
}
