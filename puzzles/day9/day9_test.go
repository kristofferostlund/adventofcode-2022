package day9_test

import (
	"os"
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2022/pkg/relative"
	"github.com/kristofferostlund/adventofcode-2022/puzzles/day9"
)

func TestPuzzle(t *testing.T) {
	firstExampleInput := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

	t.Run("Part1", func(t *testing.T) {
		t.Run("example input", func(t *testing.T) {
			want := 13
			got, err := day9.Puzzle{}.Part1(strings.NewReader(firstExampleInput))
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

			want := 5902
			got, err := day9.Puzzle{}.Part1(f)
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})

	secondExampleInput := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

	t.Run("Part2", func(t *testing.T) {
		t.Run("first example input", func(t *testing.T) {
			want := 1
			got, err := day9.Puzzle{}.Part2(strings.NewReader(firstExampleInput))
			if err != nil {
				t.Fatalf("solving part 2: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
		t.Run("second example input", func(t *testing.T) {
			want := 36
			got, err := day9.Puzzle{}.Part2(strings.NewReader(secondExampleInput))
			if err != nil {
				t.Fatalf("solving part 2: %v", err)
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

			want := 2445
			got, err := day9.Puzzle{}.Part2(f)
			if err != nil {
				t.Fatalf("solving part 2: %v", err)
			}

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})
}
