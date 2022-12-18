package day14

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2022/pkg/grids"
)

const (
	sandStart = "+"
	sand      = "o"
	empty     = "."
	rock      = "#"
)

type Grid struct {
	g              *grids.Grid[string]
	originalBounds grids.Bounds

	sandStart grids.Loc
}

func newGrid(paths [][]grids.Loc, sandFrom grids.Loc) (*Grid, error) {
	g := grids.NewGrid(empty)

	for _, path := range paths {
		for i := 1; i < len(path); i++ {
			from, to := path[i-1], path[i]

			for at := from; ; at = stepTowards(at, to) {
				g.Set(at, rock)
				if at == to {
					break
				}
			}
		}
	}
	g.Set(sandFrom, sandStart)

	grid := &Grid{
		g:         g,
		sandStart: sandFrom,

		originalBounds: g.Bounds(),
	}

	return grid, nil
}

func (g *Grid) At(at grids.Loc) string {
	if g.IsAtFloor(at) {
		return rock
	}

	r, ok := g.g.At(at)
	if !ok {
		return empty
	}
	return r
}

func (g *Grid) IsAtFloor(at grids.Loc) bool {
	return at[1] == g.originalBounds.MaxY()+2
}

func (g *Grid) String() string {
	return g.g.String()
}

func (g *Grid) Count(r string) int {
	return g.g.Count(r)
}

func (g *Grid) SetRock(at grids.Loc) {
	g.g.Set(at, rock)
}

func (g *Grid) SetSand(at grids.Loc) {
	g.g.Set(at, sand)
}

func (g *Grid) InBounds(at grids.Loc) bool {
	return g.g.InBounds(at)
}

func parseInput(reader io.Reader) ([][]grids.Loc, error) {
	paths := make([][]grids.Loc, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		path := make([]grids.Loc, 0)
		pointsToParse := strings.Split(line, " -> ")
		for _, pt := range pointsToParse {
			xs, ys, ok := strings.Cut(pt, ",")
			if !ok {
				return nil, fmt.Errorf("malformed point pair: %q (line: %q)", pt, line)
			}

			x, err := strconv.Atoi(xs)
			if err != nil {
				return nil, fmt.Errorf("parsing %q: %w", xs, err)
			}

			y, err := strconv.Atoi(ys)
			if err != nil {
				return nil, fmt.Errorf("parsing %q: %w", ys, err)
			}

			path = append(path, grids.Loc{x, y})
		}

		paths = append(paths, path)
	}

	return paths, nil
}

func stepTowards(from, to grids.Loc) grids.Loc {
	x, y := from.XY()
	xDir, yDir := to[0]-from[0], to[1]-from[1]

	if xDir != 0 {
		if xDir < 0 {
			return grids.Loc{x - 1, y}
		} else {
			return grids.Loc{x + 1, y}
		}
	}

	if yDir < 0 {
		return grids.Loc{x, y - 1}
	}
	return grids.Loc{x, y + 1}
}
