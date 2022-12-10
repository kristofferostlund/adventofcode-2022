package day7

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Puzzle struct{}

func (p Puzzle) Part1(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)

	tree, err := p.buildTree(scanner)
	if err != nil {
		return 0, fmt.Errorf("building tree: %w", err)
	}

	totalSizeUnder100000 := 0
	tree.WalkDirs(func(path string, dir *Dir) {
		size := dir.Size()
		if size < 100000 {
			totalSizeUnder100000 += size
		}
	})

	return totalSizeUnder100000, nil
}

func (p Puzzle) Part2(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)

	tree, err := p.buildTree(scanner)
	if err != nil {
		return 0, fmt.Errorf("building tree: %w", err)
	}

	totalSize := tree.Size()
	const (
		maxSize      = 70000000
		requiredSize = 30000000
	)

	var smallestDir *Dir

	tree.WalkDirs(func(path string, dir *Dir) {
		if smallestDir == nil {
			smallestDir = dir
			return
		}

		dirSize := dir.Size()
		if maxSize-totalSize+dirSize > requiredSize && dirSize < smallestDir.Size() {
			smallestDir = dir
		}
	})

	return smallestDir.Size(), nil
}

func (Puzzle) buildTree(scanner *bufio.Scanner) (*Dir, error) {
	tree := newDir("", nil)
	tree.AddDir("/")

	dir := tree

	var command string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "$") {
			args := strings.Split(line, " ")[1:]
			command = args[0]

			switch command {
			case "cd":
				dirname := args[1]
				nextDir, err := dir.Cd(dirname)
				if err != nil {
					return nil, fmt.Errorf("finding next dir: %w", err)
				}
				dir = nextDir
				continue
			case "ls":
				// No arguments to ls
			default:
				return nil, fmt.Errorf("illegal command: %q", command)
			}
		} else {
			if command == "ls" {
				a, b, ok := strings.Cut(line, " ")
				if !ok {
					return nil, fmt.Errorf("malformed output for command %q: %q", command, line)
				}

				if a == "dir" {
					dir.AddDir(b)
				} else {
					size, err := strconv.Atoi(a)
					if err != nil {
						return nil, fmt.Errorf("parsing file size of %s: %w", b, err)
					}
					dir.Files[b] = size
				}
			} else {
				return nil, fmt.Errorf("illegal output for command %q: %q", command, line)
			}
		}
	}
	return tree, nil
}
