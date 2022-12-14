package day16

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/kristofferostlund/adventofcode-2022/pkg/ints"
	"github.com/kristofferostlund/adventofcode-2022/pkg/maps"
	"github.com/kristofferostlund/adventofcode-2022/pkg/queues"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	valves, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	maxTime := 30
	keyFunc := func(s *State) string {
		// We can simplify the cache key for the first part by tracking
		// where we are and how many we've visited beforehand.
		// Because we're looking for the most efficient path, this seems
		// to be good enough to actually give us the best path while not
		// accidentally filtering out bad performers early on.
		// For the example input we go from 87898 to 3670 branches
		// to test. In terms of time, this means we end up spending
		// ~0.3s or so for the real input vs ~20s. Quite the improvement!
		return fmt.Sprintf("%s-%d-%d", s.Position, s.OpenCount(), s.Time)
	}
	endStates := simulateEndStates(valves, maxTime, keyFunc)

	maxPressure := 0
	for _, state := range endStates {
		maxPressure = ints.Max(maxPressure, state.Pressure)
	}

	return maxPressure, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	valves, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	maxTime := 26 // 30 - the 4 minutes it takes to train an elephant.
	keyFunc := func(s *State) string {
		// The key used in Part1 seems to skip certain _locally_ sub-optimal parts
		// which yields a worse result than if we extensively test all paths
		// like we do here. The runtime here is of course worse than that of
		// part 1, but not too bad compared to what I was expecting since the
		// searchspace is smaller (less time allowed).
		return fmt.Sprintf("%s-%b-%d", s.Position, s.BitMask, s.Time)
	}
	endStates := simulateEndStates(valves, maxTime, keyFunc)

	var maxOpenCount float64 = 0
	for _, v := range valves {
		if v.FlowRate > 0 {
			maxOpenCount++
		}
	}
	duoPaths := make(map[uint64]int)
	for _, state := range endStates {
		percentOpened := float64(state.OpenCount()) / maxOpenCount
		// We're aiming for a *rough* 50/50 split, 40/60 is a rough 50/5O!
		if 0.4 <= percentOpened || percentOpened <= 0.6 {
			if pressure, ok := duoPaths[state.BitMask]; !ok || pressure < state.Pressure {
				duoPaths[state.BitMask] = state.Pressure
			}
		}
	}

	max := 0
	for a, aMax := range duoPaths {
		for b, bMax := range duoPaths {
			// We're looking for pairs where there's no overlap at all.
			if a&b == 0 {
				max = ints.Max(max, aMax+bMax)
			}
		}
	}

	return max, nil
}

func simulateEndStates(valves []Valve, maxTime int, keyFunc func(next *State) string) map[string]State {
	endStates := make(map[string]State)

	pq := initPriorityQueue(valves)
	for pq.Len() > 0 {
		state := pq.PopT()

		if state.Time == maxTime-1 {
			// This branch is exhausted, no need to try further!
			// No need to try further here.
			continue
		}

		nextStates := getNextStates(*state)
		for i := range nextStates {
			next := nextStates[i]

			key := keyFunc(&next)
			if state, ok := endStates[key]; !ok || state.Pressure < next.Pressure {
				endStates[key] = next
			} else {
				continue
			}

			pq.PushT(&next, -next.Pressure)
		}
	}
	return endStates
}

func initPriorityQueue(valves []Valve) *queues.PriorityQueue[State] {
	pq := queues.NewPriorityQueue[State]()
	pq.PushT(&State{
		Time:     0,
		Pressure: 0,
		Position: "AA", // Start at AA
		// This is copie around and shared across all states.
		// I'm not entirely sure I like it, but it makes the code a bit simpler.
		valveLookup: maps.LookupOf(valves, func(v Valve) string { return v.ID }),
	}, 0)
	return pq
}

type State struct {
	Time     int
	Pressure int
	Position string
	BitMask  uint64

	valveLookup map[string]Valve
}

func (s *State) OpenValve(v Valve) {
	s.BitMask |= v.B
}

func (s *State) IsOpen(v Valve) bool {
	return s.BitMask&v.B != 0
}

func (s State) Copy() State {
	return State{
		Time:     s.Time,
		Pressure: s.Pressure,
		Position: s.Position,
		BitMask:  s.BitMask,

		valveLookup: s.valveLookup,
	}
}

func (s *State) Valve() Valve {
	return s.valveLookup[s.Position]
}

func (s *State) IncreasePressure() {
	for _, v := range s.OpenedValves() {
		s.Pressure += v.FlowRate
	}
}

func (s *State) OpenCount() int {
	// Counting the 1s in BitMask was a bit faster than calling len(state.OpenedValves()).
	return strings.Count(fmt.Sprintf("%b", s.BitMask), "1")
}

func (s *State) OpenedValves() []Valve {
	valves := make([]Valve, 0)
	for _, v := range s.valveLookup {
		if s.IsOpen(v) {
			valves = append(valves, v)
		}
	}
	return valves
}

type Valve struct {
	ID       string
	FlowRate int
	LeadsTo  []string

	B uint64
}

func (v Valve) String() string {
	return v.ID
}

func getNextStates(state State) []State {
	v := state.Valve()

	nextStates := make([]State, 0, len(v.LeadsTo)+1)
	if !state.IsOpen(v) && v.FlowRate > 0 {
		next := state.Copy()

		// Open the valve
		next.OpenValve(v)
		next.Time++
		next.IncreasePressure()
		nextStates = append(nextStates, next)
	}

	for _, id := range v.LeadsTo {
		next := state.Copy()

		// Move to another valve
		next.Position = id
		next.Time++
		next.IncreasePressure()
		nextStates = append(nextStates, next)
	}

	return nextStates
}

func parseInput(reader io.Reader) ([]Valve, error) {
	valves := make([]Valve, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		v, err := parseValve(line)
		if err != nil {
			return nil, fmt.Errorf("parsing valve: %w", err)
		}

		valves = append(valves, v)
	}

	sort.Slice(valves, func(i, j int) bool {
		return valves[i].ID < valves[j].ID
	})

	var b uint64 = 2
	for i := 0; i < len(valves); i++ {
		b = b << 1
		valves[i].B = b

	}

	return valves, nil
}

func parsePaths(input string) ([]string, error) {
	var others string
	var ok bool

	isPlural := strings.HasPrefix(input, "tunnels")
	if isPlural {
		_, others, ok = strings.Cut(input, "tunnels lead to valves ")
	} else {
		_, others, ok = strings.Cut(input, "tunnel leads to valve ")
	}
	if !ok {
		return nil, fmt.Errorf("malformed path string: %q", input)
	}

	return strings.Split(others, ", "), nil
}

func parseValve(line string) (Valve, error) {
	input, othersStr, ok := strings.Cut(line, "; ")
	if !ok {
		return Valve{}, fmt.Errorf("malformed line %q", line)
	}

	var name string
	var flowRate int
	scannedCount, err := fmt.Sscanf(input, "Valve %s has flow rate=%d", &name, &flowRate)
	if err != nil {
		return Valve{}, fmt.Errorf("scanning %q: %w", input, err)
	}
	if got, want := scannedCount, 2; got != want {
		return Valve{}, fmt.Errorf("malformed valve input %q, got %d scanned values, want %d", input, got, want)
	}

	leadsTo, err := parsePaths(othersStr)
	if err != nil {
		return Valve{}, fmt.Errorf("parsing path string: %w", err)
	}

	return Valve{ID: name, FlowRate: flowRate, LeadsTo: leadsTo}, nil
}
