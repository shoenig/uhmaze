// Author hoenig

package algos

type queue struct {
	elements [][]Point
}

func newQueue() *queue {
	return &queue{
		elements: make([][]Point, 0, 100),
	}
}

func (q *queue) Put(points []Point) {
	q.elements = append(q.elements, points)
}

func (q *queue) Poll() ([]Point, Point) {
	points := q.elements[0]
	q.elements = q.elements[1:]
	return points, points[len(points)-1]
}
