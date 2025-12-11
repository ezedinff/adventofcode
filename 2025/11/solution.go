package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	input = "2025/11/input.txt"
)

func parseInput() map[string][]string {
	data, _ := os.ReadFile(input)
	adj := make(map[string][]string)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		parts := strings.Split(line, ": ")
		adj[parts[0]] = strings.Fields(parts[1])
	}
	return adj
}

func countPaths(adj map[string][]string, start, end string, memo map[string]int) int {
	if start == end {
		return 1
	}
	if val, ok := memo[start]; ok {
		return val
	}
	count := 0
	for _, next := range adj[start] {
		count += countPaths(adj, next, end, memo)
	}
	memo[start] = count
	return count
}

func part1(adj map[string][]string) int {
	return countPaths(adj, "you", "out", make(map[string]int))
}

func part2(adj map[string][]string) int {
	// svr -> dac -> fft -> out
	p1 := countPaths(adj, "svr", "dac", make(map[string]int)) * 
	      countPaths(adj, "dac", "fft", make(map[string]int)) * 
	      countPaths(adj, "fft", "out", make(map[string]int))

	// svr -> fft -> dac -> out
	p2 := countPaths(adj, "svr", "fft", make(map[string]int)) * 
	      countPaths(adj, "fft", "dac", make(map[string]int)) * 
	      countPaths(adj, "dac", "out", make(map[string]int))
	return p1 + p2
}

func main() {
	adj := parseInput()
	fmt.Println("Part 1:", part1(adj))
	fmt.Println("Part 2:", part2(adj))
}