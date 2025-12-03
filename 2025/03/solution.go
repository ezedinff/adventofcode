package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	input = "2025/03/input.txt"
	minLengthPart1 = 2
	minLengthPart2 = 12
)

func readlines(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	return lines, nil
}

func solveBank(line string) int {
	max := 0
	for i := 0; i < len(line)-1; i++ {
		for j := i + 1; j < len(line); j++ {
			if val := int(line[i]-'0')*10 + int(line[j]-'0'); val > max {
				max = val
			}
		}
	}
	return max
}

func solveBankPart2(line string) int64 {
	if len(line) < 12 {
		return 0
	}

	var result strings.Builder
	pos := -1

	for needed := 12; needed > 0; needed-- {
		best, bestPos := -1, -1
		for i := pos + 1; i <= len(line)-needed; i++ {
			if digit := int(line[i] - '0'); digit > best {
				best, bestPos = digit, i
				if digit == 9 {
					break
				}
			}
		}
		result.WriteByte(line[bestPos])
		pos = bestPos
	}

	val, _ := strconv.ParseInt(result.String(), 10, 64)
	return val
}

func processLines[T int | int64](minLength int, solver func(string) T) T {
	lines, _ := readlines(input)
	var total T
	for _, line := range lines {
		if line = strings.TrimSpace(line); len(line) >= minLength {
			total += solver(line)
		}
	}
	return total
}

func main() {
	fmt.Println("Part 1:", processLines(minLengthPart1, solveBank))
	fmt.Println("Part 2:", processLines(minLengthPart2, solveBankPart2))
}