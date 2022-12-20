package day15

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

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
			if prev[0] <= xRange[0] && xRange[0] <= prev[1] ||
				prev[0] <= xRange[1] && xRange[1] <= prev[1] ||
				xRange[0] <= prev[0] && prev[0] <= xRange[1] ||
				xRange[0] <= prev[1] && prev[1] <= xRange[1] {
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

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	return 0, nil
}

type Sensor struct {
	At     grids.Loc
	Beacon grids.Loc
}

func (s Sensor) ManhattanDistance() int {
	return manhattanDistance(s.At, s.Beacon)
}

func manhattanDistance(a, b grids.Loc) int {
	return ints.Abs(a[0]-b[0]) + ints.Abs(a[1]-b[1])
}

func parseInput(reader io.Reader) ([]Sensor, error) {
	sensors := make([]Sensor, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		sensorStr, beaconStr, ok := strings.Cut(line, ":")
		if !ok {
			return nil, fmt.Errorf("malformed line: %q", line)
		}

		sensorAt, err := parsePoint(sensorStr)
		if err != nil {
			return nil, fmt.Errorf("parsing sensor point: %w", err)
		}

		beaconAt, err := parsePoint(beaconStr)
		if err != nil {
			return nil, fmt.Errorf("parsing sensor point: %w", err)
		}

		sensors = append(sensors, Sensor{At: sensorAt, Beacon: beaconAt})
	}
	return sensors, nil
}

func parsePoint(str string) (grids.Loc, error) {
	// str is either something like "Sensor at x=2, y=18" or "closest beacon is at x=-2, y=15"
	_, ptStr, ok := strings.Cut(str, " at ")
	if !ok {
		return grids.Loc{}, fmt.Errorf("malformed point string: %q", str)
	}

	// "x=2, y=18" => "x=2" and "y=18"
	xStr, yStr, ok := strings.Cut(ptStr, ", ")
	if !ok {
		return grids.Loc{}, fmt.Errorf("malformed xy string: %q", ptStr)
	}

	var x, y int
	if _, val, ok := strings.Cut(xStr, "="); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			return grids.Loc{}, fmt.Errorf("parsing x: %w", err)
		}
		x = v
	} else {
		return grids.Loc{}, fmt.Errorf("malformed x: %q", xStr)
	}

	if _, val, ok := strings.Cut(yStr, "="); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			return grids.Loc{}, fmt.Errorf("parsing y: %w", err)
		}
		y = v
	} else {
		return grids.Loc{}, fmt.Errorf("malformed y: %q", xStr)
	}

	return grids.Loc{x, y}, nil
}
