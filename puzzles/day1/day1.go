package day1

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	groups, err := p.parseGroups(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing groups: %w", err)
	}

	return groups[0], nil
}

func (d Puzzle) Part2(reader io.Reader) (int, error) {
	groups, err := d.parseGroups(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing groups: %w", err)
	}

	sum := 0
	for _, v := range groups[:3] {
		sum += v
	}

	return sum, nil
}

func (Puzzle) parseGroups(reader io.Reader) ([]int, error) {
	groups := make([]int, 0)

	scanner := bufio.NewScanner(reader)
	currGroup := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			groups = append(groups, currGroup)
			sort.Slice(groups, func(i, j int) bool {
				return groups[i] > groups[j]
			})

			currGroup = 0
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("parsing line %q: %w", line, err)
		}
		currGroup += val
	}

	if currGroup > 0 {
		groups = append(groups, currGroup)
		sort.Slice(groups, func(i, j int) bool {
			return groups[i] > groups[j]
		})
	}

	return groups, nil
}
