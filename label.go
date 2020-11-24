package hungarianAlgorithm

import (
	"math"

	"lukechampine.com/uint128"
)

type label struct {
	n      int
	costs  [][]uint128.Uint128 //costs
	left   []uint128.Uint128   // labels on the rows
	right  []uint128.Uint128   // labels on the columns
	slack  []uint128.Uint128   // min slack
	slackI []int               // min slack index
}

func makeLabel(n int, costs [][]uint128.Uint128) label {
	left := make([]uint128.Uint128, n)
	right := make([]uint128.Uint128, n)
	slack := make([]uint128.Uint128, n)
	slackI := make([]int, n)
	return label{n, costs, left, right, slack, slackI}
}

// initialize left = min_j cost[i][j] for each row i
func (l *label) initialize() {
	for i := 0; i < l.n; i++ {
		l.left[i] = l.costs[i][0]
		for j := 1; j < l.n; j++ {
			if l.costs[i][j].Cmp(l.left[i]) < 0 {
				l.left[i] = l.costs[i][j]
			}
		}
	}
}

// Returns whether a given edge is tight
func (l *label) isTight(i int, j int) bool {
	return l.costs[i][j].Sub(l.left[i]).Sub(l.right[j]).IsZero()
}

// Given a set s of row indices and a set of column indices update the labels.
// Assumes that each indices set is sorted and contains no duplicate.
func (l *label) update(s []int, t []int) []edge {
	// find the minimum slack
	min := uint128.New(math.MaxUint64, math.MaxUint64)
	idx := 0
	for j := 0; j < l.n; j++ {
		if idx < len(t) && j == t[idx] {
			idx++
			continue
		}
		sl := l.slack[j]
		if sl.Cmp(min) < 0 {
			min = sl
		}
	}

	// increase the label on the elements of s
	for _, i := range s {
		l.left[i] = l.left[i].Add(min)
	}
	// decrease the label on the elements of t
	for _, i := range t {
		l.right[i] = l.right[i].Sub(min)
	}

	// decrease each slack by min and cache the tight edges
	edges := make([]edge, 0, l.n)
	idx = 0
	for j := 0; j < l.n; j++ {
		if idx < len(t) && j == t[idx] {
			idx++
			continue
		}
		l.slack[j] = l.slack[j].Sub(min)
		if l.slack[j].IsZero() {
			edges = append(edges, edge{l.slackI[j], j})
		}
	}

	return edges
}

func (l *label) initializeSlacks(i int) []edge {
	edges := make([]edge, 0, l.n)
	for j := 0; j < l.n; j++ {
		l.slack[j] = l.costs[i][j].Sub(l.left[i]).Sub(l.right[j])
		l.slackI[j] = i
		if l.slack[j].IsZero() {
			edges = append(edges, edge{i, j})
		}
	}
	return edges
}

func (l *label) updateSlacks(i int) []edge {
	edges := make([]edge, 0, l.n)
	for j := 0; j < l.n; j++ {
		s := l.costs[i][j].Sub(l.left[i]).Sub(l.right[j])
		if s.Cmp(l.slack[j]) < 0 {
			l.slack[j] = s
			l.slackI[j] = i
			if l.slack[j].IsZero() {
				edges = append(edges, edge{i, j})
			}
		}
	}
	return edges
}
