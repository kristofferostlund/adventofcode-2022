package day4_test

import (
	"os"
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2022/pkg/relative"
	"github.com/kristofferostlund/adventofcode-2022/puzzles/day4"
)

func TestPuzzle(t *testing.T) {
	exampleInput := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

	t.Run("Part1", func(t *testing.T) {
		t.Run("example input", func(t *testing.T) {
			want := 2
			got, err := day4.Puzzle{}.Part1(strings.NewReader(exampleInput))
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("input.txt", func(t *testing.T) {
			f, err := os.Open(relative.Filepath("./input.txt"))
			if err != nil {
				t.Fatalf("opening file: %v", err)
			}
			defer f.Close()

			want := 651
			got, err := day4.Puzzle{}.Part1(f)
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})

	t.Run("Part2", func(t *testing.T) {
		t.Run("example input", func(t *testing.T) {
			want := 4
			got, err := day4.Puzzle{}.Part2(strings.NewReader(exampleInput))
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("input.txt", func(t *testing.T) {
			f, err := os.Open(relative.Filepath("./input.txt"))
			if err != nil {
				t.Fatalf("opening file: %v", err)
			}
			defer f.Close()

			want := 956
			got, err := day4.Puzzle{}.Part2(f)
			if err != nil {
				t.Fatalf("solving part 2: %v", err)
			}

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})
}
