package puzzles_test

import (
	"os"
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2022/pkg/relative"
	"github.com/kristofferostlund/adventofcode-2022/puzzles"
)

func TestDay2(t *testing.T) {
	exampleInput := `A Y
B X
C Z
`

	t.Run("Part1", func(t *testing.T) {
		t.Run("example input", func(t *testing.T) {
			want := 15
			got, err := puzzles.NewDay2().Part1(strings.NewReader(exampleInput))
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("input/day2.txt", func(t *testing.T) {
			f, err := os.Open(relative.Filepath("./input/day2.txt"))
			if err != nil {
				t.Fatalf("opening file: %v", err)
			}
			defer f.Close()

			want := 12740
			got, err := puzzles.NewDay2().Part1(f)
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
			want := 12
			got, err := puzzles.NewDay2().Part2(strings.NewReader(exampleInput))
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("input/day2.txt", func(t *testing.T) {
			f, err := os.Open(relative.Filepath("./input/day2.txt"))
			if err != nil {
				t.Fatalf("opening file: %v", err)
			}
			defer f.Close()

			want := 11980
			got, err := puzzles.NewDay2().Part2(f)
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})
}
