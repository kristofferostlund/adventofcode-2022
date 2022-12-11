package day11

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2022/pkg/stacks"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int64, error) {
	ops, err := parseInput(reader, func(val int64) int64 { return val / 3 })
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	activity := make([]int64, len(ops))

	for r := 0; r < 20; r++ {
		for i, op := range ops {
			for op.Next() {
				activity[i]++

				value, dest := op.Exec()
				// If any monkey were to throw to itself, this would cause
				// the loop to never finish.
				ops[dest].Append(value)
			}
		}
	}

	sort.Slice(activity, func(i, j int) bool {
		return activity[i] > activity[j]
	})

	return activity[0] * activity[1], nil
}

func (p Puzzle) Part2(reader io.Reader) (int64, error) {
	return 0, nil
}

type Operation struct {
	items *stacks.Stack[int64]
	// operation combines both the Operation and Test directives from the exercise.
	operation func(old int64) (value int64, destination int)
}

func (op *Operation) Append(val int64) {
	op.items.Push(val)
}

func (op *Operation) Next() bool {
	return op.items.Len() > 0
}

func (op *Operation) Exec() (value int64, destination int) {
	return op.operation(op.items.Shift())
}

func parseInput(reader io.Reader, reduceWorry func(value int64) int64) ([]*Operation, error) {
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

		testFunc, err := prepareTestFunc(rawOp["Test"])
		if err != nil {
			return nil, fmt.Errorf("preparing test func: %w", err)
		}

		outcomes, err := parseOutcomes(rawOp["If true"], rawOp["If false"])
		if err != nil {
			return nil, fmt.Errorf("parsing outcomes: %w", err)
		}

		ops = append(ops, &Operation{
			items: stacks.Of(startingItems),
			operation: func(old int64) (value int64, destination int) {
				val := opFunc(old)

				val = reduceWorry(val)

				result := testFunc(val)

				dest := outcomes[result]

				return val, dest
			},
		})
	}

	return ops, nil
}

func parseStartingItems(raw string) ([]int64, error) {
	if raw == "" {
		return nil, errors.New("must have values")
	}

	items := make([]int64, 0)
	for _, v := range strings.Split(raw, ",") {
		trimmed := strings.TrimSpace(v)
		item, err := strconv.ParseInt(trimmed, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %w", trimmed, err)
		}
		items = append(items, item)
	}

	return items, nil
}

var operators = map[string]func(a, b int64) int64{
	"+": func(a, b int64) int64 { return a + b },
	"*": func(a, b int64) int64 { return a * b },
}

var literals = map[string]int64{}

func prepareOperationFunc(raw string) (func(old int64) int64, error) {
	trimmed := strings.TrimSpace(raw)
	chunks := strings.Split(trimmed, " ")
	if len(chunks) != 5 {
		return nil, fmt.Errorf("unexpected statement %q", trimmed)
	}

	a, operator, b := chunks[2], chunks[3], chunks[4]
	if _, exists := literals[a]; a != "old" || exists {
		parsed, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parsing a: %w", err)
		}
		literals[a] = parsed
	}
	if _, exists := literals[b]; b != "old" || exists {
		parsed, err := strconv.ParseInt(b, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parsing b: %w", err)
		}
		literals[b] = parsed
	}

	operatorFunc, ok := operators[operator]
	if !ok {
		return nil, fmt.Errorf("illegal operator %q", operator)
	}

	return func(old int64) int64 {
		reg := copyMap(literals)
		reg["old"] = old

		return operatorFunc(reg[a], reg[b])
	}, nil
}

var tests = map[string]func(want int64) func(value int64) bool{
	"divisible": func(divisibleBy int64) func(value int64) bool {
		return func(value int64) bool {
			return value%divisibleBy == 0
		}
	},
}

func prepareTestFunc(raw string) (func(value int64) bool, error) {
	vals := strings.Split(strings.TrimSpace(raw), " ")
	if len(vals) != 3 {
		return nil, fmt.Errorf("unexpected test declaration %q, must have exactly 3 elements", raw)
	}

	testVal, err := strconv.ParseInt(vals[2], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parsing test value: %w", err)
	}

	testFunc, ok := tests[vals[0]]
	if !ok {
		return nil, fmt.Errorf("no such test: %q", vals[0])
	}

	return testFunc(testVal), nil
}

func parseOutcomes(rawIfTrue, rawIfFalse string) (map[bool]int, error) {
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
