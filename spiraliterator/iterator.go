package spiraliterator

import "errors"

type cursor struct {
	col, row   int
	dcol, drow int
}

func (c *cursor) turnRight() {
	if c.drow != 0 {
		c.drow *= -1
	}
	c.dcol, c.drow = c.drow, c.dcol
}
func (c *cursor) move() {
	c.col += c.dcol
	c.row += c.drow
}

// spiralIterator traverses over a table in spiral-like manner
type spiralIterator struct {
	c      cursor
	table  [][]int
	n      int
	nSteps int
}

// Read returns an item the iterator points on
func (s *spiralIterator) Read() int {
	return s.table[s.c.row][s.c.col]
}

// Next positions the iterator on the next item
func (s *spiralIterator) Next() bool {
	if s.nSteps == s.n*s.n {
		return false
	}

	if s.nSteps > 1 && (s.nSteps-1)%(s.n-1) == 0 {
		s.c.turnRight()
	} else if s.nSteps/(s.n-1) == 4 {
		s.c.turnRight()
		s.n -= 2
		s.nSteps = 0
	}
	s.c.move()
	s.nSteps++
	return true
}

func validateTable(t [][]int) bool {
	rows := len(t)
	for _, col := range t {
		if len(col) != rows {
			return false
		}
	}
	return true
}

func Create(t [][]int) (spiralIterator, error) {
	if !validateTable(t) {
		return spiralIterator{table: [][]int{}}, errors.New("input table is expected to be n*n dimensional")
	}
	i := spiralIterator{table: t}
	i.Reset()
	return i, nil
}

func (i *spiralIterator) Reset() {
	if len(i.table) != 0 {
		i.c.col = -1
		i.n = len(i.table) + 2
		i.nSteps = (i.n - 1) * 4
		i.c.drow = -1
	}
}
