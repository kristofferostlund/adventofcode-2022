package day15

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"

	"github.com/kristofferostlund/adventofcode-2022/pkg/grids"
	"github.com/kristofferostlund/adventofcode-2022/pkg/ints"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader, cy int, debug bool) (int, error) {
	sensors, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	var grid *grids.Grid[string]
	if debug {
		grid = prepareGrid(sensors)
		addNoBeaconZones(grid, sensors)
	}

	xRanges := make([][2]int, 0)
	for i, s := range sensors {
		sx, sy := s.At.XY()
		manhattan := s.ManhattanDistance()

		straightLine := ints.Abs(sy - cy)
		if manhattan < straightLine {
			// If a straight line is longer than the manhattan distance,
			// we can't have any overlap!
			continue
		}

		minX := sx + (straightLine - manhattan)
		maxX := sx - (straightLine - manhattan)

		if debug {
			fmt.Println(s.At)
			fmt.Println("  straight line", straightLine)
			fmt.Println("  minX", minX)
			fmt.Println("  maxX", maxX)

			// Print the current sensor's no-beacon zone
			g := prepareGrid(sensors)
			addNoBeaconZones(g, sensors[i:i+1])
			fmt.Printf("  %s\n", g.RenderRow(cy))
		}

		xRanges = append(xRanges, [2]int{minX, maxX})
	}

	counter := 0
	sort.Slice(xRanges, func(i, j int) bool {
		if xRanges[i][0] == xRanges[j][0] {
			return xRanges[i][1] < xRanges[j][1]
		}
		return xRanges[i][0] < xRanges[j][0]
	})

	toUse := make([][2]int, 0)
	for _, xRange := range xRanges {
		overlapIdx := -1
		for i, prev := range toUse {
			if overlaps(prev, xRange) {
				overlapIdx = i
				break
			}
		}

		if overlapIdx > -1 {
			prev := toUse[overlapIdx]
			toUse[overlapIdx] = [2]int{ints.Min(prev[0], xRange[0]), ints.Max(prev[1], xRange[1])}
		} else {
			toUse = append(toUse, xRange)
		}
	}

	for _, xRange := range toUse {
		counter += xRange[1] - xRange[0]
	}

	if debug {
		fmt.Printf("\n%s\n", grid.RenderRow(cy))
		fmt.Printf("\n%s\n", grid)
	}

	return counter, nil
}

func (p Puzzle) Part2(reader io.Reader, bounds grids.Bounds) (int, error) {
	sensors, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	for i, s := range sensors {
		// Since there's exactly one point on the map,
		// it must be exactly 1 space oustide the range.
		distance := sensors[i].ManhattanDistance() + 1

		// Just learned about the [...]type shorthand for creating arrays
		extremes := [...]grids.Loc{
			s.At.Add(grids.Loc{0, distance}),  // top
			s.At.Add(grids.Loc{distance, 0}),  // right
			s.At.Add(grids.Loc{0, -distance}), // bottom
			s.At.Add(grids.Loc{-distance, 0}), // left
		}

		for ei, extreme := range extremes {
			next := extremes[(ei+1)%len(extremes)]

			for loc := extreme; loc != next; loc = stepTowards(loc, next) {
				if !bounds.IsInside(loc) {
					continue
				}

				outsideAll := true
				for j, other := range sensors {
					if j == i {
						continue
					}

					if other.InRangeOf(loc) {
						outsideAll = false
						break
					}
				}

				if outsideAll {
					x, y := loc.XY()
					return x*4000000 + y, nil
				}
			}
		}
	}

	return 0, errors.New("expected exactly one available space within the area, found none")
}

type Sensor struct {
	At     grids.Loc
	Beacon grids.Loc
}

func (s Sensor) ManhattanDistance() int {
	return manhattanDistance(s.At, s.Beacon)
}

func (s Sensor) InRangeOf(loc grids.Loc) bool {
	return manhattanDistance(s.At, loc) <= s.ManhattanDistance()
}

func manhattanDistance(a, b grids.Loc) int {
	x, y := xyDiff(a, b)
	return x + y
}

func xyDiff(a, b grids.Loc) (int, int) {
	x, y := ints.Abs(a[0]-b[0]), ints.Abs(a[1]-b[1])
	return x, y
}

func parseInput(reader io.Reader) ([]Sensor, error) {
	sensors := make([]Sensor, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var sensorAt, beaconAt [2]int
		// I learned about fmt.Sscanf which is extremely convenient for AoC since
		// a lot of the input is basically known strings like below with the only
		// _variable_ data is what's to parse out!
		scannedCount, err := fmt.Sscanf(
			line,
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensorAt[0],
			&sensorAt[1],
			&beaconAt[0],
			&beaconAt[1],
		)
		if err != nil {
			return nil, fmt.Errorf("scanning line %q: %w", line, err)
		}
		if got, want := scannedCount, 4; got != want {
			return nil, fmt.Errorf("malformed line %q, got %d scanned values, want %d", line, got, want)
		}

		sensors = append(sensors, Sensor{At: sensorAt, Beacon: beaconAt})
	}
	return sensors, nil
}

func overlaps(a [2]int, b [2]int) bool {
	aLower, aUpper := a[0], a[1]
	bLower, bUpper := b[0], b[1]

	x := aLower <= bLower && bLower <= aUpper ||
		aLower <= bUpper && bUpper <= aUpper ||
		bLower <= aLower && aLower <= bUpper ||
		bLower <= aUpper && aUpper <= bUpper
	return x
}
