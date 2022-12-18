package day14

import (
	"fmt"
	"io"

	"github.com/kristofferostlund/adventofcode-2022/pkg/location"
)

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

	isValid := func(loc location.Loc) bool {
		return grid.InBounds(loc)
	}

	for {
		if !simulateSand(grid, isValid) {
			break
		}
	}

	return grid.Count(sand), nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	paths, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	sandStartFrom := location.Loc{500, 0}
	grid, err := NewGrid(paths, sandStartFrom)
	if err != nil {
		return 0, fmt.Errorf("createing grid: %w", err)
	}

	filledStart := false
	isValid := func(loc location.Loc) bool {
		if filledStart {
			return false
		}
		if loc == sandStartFrom {
			filledStart = true
		}

		return true
	}

	for {
		if !simulateSand(grid, isValid) {
			break
		}
	}

	return grid.Count(sand), nil
}

func simulateSand(grid *Grid, isValid func(loc location.Loc) bool) bool {
	down := location.Loc{0, 1}
	left := location.Loc{-1, 0}
	right := location.Loc{1, 0}

	findNext := func(loc location.Loc) (location.Loc, bool) {
		next := loc.Add(down)
		switch grid.At(next) {
		case empty, sandStart:
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
		if !isValid(loc) {
			break
		}
	}

	if isValid(loc) {
		grid.SetSand(loc)
		return true
	}

	return false
}
