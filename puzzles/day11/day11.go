package day11

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2022/pkg/sets"
	"github.com/kristofferostlund/adventofcode-2022/pkg/stacks"
	"github.com/kristofferostlund/adventofcode-2022/puzzles/day11/math"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	ops, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	mutateValue := func(val int) int { return val / 3 }

	monkeyBusiness := p.getMonkeyBusiness(ops, 20, mutateValue)
	return monkeyBusiness, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	ops, err := parseInput(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	// This is what the ✨ way ✨ from:
	// > find another way to keep your worry levels manageable
	allDivisors := make([]int, 0)
	for _, op := range ops {
		allDivisors = append(allDivisors, op.divisor)
	}
	lcm := math.LCMSlice(sets.Of(allDivisors).Values())
	mutateValue := func(value int) int { return value % lcm }

	monkeyBusiness := p.getMonkeyBusiness(ops, 10000, mutateValue)
	return monkeyBusiness, nil
}

func (Puzzle) getMonkeyBusiness(ops []*Operation, iterations int, mutateValue func(value int) int) int {
	activity := make([]int, len(ops))

	for r := 0; r < iterations; r++ {
		for i, op := range ops {
			for op.Next() {
				activity[i]++

				value, dest := op.Exec(mutateValue)
				// If any monkey were to throw to itself, this would cause
				// the loop to never finish. Luckily monkeys doesn't juggle
				// so we're fine!
				ops[dest].Append(value)
			}
		}
	}

	sort.Slice(activity, func(i, j int) bool {
		return activity[i] > activity[j]
	})

	return activity[0] * activity[1]
}

type Operation struct {
	items *stacks.Stack[int]

	operation    func(old int) int
	divisor      int
	destinations map[bool]int
}

func (op *Operation) Append(val int) {
	op.items.Push(val)
}

func (op *Operation) Next() bool {
	return op.items.Len() > 0
}

func (op *Operation) Exec(mutateValue func(int) int) (value int, destination int) {
	old := op.items.Shift()
	val := op.operation(old)

	val = mutateValue(val)

	isDivisible := val%op.divisor == 0
	dest := op.destinations[isDivisible]

	return val, dest
}

func parseInput(reader io.Reader) ([]*Operation, error) {
	rawOps := make([]map[string]string, 0)
	current := make(map[string]string, 5)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Monkey") {
			continue
		}

		key, value, ok := strings.Cut(line, ":")
		if !ok {
			return nil, fmt.Errorf("malformed line: %q", line)
		}
		current[strings.TrimSpace(key)] = strings.TrimSpace(value)

		if len(current) == 5 {
			rawOps = append(rawOps, current)
			current = make(map[string]string, 5)
		}
	}

	ops := make([]*Operation, 0, len(rawOps))
	for _, rawOp := range rawOps {
		startingItems, err := parseStartingItems(rawOp["Starting items"])
		if err != nil {
			return nil, fmt.Errorf("parsing starting items: %w", err)
		}

		opFunc, err := prepareOperationFunc(rawOp["Operation"])
		if err != nil {
			return nil, fmt.Errorf("preparing operation func: %w", err)
		}

		divisor, err := parseTestDivisor(rawOp["Test"])
		if err != nil {
			return nil, fmt.Errorf("preparing test func: %w", err)
		}

		destinations, err := parseDestinations(rawOp["If true"], rawOp["If false"])
		if err != nil {
			return nil, fmt.Errorf("parsing outcomes: %w", err)
		}

		ops = append(ops, &Operation{
			items:        stacks.Of(startingItems),
			operation:    opFunc,
			divisor:      divisor,
			destinations: destinations,
		})
	}

	return ops, nil
}

func parseStartingItems(raw string) ([]int, error) {
	if raw == "" {
		return nil, errors.New("must have values")
	}

	items := make([]int, 0)
	for _, v := range strings.Split(raw, ",") {
		trimmed := strings.TrimSpace(v)
		item, err := strconv.Atoi(trimmed)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %w", trimmed, err)
		}
		items = append(items, item)
	}

	return items, nil
}

var operators = map[string]func(a, b int) int{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
}

var literals = map[string]int{}

func prepareOperationFunc(raw string) (func(old int) int, error) {
	trimmed := strings.TrimSpace(raw)
	chunks := strings.Split(trimmed, " ")
	if len(chunks) != 5 {
		return nil, fmt.Errorf("unexpected statement %q", trimmed)
	}

	a, operator, b := chunks[2], chunks[3], chunks[4]
	if _, exists := literals[a]; a != "old" || exists {
		parsed, err := strconv.Atoi(a)
		if err != nil {
			return nil, fmt.Errorf("parsing a: %w", err)
		}
		literals[a] = parsed
	}
	if _, exists := literals[b]; b != "old" || exists {
		parsed, err := strconv.Atoi(b)
		if err != nil {
			return nil, fmt.Errorf("parsing b: %w", err)
		}
		literals[b] = parsed
	}

	operatorFunc, ok := operators[operator]
	if !ok {
		return nil, fmt.Errorf("illegal operator %q", operator)
	}

	return func(old int) int {
		reg := copyMap(literals)
		reg["old"] = old

		return operatorFunc(reg[a], reg[b])
	}, nil
}

func parseTestDivisor(raw string) (int, error) {
	vals := strings.Split(strings.TrimSpace(raw), " ")
	if len(vals) != 3 {
		return 0, fmt.Errorf("unexpected test declaration %q, must have exactly 3 elements", raw)
	}

	testVal, err := strconv.Atoi(vals[2])
	if err != nil {
		return 0, fmt.Errorf("parsing test value: %w", err)
	}

	return testVal, nil
}

func parseDestinations(rawIfTrue, rawIfFalse string) (map[bool]int, error) {
	if rawIfTrue == "" || rawIfFalse == "" {
		return nil, fmt.Errorf("malformed input: %q, %q", rawIfTrue, rawIfFalse)
	}

	parseDest := func(rawStr string) (int, error) {
		_, destStr, ok := strings.Cut(rawStr, "throw to monkey")
		if !ok {
			return 0, fmt.Errorf("malformed dest string %q", rawStr)
		}
		dest, err := strconv.Atoi(strings.TrimSpace(destStr))
		if err != nil {
			return 0, fmt.Errorf("parsing dest string: %w", err)
		}
		return dest, nil
	}

	ifTrue, err := parseDest(rawIfTrue)
	if err != nil {
		return nil, fmt.Errorf("parsing if-true string: %w", err)
	}

	ifFalse, err := parseDest(rawIfFalse)
	if err != nil {
		return nil, fmt.Errorf("parsing if-false string: %w", err)
	}

	return map[bool]int{true: ifTrue, false: ifFalse}, nil
}

func copyMap[K comparable, V any](in map[K]V) map[K]V {
	out := make(map[K]V, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}
