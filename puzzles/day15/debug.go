package day15

import "github.com/kristofferostlund/adventofcode-2022/pkg/grids"

const (
	empty    = "."
	sensor   = "S"
	beacon   = "B"
	noBeacon = "#"
)

func prepareGrid(sensors []Sensor) *grids.Grid[string] {
	grid := grids.NewGrid(empty)
	for _, s := range sensors {
		grid.Set(s.At, sensor)
		grid.Set(s.Beacon, beacon)
	}

	return grid
}

func addNoBeaconZones(grid *grids.Grid[string], sensors []Sensor) {
	for _, s := range sensors {
		manhattan := s.ManhattanDistance()
		for md := manhattan; md > 0; md-- {
			x, y := s.At.XY()

			up := grids.Loc{x, y - md}
			right := grids.Loc{x + md, y}
			down := grids.Loc{x, y + md}
			left := grids.Loc{x - md, y}

			dirs := []grids.Loc{up, right, down, left}
			for i, dir := range dirs {
				next := dirs[(i+1)%len(dirs)]

				for loc := dir; loc != next; loc = stepTowards(loc, next) {
					if _, isUsed := grid.At(loc); !isUsed {
						grid.Set(loc, noBeacon)
					}
				}
			}
		}
	}
}

func stepTowards(at, to grids.Loc) grids.Loc {
	x, y := 0, 0
	if at[0] > to[0] {
		x = -1
	}
	if at[0] < to[0] {
		x = 1
	}

	if at[1] > to[1] {
		y = -1
	}
	if at[1] < to[1] {
		y = 1
	}

	return at.Add(grids.Loc{x, y})
}
