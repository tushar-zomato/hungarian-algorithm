package hungarianAlgorithm

import (
	"errors"

	"lukechampine.com/uint128"
)

func validate(costs [][]uint128.Uint128) error {
	n := len(costs)

	if n == 0 {
		return errors.New("The costs matrix is empty.")
	}

	if m := len(costs[0]); m != n {
		return errors.New("The costs matrix is not square.")
	}

	return nil
}
