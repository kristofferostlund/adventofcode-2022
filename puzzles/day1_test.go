package puzzles_test

import (
	"os"
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2022/pkg/relative"
	"github.com/kristofferostlund/adventofcode-2022/puzzles"
)

var exampleInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestDay1_Part1(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		want := 24000
		got, err := puzzles.Day1{}.Part1(strings.NewReader(exampleInput))
		if err != nil {
			t.Fatalf("solving part 1: %v", err)
		}
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("input/day1.txt", func(t *testing.T) {
		f, err := os.Open(relative.Filepath("./input/day1.txt"))
		if err != nil {
			t.Fatalf("opening file: %v", err)
		}
		defer f.Close()

		want := 72511
		got, err := puzzles.Day1{}.Part1(f)
		if err != nil {
			t.Fatalf("solving part 1: %v", err)
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestDay1_Part2(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		want := 45000
		got, err := puzzles.Day1{}.Part2(strings.NewReader(exampleInput))
		if err != nil {
			t.Fatalf("solving part 2: %v", err)
		}
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("input/day1.txt", func(t *testing.T) {
		f, err := os.Open(relative.Filepath("./input/day1.txt"))
		if err != nil {
			t.Fatalf("opening file: %v", err)
		}
		defer f.Close()

		want := 212117
		got, err := puzzles.Day1{}.Part2(f)
		if err != nil {
			t.Fatalf("solving part 2: %v", err)
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
