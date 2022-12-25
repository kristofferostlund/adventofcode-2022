package day17

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kristofferostlund/adventofcode-2022/pkg/grids"
	"github.com/kristofferostlund/adventofcode-2022/pkg/slices"
)

const (
	left  = "<"
	right = ">"
	down  = "v"
)

var rocks = [][]grids.Loc{
	// ####
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
	// .#.
	// ###
	// .#.
	{
		{1, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{1, 2},
	},
	// ..#
	// ..#
	// ###
	{
		{2, 0},
		{2, 1},
		{0, 2},
		{1, 2},
		{2, 2},
	},
	// #
	// #
	// #
	// #
	{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	},
	// ##
	// ##
	{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
	},
}

var rockBounds = func() []grids.Bounds {
	bounds := make([]grids.Bounds, 0, len(rocks))
	for _, rock := range rocks {
		bounds = append(bounds, grids.BoundsOf(rock))
	}
	return bounds
}()

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	directions, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	grid := grids.NewGrid(".")
	for i := 0; i < 7; i++ {
		grid.Set(grids.Loc{i, 0}, "-")
	}

	ri, curr, counter := nextRock(grid, -1, 0)

	for i, dir, useDir := nextDirection(directions, 0); counter <= 2022; i, dir, useDir = nextDirection(directions, i) {
		if useDir {
			next := move(dir, curr)
			b := grids.BoundsOf(next)
			gb := grid.Bounds()

			if gb.MinX() <= b.MinX() && b.MaxX() <= gb.MaxX() && !hasCollision(grid, next) {
				curr = next
			}
		} else {
			next := move(down, curr)
			if hasCollision(grid, next) {
				addRockToGrid(grid, curr, "#")

				ri, curr, counter = nextRock(grid, ri, counter)
			} else {
				curr = next
			}
		}
	}

	return grid.Bounds().Height(), nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	return 0, nil
}

func addRockToGrid(grid *grids.Grid[string], rock []grids.Loc, c string) {
	for _, l := range rock {
		grid.Set(l, c)
	}
}

func move(direction string, rock []grids.Loc) []grids.Loc {
	var dir grids.Loc
	switch direction {
	case down:
		dir = grids.Loc{0, 1} // down
	case right:
		dir = grids.Loc{1, 0} // right
	case left:
		dir = grids.Loc{-1, 0} // left
	default:
		panic(fmt.Sprintf("illegal direction %q", direction))
	}

	return addTo(rock, dir)
}

func hasCollision(g *grids.Grid[string], rock []grids.Loc) bool {
	for _, l := range rock {
		if _, collides := g.At(l); collides {
			return true
		}
	}
	return false
}

func nextRock(grid *grids.Grid[string], i, counter int) (int, []grids.Loc, int) {
	i = (i + 1) % len(rocks)
	next := slices.Copy(rocks[i])

	gridMinY := grid.Bounds().MinY()
	b := rockBounds[i]

	offset := grids.Loc{2, gridMinY - 4 - (b.MinY() + b.Height())}

	return i, addTo(next, offset), counter + 1
}

func nextDirection(directions []string, i int) (int, string, bool) {
	next := i + 1
	ok := i%2 == 0
	if !ok {
		return next, "", ok
	}

	di := (i / 2) % len(directions)
	return next, directions[di], ok
}

func addTo(rock []grids.Loc, offset grids.Loc) []grids.Loc {
	moved := slices.Copy(rock)
	for i, l := range moved {
		moved[i] = l.Add(offset)
	}
	return moved
}

func parseInput(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)

	directions := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		for _, r := range line {
			s := string(r)
			switch s {
			case left, right:
			default:
				return nil, fmt.Errorf("illegal direction %q", s)
			}

			directions = append(directions, s)
		}
	}
	return directions, nil
}
