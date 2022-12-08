package day7

import (
	"fmt"
	"strings"
)

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
