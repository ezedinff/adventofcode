package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	input = "2025/02/input.txt"
)

func readInput(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// used for part 1
func isInvalid(id int) bool {
	idStr := strconv.Itoa(id)
	length := len(idStr)
	if length%2 != 0 {
		return false
	}
	half := length / 2
	return idStr[:half] == idStr[half:]
}

func isRepeatedPattern(s string, l int) bool {
	pattern := s[:l]
	for i := l; i < len(s); i += l {
		if s[i:i+l] != pattern {
			return false
		}
	}
	return true
}

// used for part 2
func isInvalid2(id int) bool {
	s := strconv.Itoa(id)
	n := len(s)
	for l := 1; l <= n/2; l++ {
		if n%l == 0 && isRepeatedPattern(s, l) {
			return true
		}
	}
	return false
}


func processRange(idRange string, validator func(int) bool) int {
	idRange = strings.TrimSpace(idRange)
	if idRange == "" {
		return 0
	}

	ids := strings.Split(idRange, "-")
	if len(ids) != 2 {
		fmt.Println("invalid range format:", idRange)
		return 0
	}

	start, err := strconv.Atoi(ids[0])
	if err != nil {
		fmt.Println("error parsing start:", err)
		return 0
	}

	end, err := strconv.Atoi(ids[1])
	if err != nil {
		fmt.Println("error parsing end:", err)
		return 0
	}

	sum := 0
	for id := start; id <= end; id++ {
		if validator(id) {
			sum += id
		}
	}
	return sum
}

func processRanges(content string, validator func(int) bool) int {
	idRanges := strings.Split(strings.TrimSpace(content), ",")
	sum := 0

	for _, idRange := range idRanges {
		sum += processRange(idRange, validator)
	}

	return sum
}

func part1() int {
	content, err := readInput(input)
	if err != nil {
		fmt.Println("error reading file:", err)
		return 0
	}
	return processRanges(content, isInvalid)
}

func part2() int {
	content, err := readInput(input)
	if err != nil {
		fmt.Println("error reading file:", err)
		return 0
	}
	return processRanges(content, isInvalid2)
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}