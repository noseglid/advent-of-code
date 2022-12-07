package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type FileType int

const (
	Dir FileType = iota
	File
)

type files []*file

func (ff files) Len() int {
	return len(ff)
}
func (ff files) Less(i, j int) bool {
	return ff[i].df() < ff[j].df()
}
func (ff files) Swap(i, j int) {
	ff[i], ff[j] = ff[j], ff[i]
}

type file struct {
	tt      FileType
	name    string
	size    int
	parent  *file
	content []*file
}

func (f file) find(name string) *file {
	for _, c := range f.content {
		if c.name == name {
			return c
		}
	}
	panic("file not found: " + name)
}

func (f file) stringInternal(indent int) string {
	var sb strings.Builder
	sb.WriteString(strings.Repeat(" ", indent))
	sb.WriteString(fmt.Sprintf("- %s ", f.name))
	switch f.tt {
	case Dir:
		sb.WriteString("(dir)")
	case File:
		sb.WriteString(fmt.Sprintf("(file, size=%d)", f.size))
	}
	sb.WriteRune('\n')

	for _, c := range f.content {
		sb.WriteString(c.stringInternal(indent + 2))
	}
	return sb.String()
}

func (f file) String() string {
	return f.stringInternal(0)
}

func (f file) df() int {
	switch f.tt {
	case File:
		return f.size
	case Dir:
		s := 0
		for _, c := range f.content {
			s += c.df()
		}
		return s
	}
	panic("bad type")
}

func (f *file) foldersWithMax(m int) []*file {
	var files []*file
	if f.df() < m {
		files = append(files, f)
	}

	for _, c := range f.content {
		if c.tt != Dir {
			continue
		}
		files = append(files, c.foldersWithMax(m)...)
	}

	return files
}

func main() {
	input := util.GetFileStrings("2022/Day7/input")

	root := &file{
		tt:   Dir,
		name: "root",
	}

	current := root

LL:
	for _, l := range input[1:] {
		parts := strings.Split(l, " ")
		switch parts[0] {
		case "dir":
			d := &file{
				tt:     Dir,
				name:   parts[1],
				parent: current,
			}
			current.content = append(current.content, d)
		case "$":
			switch parts[1] {
			case "ls":
				continue LL
			case "cd":
				if parts[2] == ".." {
					current = current.parent
				} else {
					current = current.find(parts[2])
				}
			}
		default:
			f := &file{
				tt:     File,
				name:   parts[1],
				parent: current,
				size:   util.MustAtoi(parts[0]),
			}
			current.content = append(current.content, f)
		}
	}

	s := 0
	for _, f := range root.foldersWithMax(100000) {
		s += f.df()
	}

	fmt.Printf("Sum of small folders (part1): %d\n", s)

	maxSpace := 70000000
	required := 30000000
	used := root.df()
	mustFree := required - (maxSpace - used)

	allFoldersSize := root.foldersWithMax(math.MaxInt)
	sort.Sort(files(allFoldersSize))
	for _, f := range allFoldersSize {
		if mustFree-f.df() < 0 {
			fmt.Printf("Smalles to delete (part2) is %s: %d", f.name, f.df())
			break
		}
	}
}
