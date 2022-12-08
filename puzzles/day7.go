package puzzles

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

/*
--- Day 7: No Space Left On Device ---

You can hear birds chirping and raindrops hitting leaves as the expedition proceeds. Occasionally, you can even hear much louder sounds in the distance; how big do the animals get out here, anyway?

The device the Elves gave you has problems with more than just its communication system. You try to run a system update:

$ system-update --please --pretty-please-with-sugar-on-top
Error: No space left on device

Perhaps you can delete some files to make space for the update?

You browse around the filesystem to assess the situation and save the resulting terminal output (your puzzle input). For example:

$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k

The filesystem consists of a tree of files (plain data) and directories (which can contain other directories or files). The outermost directory is called /. You can navigate around the filesystem, moving into or out of directories and listing the contents of the directory you're currently in.

Within the terminal output, lines that begin with $ are commands you executed, very much like some modern computers:

    cd means change directory. This changes which directory is the current directory, but the specific result depends on the argument:
        cd x moves in one level: it looks in the current directory for the directory named x and makes it the current directory.
        cd .. moves out one level: it finds the directory that contains the current directory, then makes that directory the current directory.
        cd / switches the current directory to the outermost directory, /.
    ls means list. It prints out all of the files and directories immediately contained by the current directory:
        123 abc means that the current directory contains a file named abc with size 123.
        dir xyz means that the current directory contains a directory named xyz.

Given the commands and output in the example above, you can determine that the filesystem looks visually like this:

- / (dir)
  - a (dir)
    - e (dir)
      - i (file, size=584)
    - f (file, size=29116)
    - g (file, size=2557)
    - h.lst (file, size=62596)
  - b.txt (file, size=14848514)
  - c.dat (file, size=8504156)
  - d (dir)
    - j (file, size=4060174)
    - d.log (file, size=8033020)
    - d.ext (file, size=5626152)
    - k (file, size=7214296)

Here, there are four directories: / (the outermost directory), a and d (which are in /), and e (which is in a). These directories also contain files of various sizes.

Since the disk is full, your first step should probably be to find directories that are good candidates for deletion. To do this, you need to determine the total size of each directory. The total size of a directory is the sum of the sizes of the files it contains, directly or indirectly. (Directories themselves do not count as having any intrinsic size.)

The total sizes of the directories above can be found as follows:

    The total size of directory e is 584 because it contains a single file i of size 584 and no other directories.
    The directory a has total size 94853 because it contains files f (size 29116), g (size 2557), and h.lst (size 62596), plus file i indirectly (a contains e which contains i).
    Directory d has total size 24933642.
    As the outermost directory, / contains every file. Its total size is 48381165, the sum of the size of every file.

To begin, find all of the directories with a total size of at most 100000, then calculate the sum of their total sizes. In the example above, these directories are a and e; the sum of their total sizes is 95437 (94853 + 584). (As in this example, this process can count files more than once!)

Find all of the directories with a total size of at most 100000. What is the sum of the total sizes of those directories?

--- Part Two ---

Now, you're ready to choose a directory to delete.

The total disk space available to the filesystem is 70000000. To run the update, you need unused space of at least 30000000. You need to find a directory you can delete that will free up enough space to run the update.

In the example above, the total size of the outermost directory (and thus the total amount of used space) is 48381165; this means that the size of the unused space must currently be 21618835, which isn't quite the 30000000 required by the update. Therefore, the update still requires a directory with total size of at least 8381165 to be deleted before it can run.

To achieve this, you have the following options:

    Delete directory e, which would increase unused space by 584.
    Delete directory a, which would increase unused space by 94853.
    Delete directory d, which would increase unused space by 24933642.
    Delete directory /, which would increase unused space by 48381165.

Directories e and a are both too small; deleting them would not free up enough space. However, directories d and / are both big enough! Between these, choose the smallest: d, increasing unused space by 24933642.

Find the smallest directory that, if deleted, would free up enough space on the filesystem to run the update. What is the total size of that directory?
*/

type Day7 struct{}

func (d Day7) Part1(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)

	tree, err := d.buildTree(scanner)
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

func (d Day7) Part2(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)

	tree, err := d.buildTree(scanner)
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

func (Day7) buildTree(scanner *bufio.Scanner) (*Dir, error) {
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

type Dir struct {
	Name  string
	Files map[string]int
	dirs  map[string]*Dir
}

func newDir(name string, parent *Dir) *Dir {
	return &Dir{
		Name:  name,
		Files: make(map[string]int),
		dirs:  map[string]*Dir{"..": parent},
	}
}

func (d *Dir) Cd(dirname string) (*Dir, error) {
	if nextDir := d.dirs[dirname]; nextDir != nil {
		return nextDir, nil
	}
	return nil, fmt.Errorf("no such dir %s", dirname)
}

func (d *Dir) Parent() *Dir {
	return d.dirs[".."]
}

func (d *Dir) Dirs() map[string]*Dir {
	dirs := make(map[string]*Dir, len(d.dirs)-1)
	for dirname, dir := range d.dirs {
		if dirname == ".." {
			continue
		}
		dirs[dirname] = dir
	}
	return dirs
}

func (d *Dir) AddDir(dirname string) {
	d.dirs[dirname] = newDir(dirname, d)
}

func (d Dir) Size() int {
	size := 0
	for _, fileSize := range d.Files {
		size += fileSize
	}
	for _, dir := range d.Dirs() {
		size += dir.Size()
	}
	return size
}

func (d Dir) String() string {
	padding := &strings.Builder{}
	for dd := &d; dd.Parent() != nil; dd = dd.Parent() {
		padding.WriteString("  ")
	}

	sb := &strings.Builder{}

	for dirname, dir := range d.Dirs() {
		sb.WriteString(fmt.Sprintf("%s- %s (dir, total size=%d)\n", padding, dirname, dir.Size()))
		sb.WriteString(fmt.Sprintf("%s%s", padding, dir))
	}

	for filename, size := range d.Files {
		sb.WriteString(fmt.Sprintf("%s- %s (file, size=%d)\n", padding, filename, size))
	}

	return sb.String()
}

func (d Dir) Path() string {
	path := &strings.Builder{}
	for dd := &d; dd.Parent() != nil; dd = dd.Parent() {
		if dd.Name == "/" || dd.Name == "" {
			path.WriteString(dd.Name)
		} else {
			path.WriteString(fmt.Sprintf("/%s", dd.Name))
		}
	}
	return path.String()
}

func (d Dir) WalkDirs(onDir func(path string, dir *Dir)) {
	for _, dir := range d.Dirs() {
		onDir(dir.Path(), dir)
	}

	for _, dir := range d.Dirs() {
		dir.WalkDirs(onDir)
	}
}

func (d Dir) WalkFiles(onFile func(path string, size int)) {
	for filename, size := range d.Files {
		onFile(d.Path()+filename, size)
	}
	for _, dir := range d.Dirs() {
		dir.WalkFiles(onFile)
	}
}
