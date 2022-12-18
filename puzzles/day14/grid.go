package day14

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2022/pkg/location"
)

const (
	sandStart = '+'
	sand      = 'o'
	empty     = '.'
	rock      = '#'
)

type Grid struct {
	topLeft     location.Loc
	bottomRight location.Loc

	values   map[location.Loc]rune
	sandFrom location.Loc
}

func NewGrid(paths [][]location.Loc, sandFrom location.Loc) (*Grid, error) {
	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := 0, math.MinInt
	for _, path := range paths {
		for _, point := range path {
			x, y := point[0], point[1]

			if minX > x {
				minX = x
			}
			if maxX < x {
				maxX = x
			}

			if maxY < y {
				maxY = y
			}
		}
	}

	grid := &Grid{
		topLeft:     location.Loc{minX, minY},
		bottomRight: location.Loc{maxX, maxY},
		values:      make(map[location.Loc]rune),
		sandFrom:    sandFrom,
	}

	if err := grid.set(sandFrom, sandStart); err != nil {
		return nil, err
	}

	for _, path := range paths {
		for i := 1; i < len(path); i++ {
			from, to := path[i-1], path[i]

			for at := from; ; at = stepTowards(at, to) {
				if err := grid.SetRock(at); err != nil {
					return nil, err
				}

				if at == to {
					break
				}
			}
		}
	}

	return grid, nil
}

func (g *Grid) At(at location.Loc) rune {
	r, ok := g.values[at]
	if !ok {
		return empty
	}
	return r
}

func (g *Grid) String() string {
	minX, maxX := g.topLeft[0], g.bottomRight[0]
	minY, maxY := g.topLeft[1], g.bottomRight[1]

	sb := &strings.Builder{}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			at := location.Loc{x, y}
			sb.WriteRune(g.At(at))
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (g *Grid) SetRock(at location.Loc) error {
	return g.set(at, rock)
}

func (g *Grid) SetSand(at location.Loc) error {
	return g.set(at, sand)
}

func (g *Grid) set(at location.Loc, r rune) error {
	if !g.withinBounds(at) {
		return fmt.Errorf("%w: setting %s at %v", errOutOfBounds, string(r), at)
	}

	g.values[at] = r
	return nil
}

func (g *Grid) withinBounds(p location.Loc) bool {
	x, y := p[0], p[1]

	minX, maxX := g.topLeft[0], g.bottomRight[0]
	minY, maxY := g.topLeft[1], g.bottomRight[1]

	return minX <= x && x <= maxX &&
		minY <= y && y <= maxY
}

func parseInput(reader io.Reader) ([][]location.Loc, error) {
	paths := make([][]location.Loc, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		path := make([]location.Loc, 0)
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

			path = append(path, location.Loc{x, y})
		}

		paths = append(paths, path)
	}

	return paths, nil
}

func stepTowards(from, to location.Loc) location.Loc {
	x, y := from[0], from[1]
	xDir, yDir := to[0]-from[0], to[1]-from[1]

	if xDir != 0 {
		if xDir < 0 {
			return location.Loc{x - 1, y}
		} else {
			return location.Loc{x + 1, y}
		}
	}

	if yDir < 0 {
		return location.Loc{x, y - 1}
	}
	return location.Loc{x, y + 1}
}
