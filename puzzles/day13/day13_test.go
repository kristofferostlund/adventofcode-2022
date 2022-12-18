package day13_test

import (
	"os"
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2022/pkg/relative"
	"github.com/kristofferostlund/adventofcode-2022/puzzles/day13"
)

func TestPuzzle(t *testing.T) {
	exampleInput := `
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`

	t.Run("Part1", func(t *testing.T) {
		t.Run("example input", func(t *testing.T) {
			want := 13
			got, err := day13.Puzzle{}.Part1(strings.NewReader(exampleInput))
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("input.txt", func(t *testing.T) {
			// t.SkipNow()

			f, err := os.Open(relative.Filepath("./input.txt"))
			if err != nil {
				t.Fatalf("opening file: %v", err)
			}
			defer f.Close()

			want := 5196
			got, err := day13.Puzzle{}.Part1(f)
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
			want := 140
			got, err := day13.Puzzle{}.Part2(strings.NewReader(exampleInput))
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

			want := 22134
			got, err := day13.Puzzle{}.Part2(f)
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})
}
