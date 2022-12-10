package day9

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2022/pkg/sets"
)

type Puzzle struct{}

var directions = map[string][2]int{
	"R": {0, 1},
	"L": {0, -1},
	"U": {1, 0},
	"D": {-1, 0},
}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	knots := make([][2]int, 2)
	visits, err := tailVisits(reader, knots)
	if err != nil {
		return 0, fmt.Errorf("getting tail visits: %w", err)
	}

	return len(visits), nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	knots := make([][2]int, 10)
	visits, err := tailVisits(reader, knots)
	if err != nil {
		return 0, fmt.Errorf("getting tail visits: %w", err)
	}

	return len(visits), nil
}

func tailVisits(reader io.Reader, knots [][2]int) ([][2]int, error) {
	scanner := bufio.NewScanner(reader)
	tVisits := sets.Of([][2]int{knots[len(knots)-1]})
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		v, mut, err := handleInstruction(line)
		if err != nil {
			return nil, fmt.Errorf("handling instruction: %w", err)
		}

		knots[0][0] += mut[0] * v
		knots[0][1] += mut[1] * v

		for i := 0; i < v; i++ {
			for ki := 1; ki < len(knots); ki++ {
				h, t := knots[ki-1], knots[ki]

				if absInt(h[0]-t[0]) > 1 || absInt(h[1]-t[1]) > 1 {
					step := stepTowards(t, h)

					t[0] += step[0]
					t[1] += step[1]
					knots[ki] = t

					if ki == len(knots)-1 && !tVisits.Has(t) {
						tVisits.Add(t)
					}
				}
			}
		}
	}
	return tVisits.Values(), nil
}

func handleInstruction(line string) (int, [2]int, error) {
	d, i, ok := strings.Cut(line, " ")
	if !ok {
		return 0, [2]int{}, fmt.Errorf("malformed line: %q", line)
	}

	v, err := strconv.Atoi(i)
	if err != nil {
		return 0, [2]int{}, fmt.Errorf("parsing %q: %w", i, err)
	}

	mut, ok := directions[d]
	if !ok {
		return 0, [2]int{}, fmt.Errorf("illegal direction: %q", d)
	}
	return v, mut, nil
}

func absInt(v int) int {
	if v > 0 {
		return v
	}
	return v * -1
}

func stepTowards(at, to [2]int) [2]int {
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

	return [2]int{x, y}
}
