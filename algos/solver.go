// Author hoenig

package algos

// Solver represents one solving algorithm.
type Solver interface {
	// Solve will return a list of points leading
	// from the maze start to the maze end.
	Solve(*Maze) []Point
}
