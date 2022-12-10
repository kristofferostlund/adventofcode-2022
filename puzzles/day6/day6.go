package day6

import (
	"fmt"
	"io"

	"github.com/kristofferostlund/adventofcode-2022/pkg/sets"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	seqLen := 4

	return p.findMarker(reader, seqLen)
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	seqLen := 14

	return p.findMarker(reader, seqLen)
}

func (Puzzle) findMarker(reader io.Reader, seqLen int) (int, error) {
	seq, err := io.ReadAll(reader)
	if err != nil {
		return 0, fmt.Errorf("reading input: %w", err)
	}

	for i := seqLen + 1; i < len(seq); i++ {
		set := sets.Of(seq[i-seqLen : i])
		if set.Len() == seqLen {
			return i, nil
		}
	}

	return 0, fmt.Errorf("no marker found")
}
