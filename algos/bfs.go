// Author hoenig

package algos

// BFS is a bredth-first-search maze solver.
type BFS struct {
	storer
	stack *stack
}

// NewBFS creates a new BFS solver.
func NewBFS() *BFS {
	return &BFS{stack: newStack()}
}

func (s *BFS) get() ([]Point, Point) {
	return s.stack.Pop()
}

func (s *BFS) put(points []Point) {
	s.stack.Push(points)
}

// Solve using bfs.
func (s *BFS) Solve(m *Maze) []Point {
	return solve(s, m)
}
