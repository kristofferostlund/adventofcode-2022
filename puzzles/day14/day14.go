package day14

import (
	"fmt"
	"io"

	"github.com/kristofferostlund/adventofcode-2022/pkg/grids"
)

var sandStartFrom = grids.Loc{500, 0}

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	paths, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	grid, err := newGrid(paths, sandStartFrom)
	if err != nil {
		return 0, fmt.Errorf("creating grid: %w", err)
	}

	return p.solve(grid, grid.InBounds)
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	paths, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	grid, err := newGrid(paths, sandStartFrom)
	if err != nil {
		return 0, fmt.Errorf("creating grid: %w", err)
	}

	filledStart := false
	isValid := func(loc grids.Loc) bool {
		if filledStart {
			return false
		}
		if loc == sandStartFrom {
			filledStart = true
		}
		return !grid.IsAtFloor(loc)
	}

	return p.solve(grid, isValid)
}

func (p Puzzle) solve(grid *Grid, isValid func(loc grids.Loc) bool) (int, error) {
	for {
		if !simulateSand(grid, isValid) {
			break
		}
	}

	return grid.Count(sand), nil
}

func simulateSand(grid *Grid, isValid func(loc grids.Loc) bool) bool {
	down := grids.Loc{0, 1}
	left := grids.Loc{-1, 0}
	right := grids.Loc{1, 0}

	findNext := func(loc grids.Loc) (grids.Loc, bool) {
		next := loc.Add(down)
		switch grid.At(next) {
		case empty, sandStart:
			return next, true
		}

		toTry := []grids.Loc{next.Add(left), next.Add(right)}
		for _, next := range toTry {
			if grid.At(next) == empty {
				return next, true
			}
		}

		return loc, false
	}

	loc := grid.sandStart
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
