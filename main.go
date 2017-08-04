// Author hoenig

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/shoenig/uhmaze/algos"
	"github.com/shoenig/uhmaze/maze"
)

func main() {
	algo := os.Args[1]

	filename := os.Args[2]
	loader := maze.ASCIILoader{Filename: filename}
	maze, err := loader.Load()

	if err != nil {
		fatal(err)
	}

	switch algo {
	case "bfs":
		run("BFS", algos.NewBFS(), maze)
	case "dfs":
		run("DFS", algos.NewDFS(), maze)
	default:
		fatal(errors.Errorf("i dunno what %q is", algo))
	}
}

func run(name string, s algos.Solver, maze *algos.Maze) {
	start := time.Now()
	solution := s.Solve(maze)
	end := time.Now()
	algos.Colorize(maze.Solution(solution))
	fmt.Printf("%s solved in %v\n", name, end.Sub(start))
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
