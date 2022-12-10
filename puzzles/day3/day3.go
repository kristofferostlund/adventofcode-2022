package day3

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kristofferostlund/adventofcode-2022/pkg/sets"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	pairs, err := p.readPairs(reader)
	if err != nil {
		return 0, fmt.Errorf("reading pairs: %w", err)
	}

	prios := 0
	for _, pair := range pairs {
		a, b := p.runeSetOf(pair[0]), p.runeSetOf(pair[1])
		intersection := a.Intersection(b)
		if wantLen := 1; intersection.Len() != wantLen {
			return 0, fmt.Errorf("got intersection of %v, want length %d", intersection.Values(), wantLen)
		}

		prio, err := p.priorityOf(intersection.Values()[0])
		if err != nil {
			return 0, fmt.Errorf("getting priority of %q", intersection.Values()[0])
		}
		prios += prio
	}

	return prios, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	lines, err := p.readLines(reader)
	if err != nil {
		return 0, fmt.Errorf("reading pairs: %w", err)
	}

	prios := 0
	for _, group := range p.groupsOf(lines) {
		a, b, c := group[0], group[1], group[2]
		intersection := a.Intersection(b).Intersection(c)
		if wantLen := 1; intersection.Len() != wantLen {
			return 0, fmt.Errorf("got intersection of %v, want length %d", intersection.Values(), wantLen)
		}

		prio, err := p.priorityOf(intersection.Values()[0])
		if err != nil {
			return 0, fmt.Errorf("getting priority of %q", intersection.Values()[0])
		}
		prios += prio
	}

	return prios, nil
}

func (p Puzzle) groupsOf(lines []string) [][3]sets.Set[rune] {
	groups := make([][3]sets.Set[rune], 0)

	const i = 0
	for len(lines) > 0 {
		a := p.runeSetOf(lines[i])

	second:
		for j := i + 1; j < len(lines); j++ {
			b := p.runeSetOf(lines[j])
			intersection := a.Intersection(b)
			if intersection.Len() < 1 {
				continue second
			}

		third:
			for k := j + 1; k < len(lines); k++ {
				c := p.runeSetOf(lines[k])
				if intrsctn := intersection.Intersection(c); intrsctn.Len() != 1 {
					continue third
				}

				groups = append(groups, [3]sets.Set[rune]{a, b, c})

				updated := make([]string, 0, len(lines)-3)
				updated = append(updated, lines[i+1:j]...)
				updated = append(updated, lines[j+1:k]...)
				updated = append(updated, lines[k+1:]...)

				lines = updated
				break second
			}
		}
	}
	return groups
}

func (Puzzle) readPairs(reader io.Reader) ([][2]string, error) {
	pairs := make([][2]string, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		a, b := line[:len(line)/2], line[len(line)/2:]
		pairs = append(pairs, [2]string{a, b})
	}

	return pairs, nil
}

func (Puzzle) readLines(reader io.Reader) ([]string, error) {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			lines = append(lines, line)
		}
	}
	return lines, nil
}

func (Puzzle) runeSetOf(s string) sets.Set[rune] {
	return sets.Of([]rune(s))
}

func (Puzzle) priorityOf(char rune) (int, error) {
	switch {
	case char >= 'a' && char <= 'z':
		return 1 + int(byte(char)-byte('a')), nil
	case char >= 'A' && char <= 'Z':
		return 27 + int(byte(char)-byte('A')), nil
	default:
		return 0, fmt.Errorf("illegal char %q", char)
	}
}
