package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	input = "2025/04/input.txt"
	rollOfPaper = '@'
	empty       = '.'
)

func readGrid(filename string) [][]rune {
	content, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func isValidPos(grid [][]rune, y, x int) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0])
}

func countNeighbors(grid [][]rune, y, x int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			ny, nx := y+dy, x+dx
			if isValidPos(grid, ny, nx) && grid[ny][nx] == rollOfPaper {
				count++
			}
		}
	}
	return count
}

func part1() int {
	grid := readGrid(input)
	
	count := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == rollOfPaper && countNeighbors(grid, y, x) < 4 {
				count++
			}
		}
	}
	return count
}

func part2() int {
	grid := readGrid(input)

	total := 0
	for {
		var toRemove [][2]int
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] == rollOfPaper && countNeighbors(grid, y, x) < 4 {
					toRemove = append(toRemove, [2]int{y, x})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		total += len(toRemove)
		for _, pos := range toRemove {
			grid[pos[0]][pos[1]] = empty
		}
	}
	return total
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
