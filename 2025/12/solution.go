package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	input = "2025/12/input.txt"
)

type Query struct {
	w, h   int
	counts []int
}

func parseInput(filename string) ([]int, []Query) {
	file, err := os.Open(filename)
	if err != nil {
		file, err = os.Open("2025/12/" + filename)
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()

	var shapeAreas []int
	var queries []Query
	currentShapeIdx := -1

	scanner := bufio.NewScanner(file)
	parsingShapes := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.Contains(line, "x") && strings.Contains(line, ":") {
			parsingShapes = false
		}

		if parsingShapes {
			if strings.HasSuffix(line, ":") {
				idxStr := strings.TrimSuffix(line, ":")
				idx, _ := strconv.Atoi(idxStr)
				currentShapeIdx = idx
				for len(shapeAreas) <= currentShapeIdx {
					shapeAreas = append(shapeAreas, 0)
				}
			} else {
				count := strings.Count(line, "#")
				if currentShapeIdx != -1 {
					shapeAreas[currentShapeIdx] += count
				}
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
	successCount := 0
	for _, q := range queries {
		totalArea := 0
		for i, qty := range q.counts {
			if i < len(shapeAreas) {
				totalArea += qty * shapeAreas[i]
			}
		}
		if totalArea <= q.w*q.h {
			successCount++
		}
	}
	return successCount
}


func main() {
	fmt.Println(part1())
}