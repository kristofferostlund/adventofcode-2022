package day10_test

import (
	"os"
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2022/pkg/relative"
	"github.com/kristofferostlund/adventofcode-2022/puzzles/day10"
)

func TestPuzzle(t *testing.T) {
	exampleInput := `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
`

	t.Run("Part1", func(t *testing.T) {
		t.Run("example input", func(t *testing.T) {
			want := 13140
			got, err := day10.Puzzle{}.Part1(strings.NewReader(exampleInput))
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

			want := 15680
			got, err := day10.Puzzle{}.Part1(f)
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})

	t.Run("Part2", func(t *testing.T) {
		t.Run("first example input", func(t *testing.T) {
			wantRendered := strings.Trim(`
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
`, "\n")

			gotRendered := ""
			onRender := func(str string) {
				gotRendered = strings.Trim(str, "\n")
			}

			err := day10.Puzzle{}.Part2(strings.NewReader(exampleInput), onRender)
			if err != nil {
				t.Fatalf("solving part 2: %v", err)
			}
			if gotRendered != wantRendered {
				t.Errorf("got:\n%s\nwant:\n%s", gotRendered, wantRendered)
			}
		})

		t.Run("input.txt", func(t *testing.T) {
			wantRendered := strings.Trim(`
####.####.###..####.#..#..##..#..#.###..
...#.#....#..#.#....#..#.#..#.#..#.#..#.
..#..###..###..###..####.#....#..#.#..#.
.#...#....#..#.#....#..#.#.##.#..#.###..
#....#....#..#.#....#..#.#..#.#..#.#....
####.#....###..#....#..#..###..##..#....
`, "\n") // ZFBFHGUP

			gotRendered := ""
			onRender := func(str string) {
				gotRendered = strings.Trim(str, "\n")
			}

			f, err := os.Open(relative.Filepath("./input.txt"))
			if err != nil {
				t.Fatalf("opening file: %v", err)
			}
			defer f.Close()

			err = day10.Puzzle{}.Part2(f, onRender)
			if err != nil {
				t.Fatalf("solving part 2: %v", err)
			}

			if gotRendered != wantRendered {
				t.Errorf("got:\n%s\nwant:\n%s", gotRendered, wantRendered)
			}
		})
	})
}
