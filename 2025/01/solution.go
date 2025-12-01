package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	input    = "input.txt"
	startsAt = 50
)

func readInput(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func floorDiv(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}

func part1() int {
	content, err := readInput(input)
	if err != nil {
		fmt.Println("error reading file:", err)
		return 0
	}
	lines := strings.Split(strings.TrimSpace(content), "\n")
	currentPos := startsAt
	password := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		valStr := line[1:]
		val, err := strconv.Atoi(valStr)
		if err != nil {
			fmt.Printf("Error parsing:", err)
			continue
		}

		if strings.HasPrefix(line, "L") {
			currentPos = (currentPos - val) % 100
			if currentPos < 0 {
				currentPos += 100
			}
		} else if strings.HasPrefix(line, "R") {
			currentPos = (currentPos + val) % 100
		}

		if currentPos == 0 {
			password++
		}
	}
	return password
}


func part2() int {
	content, err := readInput(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	lines := strings.Split(strings.TrimSpace(content), "\n")
	currentPos := startsAt
	password := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		valStr := line[1:]
		val, err := strconv.Atoi(valStr)
		if err != nil {
			fmt.Printf("Error parsing:", err)
			continue
		}

		if strings.HasPrefix(line, "L") {
			hits := floorDiv(currentPos-1, 100) - floorDiv(currentPos-val-1, 100)
			password += hits
			currentPos -= val
		} else if strings.HasPrefix(line, "R") {
			hits := floorDiv(currentPos+val, 100) - floorDiv(currentPos, 100)
			password += hits
			currentPos += val
		}

		currentPos = currentPos % 100
		if currentPos < 0 {
			currentPos += 100
		}
	}
	return password
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
