package annie

import (
	"annie/internal"
	"annie/pkg"
)

func NewAnnie(path string) (pkg.Parent, error) {
	return internal.NewAnnie(path)
}
