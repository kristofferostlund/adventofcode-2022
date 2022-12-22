package day16

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/kristofferostlund/adventofcode-2022/pkg/maps"
	"github.com/kristofferostlund/adventofcode-2022/pkg/queues"
	"github.com/kristofferostlund/adventofcode-2022/pkg/sets"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	valves, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	cache := make(map[string]int)

	pq := queues.NewPriorityQueue[State]()
	pq.PushT(&State{
		Time:       0,
		Pressure:   0,
		Position:   "AA", // Start at AA
		OpenValves: sets.NewSet[string](),
		// This is copie around and shared across all states.
		// I'm not entirely sure I like it, but it makes the code a bit simpler.
		valveLookup: maps.LookupOf(valves, func(v Valve) string { return v.ID }),
	}, 0)

	maxTime := 30

	maxPressure := 0
	for pq.Len() > 0 {
		state := pq.PopT()

		if state.Time == maxTime-1 {
			// This branch is exhausted, no need to try further!
			if state.Pressure > maxPressure {
				maxPressure = state.Pressure
			}
			continue
		}

		nextStates := getNextStates(*state)
		for i := range nextStates {
			next := nextStates[i]
			key := cacheKey(&next)

			if pressure, ok := cache[key]; !ok || pressure < next.Pressure {
				cache[key] = next.Pressure
			} else {
				// No need to try further here.
				continue
			}

			// Check the
			pq.PushT(&next, -next.Pressure)
		}
	}

	return maxPressure, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	return 0, nil
}

type State struct {
	Time       int
	Pressure   int
	Position   string
	OpenValves sets.Set[string]

	valveLookup map[string]Valve
}

func (s State) Copy() State {
	return State{
		Time:       s.Time,
		Pressure:   s.Pressure,
		Position:   s.Position,
		OpenValves: s.OpenValves.Copy(),

		valveLookup: s.valveLookup,
	}
}

func (s *State) Valve() Valve {
	return s.valveLookup[s.Position]
}

func (s *State) IncreasePressure() {
	for _, id := range s.OpenValves.Values() {
		s.Pressure += s.valveLookup[id].FlowRate
	}
}

func cacheKey(state *State) string {
	// I first did a thing where I stringified the open valves in order,
	// We're talking like 0.15s vs almost 22s
	// but this yields the same output but an extreme improvement in performance.
	return fmt.Sprintf("%s-%d-%d", state.Position, state.OpenValves.Len(), state.Time)
}

type Valve struct {
	ID       string
	FlowRate int
	LeadsTo  []string
}

func (v Valve) String() string {
	return v.ID
}

func getNextStates(state State) []State {
	v := state.Valve()

	nextStates := make([]State, 0, len(v.LeadsTo)+1)
	if !state.OpenValves.Has(v.ID) && v.FlowRate > 0 {
		next := state.Copy()

		// Open the valve
		next.OpenValves.Add(v.ID)
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
