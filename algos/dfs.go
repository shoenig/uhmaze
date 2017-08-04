// Author hoenig

package algos

// DFS is a depth-first-search maze solver.
type DFS struct {
	storer
	queue *queue
}

// NewDFS creates a new DFS solver.
func NewDFS() *DFS {
	return &DFS{queue: newQueue()}
}

func (s *DFS) get() ([]Point, Point) {
	return s.queue.Poll()
}

func (s *DFS) put(points []Point) {
	s.queue.Put(points)
}

// Solve using dfs.
func (s *DFS) Solve(m *Maze) []Point {
	return solve(s, m)
}
