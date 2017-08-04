// Author hoenig

package algos

type stack struct {
	elements [][]Point // tail is top of stack
}

func newStack() *stack {
	return &stack{
		elements: make([][]Point, 0, 100),
	}
}

func (s *stack) Push(pts []Point) {
	s.elements = append(s.elements, pts)
}

func (s *stack) Pop() ([]Point, Point) {
	p := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return p, p[len(p)-1]
}
