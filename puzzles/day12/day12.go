package day12

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kristofferostlund/adventofcode-2022/puzzles/day12/dijkstra"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	grid, start, dest, err := readInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing grid: %w", err)
	}

	steps := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	legalSteps := make(map[Loc][]Loc, 0)
	locs := grid.Locs()

	for _, loc := range locs {
		val, _ := grid.AtLoc(loc)

		for _, step := range steps {
			next := loc.Add(step)
			nextVal, ok := grid.AtLoc(next)
			if !ok {
				// Out of bounds
				continue
			}

			// This gives us the reverse of what we want, so we can't use this
			// even if this really is how the algo is suggested in the task.
			// is legal if either nextVal is one of {val+1, val, <val}
			if nextVal-val <= 1 {
				legalSteps[loc] = append(legalSteps[loc], next)
			}
		}
	}

	graph := dijkstra.NewGraph[Loc]()
	nodes := make(map[Loc]*dijkstra.Node[Loc])
	for _, l := range locs {
		node := dijkstra.NewNode(l)
		nodes[l] = node

		graph.AddNode(node)
	}

	for l, node := range nodes {
		for _, step := range legalSteps[l] {
			other, ok := nodes[step]
			if !ok {
				return 0, fmt.Errorf("no node found for %s", step)
			}
			graph.AddEdge(node, other, 1)
		}
	}

	dijkstra.Dijkstra(graph, start)

	cost, _, ok := graph.ShortestPath(dest)
	if !ok {
		return 0, fmt.Errorf("couldn't find a path from %s to %s", start, dest)
	}

	return cost, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	return 0, nil
}

type Loc [2]int

func (l Loc) Add(other Loc) Loc {
	return Loc{l[0] + other[0], l[1] + other[1]}
}

func (l Loc) String() string {
	return fmt.Sprintf("{x: %d, y: %d}", l[0], l[1])
}

type Grid [][]int

func (g Grid) AtLoc(loc Loc) (int, bool) {
	x, y := loc[0], loc[1]
	return g.At(x, y)
}

func (g Grid) At(x, y int) (int, bool) {
	if y < 0 || y >= len(g) {
		return 0, false
	}

	if x < 0 || x >= len(g[y]) {
		return 0, false
	}

	return g[y][x], true
}

func (g Grid) Locs() []Loc {
	locs := make([]Loc, len(g)*len(g[0]))
	for y, row := range g {
		for x := range row {
			loc := Loc{x, y}
			locs = append(locs, loc)
		}
	}
	return locs
}

func readInput(reader io.Reader) (Grid, Loc, Loc, error) {
	var start, dest Loc
	var grid [][]int

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		row := make([]int, 0, len(line))
		for i, r := range line {
			switch r {
			case 'S':
				start = Loc{i, len(grid)}
				row = append(row, int('a'))
			case 'E':
				dest = Loc{i, len(grid)}
				row = append(row, int('z'))
			default:
				row = append(row, int(r))
			}
		}

		grid = append(grid, row)
	}

	return Grid(grid), start, dest, nil
}
