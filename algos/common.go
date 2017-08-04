// Author hoenig

package algos

// usd to
type storer interface {
	put([]Point)
	get() ([]Point, Point)
}

// evidence that BFS and DFS are the same thing
// apart from the undelrying datastructure used.
func solve(paths storer, m *Maze) []Point {
	visited := make(map[Point]bool)

	paths.put([]Point{m.Start()})

	for {
		trail, node := paths.get()

		if node.Equal(m.End()) {
			return trail
		}

		if up := node.North(); !avoid(up, m, visited) {
			paths.put(add(trail, up))
			visited[up] = true
		}
		if left := node.East(); !avoid(left, m, visited) {
			paths.put(add(trail, left))
			visited[left] = true
		}
		if down := node.South(); !avoid(down, m, visited) {
			paths.put(add(trail, down))
			visited[down] = true
		}
		if right := node.West(); !avoid(right, m, visited) {
			paths.put(add(trail, right))
			visited[right] = true
		}
	}
}

// ideally we would be less lazy and avoid using slices with
// their quirkey behavior when using append while holding on
// to an old reference
func add(points []Point, point Point) []Point {
	elms := make([]Point, len(points), len(points)+1)
	copy(elms, points)
	elms = append(elms, point)
	return elms
}
