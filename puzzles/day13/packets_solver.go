package day13

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/kristofferostlund/adventofcode-2022/puzzles/day13/packets"
)

/*
This was the solution I started with but didn't manage to get working with the real input...

I felt pretty good about the parser as I found it highly rewarding to build the parser but
ended up feeling like I was walking in circles trying to get it to pass...
I had to start over to realise what was wrong - it wasn't the data structure, which I was
a bit worried about - and quite quickly realised the problem was I returned too early when
comparing the slices!

In the slice solution (used in day13.Puzzle{}) I used *bool where nil would indicate
_keep checking_, which worked but I wasn't happy about it. I ported it over here again
but instead used an ordering state. Felt real good to get it working!
*/

type PacketSolver struct{}

func (p PacketSolver) Part1(reader io.Reader) (int, error) {
	packs, err := packets.Parse(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	counter := 0
	for i := 0; i < len(packs); i += 2 {
		a, b := packs[i], packs[i+1]

		pair := i/2 + 1
		if a.Compare(b) {
			counter += pair
		}
	}

	return counter, nil
}

func (p PacketSolver) Part2(reader io.Reader) (int, error) {
	packs, err := packets.Parse(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	controlPairs := `
[[2]]
[[6]]
`
	controlPackets, err := packets.Parse(strings.NewReader(controlPairs))
	if err != nil {
		return 0, fmt.Errorf("parsing control pairs: %w", err)
	}

	packs = append(packs, controlPackets...)

	sort.Slice(packs, func(i, j int) bool {
		return packs[i].Compare(packs[j])
	})

	idxs := make([]int, 0, 2)
outer:
	for i, s := range packs {
		switch fmt.Sprint(s) {
		case fmt.Sprint(controlPackets[0]), fmt.Sprint(controlPackets[1]):
			idxs = append(idxs, i+1)

			if len(idxs) == 2 {
				break outer
			}
		}
	}

	return idxs[0] * idxs[1], nil
}
