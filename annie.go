package annie

import (
	"annie/internal"
	"annie/pkg"
)

func NewAnnie(path string) (pkg.Annie, error) {
	return internal.NewAnnie(path)
}
