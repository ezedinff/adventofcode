package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input = "2025/06/input.txt"

type Problem struct {
	Grid []string
	Op   string
}

func parseInput(filename string) []Problem {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")

	maxLen := 0
	for _, line := range lines {
		maxLen = max(maxLen, len(line))
	}
	for i := range lines {
		lines[i] = fmt.Sprintf("%-*s", maxLen, lines[i])
	}

	var problems []Problem
	start := -1

	for col := 0; col <= maxLen; col++ {
		isSpace := col == maxLen || isEmptyColumn(lines, col)

		if !isSpace && start == -1 {
			start = col
		} else if isSpace && start != -1 {
			problems = append(problems, extractProblem(lines, start, col))
			start = -1
		}
	}

	return problems
}

func isEmptyColumn(lines []string, col int) bool {
	for _, line := range lines {
		if line[col] != ' ' {
			return false
		}
	}
	return true
}

func extractProblem(lines []string, start, end int) Problem {
	op := strings.TrimSpace(lines[len(lines)-1][start:end])
	
	var grid []string
	for row := 0; row < len(lines)-1; row++ {
		grid = append(grid, lines[row][start:end])
	}
	
	return Problem{Grid: grid, Op: op}
}

func calculate(numbers []int, op string) int {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]
	for _, num := range numbers[1:] {
		if op == "+" {
			result += num
		} else {
			result *= num
		}
	}
	return result
}

func part1() int {
	problems := parseInput(input)
	total := 0

	for _, p := range problems {
		var numbers []int
		for _, row := range p.Grid {
			numStr := strings.TrimSpace(row)
			if numStr != "" {
				num, _ := strconv.Atoi(numStr)
				numbers = append(numbers, num)
			}
		}
		total += calculate(numbers, p.Op)
	}

	return total
}

func part2() int {
	problems := parseInput(input)
	total := 0

	for _, p := range problems {
		if len(p.Grid) == 0 {
			continue
		}

		var numbers []int
		width := len(p.Grid[0])
		
		for col := width - 1; col >= 0; col-- {
			var numStr strings.Builder
			for row := 0; row < len(p.Grid); row++ {
				if char := p.Grid[row][col]; char != ' ' {
					numStr.WriteByte(char)
				}
			}
			if s := numStr.String(); s != "" {
				num, _ := strconv.Atoi(s)
				numbers = append(numbers, num)
			}
		}
		
		total += calculate(numbers, p.Op)
	}

	return total
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}