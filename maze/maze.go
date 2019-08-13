package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	i int
	j int
}

func (p point) add(q point) point {
	return point{p.i + q.i, p.j + q.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}
func (p point) String() string {
	return fmt.Sprintf("(%d,%d)", p.i, p.j)
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func readmaze() [][]int {
	file, ok := os.Open("./maze.in")
	if ok != nil {
		panic(ok)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	var maze [][]int
	maze = make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze

}

func domake(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func dofind(maze [][]int, start point, end point) []point {
	Q := []point{end}
	path := []point{end}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		maxval, _ := cur.at(maze)

		if cur == start {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(maze)
			if !ok || val != (maxval-1) {
				continue
			}
			path = append(path, next)
			Q = append(Q, next)
		}

	}

	return path
}

func main() {
	maze := readmaze()

	steps := domake(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
	path := dofind(steps, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	var spath []string
	for i := len(path) - 1; i >= 0; i-- {
		spath = append(spath, fmt.Sprintf("(%d,%d)", path[i].i, path[i].j))
	}
	str := strings.Join(spath, "-->")
	fmt.Println(str)
}
