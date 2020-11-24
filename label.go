package hungarianAlgorithm

import (
	"github.com/shabbyrobe/go-num"
)

type label struct {
	n      int
	costs  [][]num.I128 //costs
	left   []num.I128   // labels on the rows
	right  []num.I128   // labels on the columns
	slack  []num.I128   // min slack
	slackI []int        // min slack index
}

func makeLabel(n int, costs [][]num.I128) label {
	left := make([]num.I128, n)
	right := make([]num.I128, n)
	slack := make([]num.I128, n)
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
	min := num.MaxI128
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
