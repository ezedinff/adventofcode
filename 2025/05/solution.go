package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	input = "2025/05/input.txt"
)

type Range struct{ Start, End int }

func parseInput(filename string) ([]Range, []int) {
	data, _ := os.ReadFile(filename)
	parts := strings.Split(string(data), "\n\n")

	var ranges []Range
	for _, line := range strings.Split(strings.TrimSpace(parts[0]), "\n") {
		var r Range
		fmt.Sscanf(line, "%d-%d", &r.Start, &r.End)
		ranges = append(ranges, r)
	}

	var ids []int
	for _, line := range strings.Split(strings.TrimSpace(parts[1]), "\n") {
		id, _ := strconv.Atoi(strings.TrimSpace(line))
		ids = append(ids, id)
	}

	return ranges, ids
}

func inRange(id int, ranges []Range) bool {
	left, right := 0, len(ranges) - 1
	for left <= right {
		mid := (left + right) / 2
		if id >= ranges[mid].Start && id <= ranges[mid].End {
			return true
		}
		if id < ranges[mid].Start {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func mergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return nil
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := []Range{ranges[0]}
	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]
		if r.Start <= last.End+1 {
			last.End = max(last.End, r.End)
		} else {
			merged = append(merged, r)
		}
	}
	return merged
}

func countIDs(ranges []Range) int {
	total := 0
	for _, r := range ranges {
		total += r.End - r.Start + 1
	}
	return total
}

func part1(ranges []Range, ids []int) int {
	merged := mergeRanges(ranges)
	count := 0
	for _, id := range ids {
		if inRange(id, merged) {
			count++
		}
	}
	return count
}

func part2(ranges []Range) int {
	return countIDs(mergeRanges(ranges))
}

func main() {
	ranges, ids := parseInput(input)
	fmt.Println("Part 1:", part1(ranges, ids))
	fmt.Println("Part 2:", part2(ranges))
}