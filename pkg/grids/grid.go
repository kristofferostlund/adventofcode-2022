package grids

import (
	"fmt"
	"math"
	"strings"
)

type Grid[T comparable] struct {
	bounds Bounds

	values   map[Loc]T
	emptyVal T
}

func NewGrid[T comparable](emptyVal T) *Grid[T] {
	emptyBounds := NewBounds(math.MaxInt, math.MinInt, math.MaxInt, math.MinInt)

	grid := &Grid[T]{
		bounds:   emptyBounds,
		values:   make(map[Loc]T),
		emptyVal: emptyVal,
	}

	return grid
}

func (g *Grid[T]) At(at Loc) (T, bool) {
	value, ok := g.values[at]
	return value, ok
}

func (g *Grid[T]) String() string {
	return g.RenderArea(g.bounds)
}

func (g *Grid[T]) RenderRow(y int) string {
	sb := &strings.Builder{}
	for x := g.bounds.MinX(); x <= g.bounds.MaxX(); x++ {
		at := Loc{x, y}
		value, ok := g.At(at)
		if !ok {
			value = g.emptyVal
		}
		sb.WriteString(fmt.Sprint(value))
	}

	return sb.String()
}

func (g *Grid[T]) RenderArea(bounds Bounds) string {
	sb := &strings.Builder{}
	for y := bounds.MinY(); y <= bounds.MaxY(); y++ {
		for x := bounds.MinX(); x <= bounds.MaxX(); x++ {
			at := Loc{x, y}
			value, ok := g.At(at)
			if !ok {
				value = g.emptyVal
			}
			sb.WriteString(fmt.Sprint(value))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (g *Grid[T]) Count(value T) int {
	counter := 0
	for _, v := range g.values {
		if v == value {
			counter++
		}
	}
	return counter
}

func (g *Grid[T]) Set(loc Loc, value T) {
	g.bounds = g.bounds.Extend(loc)
	g.values[loc] = value
}

func (g *Grid[T]) Bounds() Bounds {
	return g.bounds
}

func (g *Grid[T]) InBounds(loc Loc) bool {
	return g.bounds.IsInside(loc)
}
