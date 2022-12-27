package day18

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	points, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	grid3d := gridOf(points)

	noncovered := 0
	for _, pt := range points {
		// There are 6 sides to a cube.
		noncovered += 6 - len(grid3d.SiblingsOf(pt))
	}

	return noncovered, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	return 0, nil
}

type Grid3d struct {
	points map[Point3D]struct{}
}

func gridOf(points []Point3D) *Grid3d {
	g := &Grid3d{points: make(map[Point3D]struct{})}
	for _, pt := range points {
		g.Set(pt)
	}
	return g
}

var dirs = []Point3D{
	{1, 0, 0},  // Left
	{-1, 0, 0}, // Right
	{0, 1, 0},  // Up
	{0, -1, 0}, // Down
	{0, 0, 1},  // In front of
	{0, 0, -1}, // Behind
}

func (g *Grid3d) SiblingsOf(point Point3D) []Point3D {
	siblings := make([]Point3D, 0)
	for _, dir := range dirs {
		sibling := point.Add(dir)
		if g.Has(sibling) {
			siblings = append(siblings, sibling)
		}
	}
	return siblings
}

func (g *Grid3d) Has(point Point3D) bool {
	_, has := g.points[point]
	return has
}

func (g *Grid3d) Set(point Point3D) {
	g.points[point] = struct{}{}
}

type Point3D [3]int

func (p Point3D) X() int {
	return p[0]
}

func (p Point3D) Y() int {
	return p[1]
}

func (p Point3D) Z() int {
	return p[2]
}

func (p Point3D) Add(other Point3D) Point3D {
	return Point3D{p[0] + other[0], p[1] + other[1], p[2] + other[2]}
}

func (p Point3D) XYZ() (x, y, z int) {
	return p[0], p[1], p[2]
}

func parseInput(reader io.Reader) ([]Point3D, error) {
	points := make([]Point3D, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		ss := strings.Split(line, ",")
		if got, want := len(ss), 3; got != want {
			return nil, fmt.Errorf("got %d dimensions, want %d", got, want)
		}

		point := Point3D{}
		for i, ds := range ss {
			val, err := strconv.Atoi(ds)
			if err != nil {
				return nil, fmt.Errorf("parsing %q: %w", ds, err)
			}
			point[i] = val
		}
		points = append(points, point)
	}
	return points, nil
}
