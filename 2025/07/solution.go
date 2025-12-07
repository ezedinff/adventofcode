package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	input = "2025/07/input.txt"
)

type Grid struct {
	lines    []string
	startRow int
	startCol int
	rows     int
	cols     int
}

func readGrid() *Grid {
	data, _ := os.ReadFile(input)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	
	g := &Grid{lines: lines, rows: len(lines), cols: len(lines[0])}
	for r, line := range lines {
		for c, char := range line {
			if char == 'S' {
				g.startRow, g.startCol = r, c
				return g
			}
		}
	}
	return g
}

func (g *Grid) inBounds(c int) bool {
	return c >= 0 && c < g.cols
}

func part1() int {
	grid := readGrid()
	activeCols := map[int]bool{grid.startCol: true}
	splits := 0

	for r := grid.startRow; r < grid.rows; r++ {
		nextCols := make(map[int]bool)
		for c := range activeCols {
			if !grid.inBounds(c) {
				continue
			}
			if grid.lines[r][c] == '^' {
				splits++
				if grid.inBounds(c - 1) {
					nextCols[c-1] = true
				}
				if grid.inBounds(c + 1) {
					nextCols[c+1] = true
				}
			} else {
				nextCols[c] = true
			}
		}
		activeCols = nextCols
	}
	return splits
}

func part2() int {
	grid := readGrid()
	counts := map[int]int{grid.startCol: 1}
	var exited int

	for r := grid.startRow; r < grid.rows; r++ {
		nextCounts := make(map[int]int)
		for c, count := range counts {
			if !grid.inBounds(c) {
				exited += count
				continue
			}
			if grid.lines[r][c] == '^' {
				if grid.inBounds(c - 1) {
					nextCounts[c-1] += count
				} else {
					exited += count
				}
				if grid.inBounds(c + 1) {
					nextCounts[c+1] += count
				} else {
					exited += count
				}
			} else {
				nextCounts[c] += count
			}
		}
		counts = nextCounts
	}
	for _, count := range counts {
		exited += count
	}
	return exited
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}