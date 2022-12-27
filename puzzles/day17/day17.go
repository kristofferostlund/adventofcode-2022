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
	directions, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	return simulateRocks(directions, 1000000000000), nil
}

func simulateRocks(directions []string, simulateCount int) int {
	grid := grids.NewGrid(".")
	for i := 0; i < 7; i++ {
		grid.Set(grids.Loc{i, 0}, "-")
	}

	type checkpoint struct {
		rockCount int
		height    int
	}

	isFirstRepeatPattern := func() func(grid *grids.Grid[string], ri, di, rockCount int) ([2]checkpoint, bool) {
		seen := make(map[string]checkpoint)
		repeatAlreadyFound := false

		return func(grid *grids.Grid[string], ri, di, rockCount int) ([2]checkpoint, bool) {
			if repeatAlreadyFound {
				return [2]checkpoint{}, false
			}

			pt := checkpoint{
				rockCount: rockCount,
				height:    grid.Bounds().Height(),
			}

			sg := subsetGridOf(grid)
			key := fmt.Sprintf("%d-%d-%s", ri, di, sg)

			prev, isSeen := seen[key]
			if !isSeen {
				seen[key] = pt
				return [2]checkpoint{}, false
			}

			repeatAlreadyFound = true
			return [2]checkpoint{prev, pt}, true
		}
	}()

	rockCount := 0
	rock, ri := nextRock(grid, rockCount)
	simulatedHeight := 0

	for i := 0; rockCount < simulateCount; i++ {
		vec, di := nextDirection(directions, i)
		next := rock.Add(vec)

		isHorizontal := vec[0] != 0
		switch {
		case isHorizontal && next.WithinSides(grid) && !next.Collides(grid):
			rock = next
		case isHorizontal:
			// Do nothing if either !next.WithinSides(grid) or next.Collides(grid
		case !isHorizontal && next.Collides(grid):
			addToGrid(grid, rock)
			rockCount++

			// The first time we see a repeating pattern we can calculate
			// the height difference from the first occurrence and the repeat
			// to then calculate how many extra iterations would be needed
			// to eventually get to the final count. Since it's highly unlikely
			// all remaining iterations would be "exhausted" like this, we
			// increase the rockCount and store the simulated height.
			if checkpoints, ok := isFirstRepeatPattern(grid, ri, di, rockCount); ok {
				a, b := checkpoints[0], checkpoints[1]
				rDelta := b.rockCount - a.rockCount
				hDelta := b.height - a.height

				remainder := simulateCount - rockCount

				mul := remainder / rDelta
				rockCount += mul * rDelta
				simulatedHeight = mul * hDelta
			}

			rock, ri = nextRock(grid, rockCount)
		case !isHorizontal:
			rock = next
		}
	}

	return grid.Bounds().Height() + simulatedHeight
}

func addToGrid(grid *grids.Grid[string], rock Rock) {
	for _, l := range rock {
		grid.Set(l, "#")
	}
}

func nextRock(grid *grids.Grid[string], rockCount int) (Rock, int) {
	i := rockCount % len(rocks)
	gridMinY := grid.Bounds().MinY()
	b := rockBounds[i]

	offset := grids.Loc{2, gridMinY - 4 - (b.MinY() + b.Height())}

	return rocks[i].Add(offset), i
}

func nextDirection(directions []string, i int) (grids.Loc, int) {
	di := (i / 2) % len(directions)

	isHorizontal := i%2 == 0
	if !isHorizontal {
		return grids.Loc{0, 1}, di
	}

	switch dir := directions[di]; dir {
	case right:
		return grids.Loc{1, 0}, di
	case left:
		return grids.Loc{-1, 0}, di
	default:
		panic(fmt.Sprintf("illegal direction %q", dir))
	}
}

func subsetGridOf(grid *grids.Grid[string]) *grids.Grid[string] {
	var got, want uint64 = 0, 0

	bounds := grid.Bounds()
	xb := make([]uint64, bounds.Width()+1)
	var b uint64 = 1
	for i := 0; i < len(xb); i++ {
		b <<= 1
		xb[i] = b
		want |= b
	}

	subGrid := grids.NewGrid(".")

	for y := bounds.MinY(); y <= bounds.MaxY(); y++ {
		for x := bounds.MinX(); x <= bounds.MaxX(); x++ {
			at := grids.Loc{x, y}
			value, ok := grid.At(at)
			if ok {
				got |= xb[x]
				subGrid.Set(at, value)

				if got == want {
					return subGrid
				}
			}
		}
	}

	// for i := len(fallenRocks) - 1; i >= 0; i-- {
	// 	rock := fallenRocks[i]
	// 	for _, l := range rock {
	// 		subGrid.Set(l, "#")
	// 		x, _ := l.XY()
	// 		got |= xb[x]
	// 	}

	// 	if got == want {
	// 		return subGrid
	// 	}
	// }

	return subGrid
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

func (r Rock) Copy() Rock {
	return slices.Copy(r)
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
