package day2

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Puzzle struct {
	scores        map[string]int
	opponentMoves map[string]string
	beatenBy      map[string]string
}

func NewPuzzle() Puzzle {
	return Puzzle{
		scores: map[string]int{
			"rock":     1,
			"paper":    2,
			"scissors": 3,

			"win":  6,
			"draw": 3,
			"loss": 0,
		},
		opponentMoves: map[string]string{
			"A": "rock",
			"B": "paper",
			"C": "scissors",
		},
		beatenBy: map[string]string{
			"rock":     "scissors",
			"paper":    "rock",
			"scissors": "paper",
		},
	}
}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	pairs, err := p.readStrategy(reader)
	if err != nil {
		return 0, fmt.Errorf("reading strategy: %w", err)
	}

	playerMoves := map[string]string{
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	score := 0
	for _, pair := range pairs {
		score += p.scoreOf(p.opponentMoves[pair[0]], playerMoves[pair[1]])
	}

	return score, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	pairs, err := p.readStrategy(reader)
	if err != nil {
		return 0, fmt.Errorf("reading strategy: %w", err)
	}

	beats := make(map[string]string, len(p.beatenBy))
	for k, v := range p.beatenBy {
		beats[v] = k
	}

	outcomes := map[string]string{
		"X": "loss",
		"Y": "draw",
		"Z": "win",
	}

	score := 0
	for _, pair := range pairs {
		opponentMove := p.opponentMoves[pair[0]]

		var playerMove string
		switch outcomes[pair[1]] {
		case "loss":
			playerMove = p.beatenBy[opponentMove]
		case "win":
			playerMove = beats[opponentMove]
		case "draw":
			playerMove = opponentMove
		default:
			return 0, fmt.Errorf("unknown outcome for %q", pair[1])
		}

		score += p.scoreOf(opponentMove, playerMove)
	}

	return score, nil
}

func (Puzzle) readStrategy(reader io.Reader) ([][2]string, error) {
	pairs := make([][2]string, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		a, b, ok := strings.Cut(line, " ")
		if !ok {
			return nil, fmt.Errorf("malformed line %q", line)
		}
		pairs = append(pairs, [2]string{a, b})
	}

	return pairs, nil
}

func (p Puzzle) scoreOf(a, b string) int {
	moveScore := p.scores[b]
	if a == b {
		return p.scores["draw"] + moveScore
	}
	if p.beatenBy[a] == b {
		return p.scores["loss"] + moveScore
	}

	return p.scores["win"] + moveScore
}
