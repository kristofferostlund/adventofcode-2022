package debug

import (
	"math"
	"strings"
)

func Render(points [][2]int) string {
	minX, maxX := math.MaxInt64, -math.MaxInt
	minY, maxY := math.MaxInt64, -math.MaxInt

	pps := make(map[[2]int]struct{}, len(points))
	for _, p := range points {
		pps[p] = struct{}{}

		if p[0] < minX {
			minX = p[0]
		} else if maxX < p[0] {
			maxX = p[0]
		}

		if p[1] < minY {
			minY = p[1]
		} else if maxY < p[1] {
			maxY = p[1]
		}
	}

	sb := &strings.Builder{}
	for x := maxX; x >= minX; x-- {
		sb.WriteRune('\n')
		for y := minY; y <= maxY+1; y++ {
			if x == 0 && y == 0 {
				sb.WriteRune('s')
				continue
			}

			if _, ok := pps[[2]int{x, y}]; ok {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
	}

	return sb.String()
}
