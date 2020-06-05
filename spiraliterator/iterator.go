package spiraliterator

type cursor struct {
	col, row  int
	direction int
}

func (c *cursor) changeDirection() {
	c.direction = (c.direction + 1) % (up + 1)
}
func (c *cursor) move() {
	switch c.direction {
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
	c         cursor
	table     [][]int
	n         int
	stepsLeft int
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
		if s.n == 1 || (s.n == 2 && s.c.direction+1 == up) {
			s.n = 0
			return false
		}
		s.c.changeDirection()
		s.stepsLeft = s.n - 1
	} else if s.stepsLeft == 1 && s.c.direction == up {
		s.n -= 2
		if s.n == 0 {
			return false
		}
		s.c.changeDirection()
		s.stepsLeft = s.n
	}
	s.c.move()
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
		i.c.direction = up
	}
	return &i
}
