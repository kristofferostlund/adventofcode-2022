package day10

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	cmdNoop = "noop"
	cmdAddX = "addx"
)

var cmdCycles = map[string]int{
	cmdNoop: 1,
	cmdAddX: 2,
}

type operation struct {
	cmd string
	val int
}

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	ops, err := parseOps(reader)
	if err != nil {
		return 0, fmt.Errorf("parsing operations: %w", err)
	}

	x := 1
	counter := 0

	opI := 0
	nextAt := 1
	var op operation
	shouldBreak := false
	for c := 1; !shouldBreak; c++ {
		if c == nextAt {
			x += op.val

			if opI < len(ops) {
				op = ops[opI]
				opI++
				nextAt += cmdCycles[op.cmd]
			} else {
				shouldBreak = true
			}
		}

		if (c-20)%40 == 0 {
			counter += x * c
		}
	}

	return counter, nil
}

func (p Puzzle) Part2(reader io.Reader, onRender func(str string)) error {
	ops, err := parseOps(reader)
	if err != nil {
		return fmt.Errorf("parsing operations: %w", err)
	}

	x := 1

	rows := [6][40]rune{}
	rowi, linei := 0, 0

	lineLen := len(rows[0])
	rowLen := len(rows)

	opI := 0
	nextAt := 1
	var op operation
	for c := 1; ; c++ {
		if c == nextAt {
			x += op.val

			if opI < len(ops) {
				op = ops[opI]
				opI++
				nextAt += cmdCycles[op.cmd]
			} else {
				break
			}
		}

		if shouldRenderX(linei, x, lineLen) {
			rows[rowi][linei] = '#'
		} else {
			rows[rowi][linei] = '.'
		}

		if linei+1 == lineLen {
			rowi = (rowi + 1) % rowLen
		}
		linei = (linei + 1) % lineLen
	}

	sb := &strings.Builder{}
	for _, row := range rows {
		for _, r := range row {
			sb.WriteRune(r)
		}
		sb.WriteRune('\n')
	}

	onRender(sb.String())

	return nil
}

func shouldRenderX(i, x, max int) bool {
	if i == (max+x-1)%max {
		return true
	}
	if i == (x+1)%max {
		return true
	}
	return i == x
}

func parseOps(reader io.Reader) ([]operation, error) {
	ops := make([]operation, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		cmd := strings.Split(line, " ")
		if len(cmd) == 0 {
			return nil, fmt.Errorf("malformed line: %q", line)
		}

		switch cmd[0] {
		case cmdNoop:
			ops = append(ops, operation{cmd: cmdNoop})
		case cmdAddX:
			if len(cmd) != 2 {
				return nil, fmt.Errorf("malformed line: %q", line)
			}
			v, err := strconv.Atoi(cmd[1])
			if err != nil {
				return nil, fmt.Errorf("parsing value: %w", err)
			}
			ops = append(ops, operation{cmd: cmdAddX, val: v})
		default:
			return nil, fmt.Errorf("illegal op: %q", cmd[0])
		}
	}

	return ops, nil
}
