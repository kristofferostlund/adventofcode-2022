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
	minX, maxX int
	minY, maxY int

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
		minX: minX,
		maxX: maxX,
		minY: minY,
		maxY: maxY,

		values:   make(map[location.Loc]rune),
		sandFrom: sandFrom,
	}

	grid.set(sandFrom, sandStart)

	for _, path := range paths {
		for i := 1; i < len(path); i++ {
			from, to := path[i-1], path[i]

			for at := from; ; at = stepTowards(at, to) {
				grid.SetRock(at)
				if at == to {
					break
				}
			}
		}
	}

	return grid, nil
}

func (g *Grid) At(at location.Loc) rune {
	if at[1] == g.maxY+2 {
		return rock
	}

	r, ok := g.values[at]
	if !ok {
		return empty
	}
	return r
}

func (g *Grid) String() string {
	minX, maxX := g.minX, g.maxX
	minY, maxY := g.minY, g.maxY

	for loc := range g.values {
		if maxX < loc[0] {
			maxX = loc[0]
		}
		if loc[0] < minX {
			minX = loc[0]
		}

		if maxY < loc[1] {
			maxY = loc[1]
		}
		if loc[1] < minY {
			minY = loc[1]
		}
	}

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

func (g *Grid) Count(r rune) int {
	counter := 0
	for _, rr := range g.values {
		if rr == r {
			counter++
		}
	}
	return counter
}

func (g *Grid) SetRock(at location.Loc) {
	g.set(at, rock)
}

func (g *Grid) SetSand(at location.Loc) {
	g.set(at, sand)
}

func (g *Grid) set(at location.Loc, r rune) {
	g.values[at] = r
}

func (g *Grid) InBounds(p location.Loc) bool {
	x, y := p[0], p[1]

	minX, maxX := g.minX, g.maxX
	minY, maxY := g.minY, g.maxY

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
