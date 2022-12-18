package day14_test

import (
	"os"
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2022/pkg/relative"
	"github.com/kristofferostlund/adventofcode-2022/puzzles/day14"
)

func TestPuzzle(t *testing.T) {
	exampleInput := `
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`

	t.Run("Part1", func(t *testing.T) {
		t.Run("example input", func(t *testing.T) {
			want := 24
			got, err := day14.Puzzle{}.Part1(strings.NewReader(exampleInput))
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

			want := 964
			got, err := day14.Puzzle{}.Part1(f)
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
			want := 93
			got, err := day14.Puzzle{}.Part2(strings.NewReader(exampleInput))
			if err != nil {
				t.Fatalf("solving part 2: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		// t.Run("input.txt", func(t *testing.T) {
		// 	f, err := os.Open(relative.Filepath("./input.txt"))
		// 	if err != nil {
		// 		t.Fatalf("opening file: %v", err)
		// 	}
		// 	defer f.Close()

		// 	want := 22134
		// 	got, err := day14.Puzzle{}.Part2(f)
		// 	if err != nil {
		// 		t.Fatalf("solving part 1: %v", err)
		// 	}

		// 	if got != want {
		// 		t.Errorf("got %d, want %d", got, want)
		// 	}
		// })
	})
}
