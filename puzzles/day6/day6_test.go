package day6_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2022/pkg/relative"
	"github.com/kristofferostlund/adventofcode-2022/puzzles/day6"
)

func TestPuzzle(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		examples := []struct {
			sequence string
			want     int
		}{
			{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
			{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
			{"nppdvjthqldpwncqszvftbrmjlhg", 6},
			{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
			{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
		}

		t.Run("example input", func(t *testing.T) {
			for _, example := range examples {
				exampleInput, want := example.sequence, example.want
				t.Run(fmt.Sprintf("%s -> %d", exampleInput, want), func(t *testing.T) {
					got, err := day6.Puzzle{}.Part1(strings.NewReader(exampleInput))
					if err != nil {
						t.Fatalf("solving part 1: %v", err)
					}
					if got != want {
						t.Errorf("got %d, want %d", got, want)
					}
				})
			}
		})

		t.Run("input.txt", func(t *testing.T) {
			f, err := os.Open(relative.Filepath("./input.txt"))
			if err != nil {
				t.Fatalf("opening file: %v", err)
			}
			defer f.Close()

			want := 1794
			got, err := day6.Puzzle{}.Part1(f)
			if err != nil {
				t.Fatalf("solving part 1: %v", err)
			}

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})

	t.Run("Part2", func(t *testing.T) {
		examples := []struct {
			sequence string
			want     int
		}{
			{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
			{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
			{"nppdvjthqldpwncqszvftbrmjlhg", 23},
			{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
			{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
		}

		t.Run("example input", func(t *testing.T) {
			for _, example := range examples {
				exampleInput, want := example.sequence, example.want
				t.Run(fmt.Sprintf("%s -> %d", exampleInput, want), func(t *testing.T) {
					got, err := day6.Puzzle{}.Part2(strings.NewReader(exampleInput))
					if err != nil {
						t.Fatalf("solving part 2: %v", err)
					}
					if got != want {
						t.Errorf("got %d, want %d", got, want)
					}
				})
			}
		})

		t.Run("input.txt", func(t *testing.T) {
			f, err := os.Open(relative.Filepath("./input.txt"))
			if err != nil {
				t.Fatalf("opening file: %v", err)
			}
			defer f.Close()

			want := 2851
			got, err := day6.Puzzle{}.Part2(f)
			if err != nil {
				t.Fatalf("solving part 2: %v", err)
			}

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})
}
