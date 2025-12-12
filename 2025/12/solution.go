package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input = "2025/12/input.txt"

type Query struct {
	w, h   int
	counts []int
}

func parseInput(input string) ([]int, []Query) {
	data, _ := os.ReadFile(input)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var shapeAreas []int
	var queries []Query
	parsingShapes := true
	current := -1

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "x") && strings.Contains(line, ":") {
			parsingShapes = false
		}
		if parsingShapes {
			if strings.HasSuffix(line, ":") {
				idx, _ := strconv.Atoi(strings.TrimSuffix(line, ":"))
				current = idx
				for len(shapeAreas) <= current {
					shapeAreas = append(shapeAreas, 0)
				}
			} else {
				shapeAreas[current] += strings.Count(line, "#")
			}
		} else {
			parts := strings.Split(line, ":")
			dims := strings.Split(strings.TrimSpace(parts[0]), "x")
			w, _ := strconv.Atoi(dims[0])
			h, _ := strconv.Atoi(dims[1])
			countsStr := strings.Fields(parts[1])
			counts := make([]int, len(countsStr))
			for i, s := range countsStr {
				counts[i], _ = strconv.Atoi(s)
			}
			queries = append(queries, Query{w, h, counts})
		}
	}

	return shapeAreas, queries
}

func part1() int {
	shapeAreas, queries := parseInput(input)
	success := 0
	for _, q := range queries {
		total := 0
		for i, qty := range q.counts {
			if i < len(shapeAreas) {
				total += qty * shapeAreas[i]
			}
		}
		area := q.w * q.h
		if total <= area {
			success++
		}
	}
	return success
}

func main() {
	fmt.Println(part1())
}
