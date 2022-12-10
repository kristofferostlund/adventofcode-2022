package day8

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	grid, err := p.parseGrid(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing grid: %w", err)
	}

	counter := 0

	for x := 0; x < len(grid); x++ {
		row := grid[x]

		for y := 0; y < len(row); y++ {
			val := row[y]
			sides := 4

			for i := y - 1; i >= 0; i-- {
				if row[i] >= val {
					sides--
					break
				}
			}
			for i := y + 1; i < len(row); i++ {
				if row[i] >= val {
					sides--
					break
				}
			}

			for i := x - 1; i >= 0; i-- {
				if grid[i][y] >= val {
					sides--
					break
				}
			}
			for i := x + 1; i < len(grid); i++ {
				if grid[i][y] >= val {
					sides--
					break
				}
			}

			if sides > 0 {
				counter++
			}
		}
	}

	return counter, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	grid, err := p.parseGrid(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing grid: %w", err)
	}

	max := -1

	for x := 0; x < len(grid); x++ {
		row := grid[x]

		for y := 0; y < len(row); y++ {
			val := row[y]

			d1 := 0
			for i := y - 1; i >= 0; i-- {
				d1++
				if row[i] >= val {
					break
				}
			}

			d2 := 0
			for i := y + 1; i < len(row); i++ {
				d2++
				if row[i] >= val {
					break
				}
			}

			d3 := 0
			for i := x - 1; i >= 0; i-- {
				d3++
				if grid[i][y] >= val {
					break
				}
			}

			d4 := 0
			for i := x + 1; i < len(grid); i++ {
				d4++
				if grid[i][y] >= val {
					break
				}
			}

			viewingDistances := d1 * d2 * d3 * d4
			if viewingDistances > max {
				max = viewingDistances
			}
		}
	}

	return max, nil
}

func (Puzzle) parseGrid(reader io.Reader) ([][]int, error) {
	grid := make([][]int, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		row := make([]int, 0, len(line))
		for _, r := range line {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, fmt.Errorf("parsing %s: %w", string(r), err)
			}

			row = append(row, v)
		}

		grid = append(grid, row)
	}

	return grid, nil
}
