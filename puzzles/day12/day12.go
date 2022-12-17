package day12

import (
	"bufio"
	"fmt"
	"io"
	"math"

	"github.com/kristofferostlund/adventofcode-2022/puzzles/day12/dijkstra"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	grid, start, dest, err := readInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing grid: %w", err)
	}

	graph, err := p.setupGraph(grid)
	if err != nil {
		return 0, fmt.Errorf("setting up graph: %w", err)
	}

	cost, _, ok := graph.ShortestPath(start, dest)
	if !ok {
		return 0, fmt.Errorf("couldn't find a path from %s to %s", start, dest)
	}

	return cost, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	grid, _, dest, err := readInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing grid: %w", err)
	}

	graph, err := p.setupGraph(grid)
	if err != nil {
		return 0, fmt.Errorf("setting up graph: %w", err)
	}

	startLocs := make([]Loc, 0)
	for _, loc := range grid.Locs() {
		val, _ := grid.AtLoc(loc)
		if val == int('a') {
			startLocs = append(startLocs, loc)
		}
	}

	smallest := math.MaxInt64
	for _, start := range startLocs {
		cost, _, ok := graph.ShortestPath(start, dest)
		if !ok {
			return 0, fmt.Errorf("couldn't find a path from %s to %s", start, dest)
		}
		if cost < smallest {
			smallest = cost
		}
	}

	return smallest, nil
}

func (Puzzle) setupGraph(grid Grid) (*dijkstra.Graph[Loc], error) {
	graph := dijkstra.NewGraph[Loc]()
	for _, l := range grid.Locs() {
		graph.AddNode(dijkstra.NewNode(l))
	}

	steps := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for _, node := range graph.Nodes {
		loc := node.Value()
		val, _ := grid.AtLoc(loc)

		for _, step := range steps {
			next := loc.Add(step)
			nextVal, ok := grid.AtLoc(next)
			if !ok {
				// Out of bounds
				continue
			}

			// "at most one higher than the elevation of your current square"
			if nextVal-val <= 1 {
				other := graph.GetNode(next)
				if other == nil {
					return nil, fmt.Errorf("no node found for %s", next)
				}
				graph.AddEdge(node, other, 1)
			}
		}
	}

	return graph, nil
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
