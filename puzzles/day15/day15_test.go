package day15_test

import (
	"os"
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2022/pkg/relative"
	"github.com/kristofferostlund/adventofcode-2022/puzzles/day15"
)

func TestPuzzle(t *testing.T) {
	exampleInput := `
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`

	t.Run("Part1", func(t *testing.T) {
		t.Run("example input", func(t *testing.T) {
			want := 26
			got, err := day15.Puzzle{}.Part1(strings.NewReader(exampleInput), 10, false)
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

			want := 4748135
			got, err := day15.Puzzle{}.Part1(f, 2000000, false)
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})

	// t.Run("Part2", func(t *testing.T) {
	// 	t.Run("example input", func(t *testing.T) {
	// 		want := 140
	// 		got, err := day15.Puzzle{}.Part2(strings.NewReader(exampleInput))
	// 		if err != nil {
	// 			t.Fatalf("solving part 2: %v", err)
	// 		}
	// 		if got != want {
	// 			t.Errorf("got %d, want %d", got, want)
	// 		}
	// 	})

	// 	// t.Run("input.txt", func(t *testing.T) {
	// 	// 	f, err := os.Open(relative.Filepath("./input.txt"))
	// 	// 	if err != nil {
	// 	// 		t.Fatalf("opening file: %v", err)
	// 	// 	}
	// 	// 	defer f.Close()

	// 	// 	want := 22134
	// 	// 	got, err := day15.Puzzle{}.Part2(f)
	// 	// 	if err != nil {
	// 	// 		t.Fatalf("solving part 1: %v", err)
	// 	// 	}

	// 	// 	if got != want {
	// 	// 		t.Errorf("got %d, want %d", got, want)
	// 	// 	}
	// 	// })
	// })
}
