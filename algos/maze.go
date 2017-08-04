// Author hoenig

package algos

import (
	"bytes"
	"fmt"

	"github.com/fatih/color"
)

// Point is a cardinal point.
type Point struct {
	Row int
	Col int
}

// Is returns true if p represents the point at (row, col)
func (p Point) Is(row, col int) bool {
	return p.Row == row && p.Col == col
}

// Equal returns true if p and o represent the same point.
func (p Point) Equal(o Point) bool {
	return p.Row == o.Row && p.Col == o.Col
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.Col, p.Row)
}

// North returns a Point representing one cell above p.
func (p Point) North() Point {
	return Point{Row: p.Row - 1, Col: p.Col}
}

// South returns a Point representing one cell below p.
func (p Point) South() Point {
	return Point{Row: p.Row + 1, Col: p.Col}
}

// East returns a Point representing one cell left of p.
func (p Point) East() Point {
	return Point{Row: p.Row, Col: p.Col - 1}
}

// West returns a Point representing one cell right of p.
func (p Point) West() Point {
	return Point{Row: p.Row, Col: p.Col + 1}
}

// Maze represents a maze. For simplicity we assume
// the size of a maze it at least 3x3.
type Maze struct {
	cells [][]bool
	start Point
	end   Point
}

// NewMaze creates a maze given a set of cells and
// a start and end point.
func NewMaze(cells [][]bool, start, end Point) *Maze {
	return &Maze{
		cells: cells,
		start: start,
		end:   end,
	}
}

func (m *Maze) String() string {
	return m.Solution([]Point{})
}

// Solution creates a visualization of the solved maze.
func (m *Maze) Solution(points []Point) string {
	var buf bytes.Buffer
	for row := 0; row < len(m.cells); row++ {
		for col := 0; col < len(m.cells[0]); col++ {
			if m.start.Is(row, col) {
				buf.WriteString("s")
				continue
			}
			if m.end.Is(row, col) {
				buf.WriteString("e")
				continue
			}

			if contains(points, Point{Row: row, Col: col}) {
				buf.WriteString("x")
				continue
			}

			if m.cells[row][col] {
				buf.WriteString("o")
			} else {
				buf.WriteString(" ")
			}
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func contains(points []Point, point Point) bool {
	for _, p := range points {
		if p.Equal(point) {
			return true
		}
	}
	return false
}

// Wall returns true if the maze cell at p is a wall or is out of bounds.
func (m *Maze) Wall(p Point) bool {
	if !m.valid(p.Row, p.Col) {
		return true
	}

	return m.cells[p.Row][p.Col]
}

func (m *Maze) valid(row, col int) bool {
	return row >= 0 &&
		col >= 0 &&
		row < len(m.cells) &&
		col < len(m.cells[0])
}

// Start returns a copy of the maze start point.
func (m *Maze) Start() Point {
	return m.start
}

// End returns a copy of the maze end point.
func (m *Maze) End() Point {
	return m.end
}

// Colorize will print a maze solution to the terminal
// with lots of fancy colors.
func Colorize(maze string) {
	for _, c := range maze {
		switch c {
		case 'o':
			color.Set(color.FgBlue)
		case 's', 'e':
			color.Set(color.FgGreen)
		case 'x':
			color.Set(color.FgRed)
		}
		fmt.Print(string(c))
		color.Unset()
	}
}

func avoid(p Point, m *Maze, visited map[Point]bool) bool {
	if _, exists := visited[p]; exists {
		return true
	}
	if wall := m.Wall(p); wall {
		return true
	}
	return false
}
