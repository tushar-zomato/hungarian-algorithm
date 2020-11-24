package hungarianAlgorithm

import (
	"errors"

	"github.com/shabbyrobe/go-num"
)

func validate(costs [][]num.I128) error {
	n := len(costs)

	if n == 0 {
		return errors.New("The costs matrix is empty.")
	}

	if m := len(costs[0]); m != n {
		return errors.New("The costs matrix is not square.")
	}

	return nil
}
