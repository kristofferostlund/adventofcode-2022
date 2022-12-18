package day14

import (
	"errors"
	"fmt"
	"io"

	"github.com/kristofferostlund/adventofcode-2022/pkg/location"
)

var errOutOfBounds = errors.New("out of bounds")

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	paths, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	sandStartFrom := location.Loc{500, 0}
	grid, err := NewGrid(paths, sandStartFrom)
	if err != nil {
		return 0, fmt.Errorf("createing grid: %w", err)
	}

	counter := 0
	for err := simulateSand(grid); err == nil; err = simulateSand(grid) {
		counter++
	}

	return counter, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	return 0, nil
}

func simulateSand(grid *Grid) error {
	down := location.Loc{0, 1}
	left := location.Loc{-1, 0}
	right := location.Loc{1, 0}

	findNext := func(loc location.Loc) (location.Loc, bool) {
		next := loc.Add(down)
		if grid.At(next) == empty {
			return next, true
		}

		toTry := []location.Loc{next.Add(left), next.Add(right)}
		for _, next := range toTry {
			if grid.At(next) == empty {
				return next, true
			}
		}

		return loc, false
	}

	loc := grid.sandFrom
	for next, valid := findNext(loc); valid; next, valid = findNext(loc) {
		loc = next
		if !grid.withinBounds(loc) {
			break
		}
	}

	return grid.SetSand(loc)
}
