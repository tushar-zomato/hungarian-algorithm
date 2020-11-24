package hungarianAlgorithm

import (
	"reflect"
	"testing"

	"github.com/shabbyrobe/go-num"
)

type testCase struct {
	in   [][]num.I128
	want []int
}

func TestSolve(t *testing.T) {
	cases := []testCase{
		{
			in: [][]num.I128{
				{num.I128From64(11), num.I128From64(6), num.I128From64(12)},
				{num.I128From64(12), num.I128From64(4), num.I128From64(6)},
				{num.I128From64(8), num.I128From64(12), num.I128From64(11)},
			},
			want: []int{1, 2, 0},
		},
		{
			in: [][]num.I128{
				{num.I128From64(13), num.I128From64(13), num.I128From64(19), num.I128From64(50), num.I128From64(33), num.I128From64(38)},
				{num.I128From64(73), num.I128From64(33), num.I128From64(71), num.I128From64(77), num.I128From64(97), num.I128From64(95)},
				{num.I128From64(20), num.I128From64(8), num.I128From64(56), num.I128From64(55), num.I128From64(64), num.I128From64(35)},
				{num.I128From64(26), num.I128From64(25), num.I128From64(72), num.I128From64(32), num.I128From64(55), num.I128From64(77)},
				{num.I128From64(83), num.I128From64(40), num.I128From64(69), num.I128From64(3), num.I128From64(53), num.I128From64(49)},
				{num.I128From64(67), num.I128From64(20), num.I128From64(44), num.I128From64(29), num.I128From64(86), num.I128From64(61)},
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
