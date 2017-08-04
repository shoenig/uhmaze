// Author hoenig

package maze

import (
	"bufio"
	"io"
	"os"

	"github.com/shoenig/uhmaze/algos"
)

// Loader loads a maze file.
type Loader interface {
	Load() (algos.Maze, error)
}

// ASCIILoader loads ascii mazes in the samples directory.
type ASCIILoader struct {
	Filename string
}

// Load the maze
func (l *ASCIILoader) Load() (*algos.Maze, error) {
	file, err := os.Open(l.Filename)
	if err != nil {
		return nil, err
	}

	return l.parse(bufio.NewReader(file))
}

func (l *ASCIILoader) parse(reader io.Reader) (*algos.Maze, error) {
	scanner := bufio.NewScanner(reader)
	var cells [][]bool
	var start algos.Point
	var end algos.Point
	var i int
	for scanner.Scan() {
		line := scanner.Text()
		row, entrance, exit := l.line(line)
		cells = append(cells, row)
		if entrance {
			start = algos.Point{Row: i, Col: 0}
		}
		if exit {
			end = algos.Point{Row: i, Col: len(line) - 1}
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return algos.NewMaze(cells, start, end), nil
}

func (l *ASCIILoader) line(s string) ([]bool, bool, bool) {
	var cells []bool
	for _, c := range s {
		if c == ' ' {
			cells = append(cells, false)
		} else {
			cells = append(cells, true)
		}
	}
	return cells, !cells[0], !cells[len(cells)-1]
}
