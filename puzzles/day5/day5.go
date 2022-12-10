package day5

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (string, error) {
	stacks, instructions, err := p.parseInput(reader)
	if err != nil {
		return "", fmt.Errorf("parsing input: %w", err)
	}

	for _, inst := range instructions {
		count, from, to := inst[0], inst[1], inst[2]

		values := stacks[from].PopN(count)
		stacks[to].Push(values...)
	}

	sb := &strings.Builder{}
	for i := 1; i <= len(stacks); i++ {
		sb.WriteString(stacks[i].Pop())
	}

	return sb.String(), nil
}

func (p Puzzle) Part2(reader io.Reader) (string, error) {
	stacks, instructions, err := p.parseInput(reader)
	if err != nil {
		return "", fmt.Errorf("parsing input: %w", err)
	}

	for _, inst := range instructions {
		count, from, to := inst[0], inst[1], inst[2]

		values := stacks[from].PopN(count)
		for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
			values[i], values[j] = values[j], values[i]
		}

		stacks[to].Push(values...)
	}

	sb := &strings.Builder{}
	for i := 1; i <= len(stacks); i++ {
		sb.WriteString(stacks[i].Pop())
	}

	return sb.String(), nil
}

func (p Puzzle) parseInput(reader io.Reader) (map[int]*Stack[string], [][3]int, error) {
	stacks := make(map[int]*Stack[string], 0)
	instructions := make([][3]int, 0)

	state := "parsing_stacks"
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		// For example input to work in tests...
		if line == "" && len(stacks) == 0 {
			continue
		}

		if state == "parsing_stacks" {
			if line == "" && len(stacks) > 0 {
				// We've read all stacks, time to read instructions!
				state = "parsing_instructions"
				continue
			}
			if line[:2] == " 1" {
				// Reading the line with the numbers.
				// Ignoring this as the convention is easy enough.
				continue
			}

			letters := p.getIndexedLetters(line)
			for idx, letter := range letters {
				if stacks[idx] == nil {
					stacks[idx] = &Stack[string]{}
				}
				stacks[idx].Unshift(letter)
			}
		} else if state == "parsing_instructions" {
			if line == "" {
				continue
			}

			count, from, to, err := p.parseInstruction(line)
			if err != nil {
				return nil, nil, fmt.Errorf("parsing instructions")
			}

			instructions = append(instructions, [3]int{count, from, to})
		}
	}

	return stacks, instructions, nil
}

func (Puzzle) getIndexedLetters(line string) map[int]string {
	idx := 1
	letters := make(map[int]string, 0)
	for i := 0; i < len(line); i++ {
		if (i-1)%4 == 0 {
			if line[i] != ' ' {
				letters[idx] = string(line[i])
			}
			idx++
		}
	}
	return letters
}

func (Puzzle) parseInstruction(line string) (int, int, int, error) {
	ss := strings.Split(line, " ")
	if len(ss) != 6 {
		return 0, 0, 0, fmt.Errorf("malformed instruction line %q", line)
	}

	count, err := strconv.Atoi(ss[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("parsing count: %w", err)
	}
	from, err := strconv.Atoi(ss[3])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("parsing from: %w", err)
	}
	to, err := strconv.Atoi(ss[5])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("parsing to: %w", err)
	}

	return count, from, to, nil
}
