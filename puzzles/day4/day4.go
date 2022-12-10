package day4

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2022/pkg/sets"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	pairs, err := p.readPairs(reader)
	if err != nil {
		return 0, fmt.Errorf("reading pairs: %w", err)
	}

	counter := 0
	for _, pair := range pairs {
		a, b := p.setRangeOf(pair[0]), p.setRangeOf(pair[1])
		if a.Contains(b) || b.Contains(a) {
			counter++
		}
	}

	return counter, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	pairs, err := p.readPairs(reader)
	if err != nil {
		return 0, fmt.Errorf("reading pairs: %w", err)
	}

	counter := 0
	for _, pair := range pairs {
		a, b := p.setRangeOf(pair[0]), p.setRangeOf(pair[1])
		if a.Overlaps(b) {
			counter++
		}
	}

	return counter, nil
}

func (p Puzzle) readPairs(reader io.Reader) ([][2][2]int, error) {
	pairs := make([][2][2]int, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		rangeStrA, rangeStrB, ok := strings.Cut(line, ",")
		if !ok {
			return nil, fmt.Errorf("malformed line %q", line)
		}

		rangeA, err := p.parseRange(rangeStrA)
		if err != nil {
			return nil, fmt.Errorf("parsing range A: %w", err)
		}
		rangeB, err := p.parseRange(rangeStrB)
		if err != nil {
			return nil, fmt.Errorf("parsing range B: %w", err)
		}

		pairs = append(pairs, [2][2]int{rangeA, rangeB})
	}

	return pairs, nil
}

func (Puzzle) parseRange(rangeStr string) ([2]int, error) {
	fromStr, toStr, ok := strings.Cut(rangeStr, "-")
	if !ok {
		return [2]int{}, fmt.Errorf("malformed range %q", rangeStr)
	}

	from, err := strconv.Atoi(fromStr)
	if err != nil {
		return [2]int{}, fmt.Errorf("parsing from: %w", err)
	}
	to, err := strconv.Atoi(toStr)
	if err != nil {
		return [2]int{}, fmt.Errorf("parsing to: %w", err)
	}

	return [2]int{from, to}, nil
}

func (Puzzle) setRangeOf(fromTo [2]int) sets.Set[int] {
	vals := make([]int, 0, fromTo[1]-fromTo[0])
	for i := fromTo[0]; i <= fromTo[1]; i++ {
		vals = append(vals, i)
	}
	return sets.Of(vals)
}
