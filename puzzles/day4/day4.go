package day4

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2022/pkg/sets"
)

/*
--- Day 4: Camp Cleanup ---

Space needs to be cleared before the last supplies can be unloaded from the ships, and so several Elves have been assigned the job of cleaning up sections of the camp. Every section has a unique ID number, and each Elf is assigned a range of section IDs.

However, as some of the Elves compare their section assignments with each other, they've noticed that many of the assignments overlap. To try to quickly find overlaps and reduce duplicated effort, the Elves pair up and make a big list of the section assignments for each pair (your puzzle input).

For example, consider the following list of section assignment pairs:

2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8

For the first few pairs, this list means:

    Within the first pair of Elves, the first Elf was assigned sections 2-4 (sections 2, 3, and 4), while the second Elf was assigned sections 6-8 (sections 6, 7, 8).
    The Elves in the second pair were each assigned two sections.
    The Elves in the third pair were each assigned three sections: one got sections 5, 6, and 7, while the other also got 7, plus 8 and 9.

This example list uses single-digit section IDs to make it easier to draw; your actual list might contain larger numbers. Visually, these pairs of section assignments look like this:

.234.....  2-4
.....678.  6-8

.23......  2-3
...45....  4-5

....567..  5-7
......789  7-9

.2345678.  2-8
..34567..  3-7

.....6...  6-6
...456...  4-6

.23456...  2-6
...45678.  4-8

Some of the pairs have noticed that one of their assignments fully contains the other. For example, 2-8 fully contains 3-7, and 6-6 is fully contained by 4-6. In pairs where one assignment fully contains the other, one Elf in the pair would be exclusively cleaning sections their partner will already be cleaning, so these seem like the most in need of reconsideration. In this example, there are 2 such pairs.

In how many assignment pairs does one range fully contain the other?

--- Part Two ---

It seems like there is still quite a bit of duplicate work planned. Instead, the Elves would like to know the number of pairs that overlap at all.

In the above example, the first two pairs (2-4,6-8 and 2-3,4-5) don't overlap, while the remaining four pairs (5-7,7-9, 2-8,3-7, 6-6,4-6, and 2-6,4-8) do overlap:

    5-7,7-9 overlaps in a single section, 7.
    2-8,3-7 overlaps all of the sections 3 through 7.
    6-6,4-6 overlaps in a single section, 6.
    2-6,4-8 overlaps in sections 4, 5, and 6.

So, in this example, the number of overlapping assignment pairs is 4.

In how many assignment pairs do the ranges overlap?
*/

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
