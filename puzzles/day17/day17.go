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
)

var rocks = []Rock{
	{
		// ####
		{0, 0}, {1, 0}, {2, 0}, {3, 0},
	},
	{
		// .#.
		// ###
		// .#.
		{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2},
	},
	{
		// ..#
		// ..#
		// ###
		{2, 0}, {2, 1}, {0, 2}, {1, 2}, {2, 2},
	},
	{
		// #
		// #
		// #
		// #
		{0, 0}, {0, 1}, {0, 2}, {0, 3},
	},
	{
		// ##
		// ##
		{0, 0}, {1, 0}, {0, 1}, {1, 1},
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

	return simulateRocks(directions, 2022), nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	return 0, nil
}

func simulateRocks(directions []string, simulateCount int) int {
	grid := grids.NewGrid(".")
	for i := 0; i < 7; i++ {
		grid.Set(grids.Loc{i, 0}, "-")
	}

	rox := make([]Rock, 0)

	rock := nextRock(grid, len(rox))
	rox = append(rox, rock)

	for i := 0; len(rox) <= simulateCount; i++ {
		vec := nextDirection(directions, i)
		next := rock.Add(vec)

		isHorizontal := vec[0] != 0
		switch {
		case isHorizontal && next.WithinSides(grid) && !next.Collides(grid):
			rock = next
		case isHorizontal:
			// Do nothing if either !next.WithinSides(grid) or next.Collides(grid
		case !isHorizontal && next.Collides(grid):
			addToGrid(grid, rock)

			rock = nextRock(grid, len(rox))
			rox = append(rox, rock)
		case !isHorizontal:
			rock = next
		}
	}
	return grid.Bounds().Height()
}

func addToGrid(grid *grids.Grid[string], rock Rock) {
	for _, l := range rock {
		grid.Set(l, "#")
	}
}

func nextRock(grid *grids.Grid[string], rockCount int) Rock {
	i := rockCount % len(rocks)
	next := Rock(slices.Copy(rocks[i]))

	gridMinY := grid.Bounds().MinY()
	b := rockBounds[i]

	offset := grids.Loc{2, gridMinY - 4 - (b.MinY() + b.Height())}

	return next.Add(offset)
}

func nextDirection(directions []string, i int) grids.Loc {
	isHorizontal := i%2 == 0
	if !isHorizontal {
		return grids.Loc{0, 1}
	}

	switch dir := directions[(i/2)%len(directions)]; dir {
	case right:
		return grids.Loc{1, 0}
	case left:
		return grids.Loc{-1, 0}
	default:
		panic(fmt.Sprintf("illegal direction %q", dir))
	}
}

type Rock []grids.Loc

func (r Rock) Add(offset grids.Loc) Rock {
	moved := slices.Copy(r)
	for i, l := range moved {
		moved[i] = l.Add(offset)
	}
	return moved
}

func (r Rock) Collides(g *grids.Grid[string]) bool {
	for _, l := range r {
		if _, collides := g.At(l); collides {
			return true
		}
	}
	return false
}

func (r Rock) WithinSides(g *grids.Grid[string]) bool {
	b := grids.BoundsOf(r)
	gb := g.Bounds()

	return gb.MinX() <= b.MinX() && b.MaxX() <= gb.MaxX()
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
