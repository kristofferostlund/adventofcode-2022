package day13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	ss, err := parseAsSlices(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	counter := 0
	for i := 0; i < len(ss); i += 2 {
		a, b := ss[i], ss[i+1]

		pair := i/2 + 1
		isCorrect := compare(a, b, 0)
		if *isCorrect {
			counter += pair
		}
	}

	return counter, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	ss, err := parseAsSlices(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	controlPairs := `
[[2]]
[[6]]
`
	css, err := parseAsSlices(strings.NewReader(controlPairs))
	if err != nil {
		return 0, fmt.Errorf("parsing control pairs: %w", err)
	}

	ss = append(ss, css...)

	sort.Slice(ss, func(i, j int) bool {
		return *compare(ss[i], ss[j], 0)
	})

	idxs := make([]int, 0, 2)
outer:
	for i, s := range ss {
		switch fmt.Sprint(s) {
		case fmt.Sprint(css[0]), fmt.Sprint(css[1]):
			idxs = append(idxs, i+1)

			if len(idxs) == 2 {
				break outer
			}
		}
	}

	return idxs[0] * idxs[1], nil
}

func parseAsSlices(reader io.Reader) ([][]any, error) {
	var out [][]any
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		var l []any
		if err := json.Unmarshal(line, &l); err != nil {
			return nil, fmt.Errorf("decoding: %w", err)
		}

		out = append(out, l)
	}

	return out, nil
}

var (
	bTrue  = true
	bFalse = false
	True   = &bTrue  // Snakey.
	False  = &bFalse // Snakey.
)

func compare(left, right []any, nestLevel int) *bool {
	lLen, rLen := len(left), len(right)

	maxIter := lLen
	if rLen > maxIter {
		maxIter = rLen
	}

	for i := 0; i < maxIter; i++ {
		if i == lLen {
			return True
		} else if i == rLen {
			return False
		}

		l, r := left[i], right[i]

		lNum, lIsNum := l.(float64)
		rNum, rIsNum := r.(float64)

		if lIsNum && rIsNum {
			if lNum < rNum {
				return True
			} else if lNum > rNum {
				return False
			}

			continue
		}

		var a, b []any
		switch {
		case lIsNum && !rIsNum:
			a, b = []any{lNum}, r.([]any)
		case !lIsNum && rIsNum:
			a, b = l.([]any), []any{rNum}
		default:
			a, b = l.([]any), r.([]any)
		}

		result := compare(a, b, nestLevel+1)
		if result == nil {
			continue
		}

		return result
	}

	if nestLevel > 0 {
		// Continue iterating. Could be made better...
		return nil
	}

	return True
}
