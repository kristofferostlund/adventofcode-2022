package puzzles_test

import (
	"os"
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2022/pkg/relative"
	"github.com/kristofferostlund/adventofcode-2022/puzzles"
)

func TestDay5(t *testing.T) {
	exampleInput := `
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

	t.Run("Part1", func(t *testing.T) {
		t.Run("example input", func(t *testing.T) {
			want := "CMZ"
			got, err := puzzles.Day5{}.Part1(strings.NewReader(exampleInput))
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})

		t.Run("input/day5.txt", func(t *testing.T) {
			f, err := os.Open(relative.Filepath("./input/day5.txt"))
			if err != nil {
				t.Fatalf("opening file: %v", err)
			}
			defer f.Close()

			want := "FZCMJCRHZ"
			got, err := puzzles.Day5{}.Part1(f)
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	})

	t.Run("Part2", func(t *testing.T) {
		t.Run("example input", func(t *testing.T) {
			want := "MCD"
			got, err := puzzles.Day5{}.Part2(strings.NewReader(exampleInput))
			if err != nil {
				t.Fatalf("solving part 2: %v", err)
			}
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})

		t.Run("input/day5.txt", func(t *testing.T) {
			f, err := os.Open(relative.Filepath("./input/day5.txt"))
			if err != nil {
				t.Fatalf("opening file: %v", err)
			}
			defer f.Close()

			want := "JSDHQMZGF"
			got, err := puzzles.Day5{}.Part2(f)
			if err != nil {
				t.Fatalf("solving part 2: %v", err)
			}

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	})
}
