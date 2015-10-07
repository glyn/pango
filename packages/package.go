package packages

import "strings"

type Pkg string

func (p Pkg) HasSubpackage(q Pkg) bool {
	return strings.HasPrefix(string(q), string(p))
}
