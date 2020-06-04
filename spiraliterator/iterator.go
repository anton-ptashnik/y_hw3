package spiraliterator

type cursor struct {
	col, row int
}

func (c *cursor) move(d int) {
	switch d {
	case right:
		c.col++
	case down:
		c.row++
	case left:
		c.col--
	case up:
		c.row--
	}
}

// spiralIterator traverses over a table in spiral-like manner
type spiralIterator struct {
	c             cursor
	table         [][]int
	n             int
	stepsLeft     int
	moveDirection int
}

// Read returns an item the iterator points on
func (s *spiralIterator) Read() int {
	return s.table[s.c.row][s.c.col]
}

// Next positions the iterator on the next item
func (s *spiralIterator) Next() bool {
	if s.n == 0 {
		return false
	}
	if s.stepsLeft == 0 {
		if s.n == 1 {
			s.n = 0
			return false
		}
		s.moveDirection++
		s.stepsLeft = s.n - 1
	} else if s.moveDirection == up && s.stepsLeft == 1 {
		s.moveDirection = right
		s.n -= 2
		s.stepsLeft = s.n
		if s.n == 0 {
			return false
		}
	}
	s.c.move(s.moveDirection)
	s.stepsLeft--
	return true
}

const (
	right = iota
	down
	left
	up
)

func New(t [][]int) *spiralIterator {
	i := spiralIterator{}
	if len(t) != 0 {
		i.table = t
		i.c.col = -1
		i.n = len(t) + 2
		i.stepsLeft = 1
		i.moveDirection = up
	}
	return &i
}
