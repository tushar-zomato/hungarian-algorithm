package hungarianAlgorithm

import (
	"reflect"
	"testing"

	"lukechampine.com/uint128"
)

type testCase struct {
	in   [][]uint128.Uint128
	want []int
}

func TestSolve(t *testing.T) {
	cases := []testCase{
		{
			in: [][]uint128.Uint128{
				{uint128.From64(11), uint128.From64(6), uint128.From64(12)},
				{uint128.From64(12), uint128.From64(4), uint128.From64(6)},
				{uint128.From64(8), uint128.From64(12), uint128.From64(11)},
			},
			want: []int{1, 2, 0},
		},
		{
			in: [][]uint128.Uint128{
				{uint128.From64(13), uint128.From64(13), uint128.From64(19), uint128.From64(50), uint128.From64(33), uint128.From64(38)},
				{uint128.From64(73), uint128.From64(33), uint128.From64(71), uint128.From64(77), uint128.From64(97), uint128.From64(95)},
				{uint128.From64(20), uint128.From64(8), uint128.From64(56), uint128.From64(55), uint128.From64(64), uint128.From64(35)},
				{uint128.From64(26), uint128.From64(25), uint128.From64(72), uint128.From64(32), uint128.From64(55), uint128.From64(77)},
				{uint128.From64(83), uint128.From64(40), uint128.From64(69), uint128.From64(3), uint128.From64(53), uint128.From64(49)},
				{uint128.From64(67), uint128.From64(20), uint128.From64(44), uint128.From64(29), uint128.From64(86), uint128.From64(61)},
			},
			want: []int{4, 1, 5, 0, 3, 2},
		},
	}

	for _, c := range cases {
		got, err := Solve(c.in)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Algo(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
