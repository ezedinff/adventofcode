package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

import "adventofcode/2025/utils"

const (
	input = "09/input.txt"
)

type Point struct{ x, y int }

type Rect struct {
	minX, maxX, minY, maxY int
}

func readPoints(file string) []Point {
	data, _ := os.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	
	var points []Point
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{x, y})
	}
	return points
}

func isInsidePolygon(p1, p2 Point, polygon []Point) bool {
	rect := Rect{
		minX: utils.Min(p1.x, p2.x),
		maxX: utils.Max(p1.x, p2.x),
		minY: utils.Min(p1.y, p2.y),
		maxY: utils.Max(p1.y, p2.y),
	}
	
	cx := float64(rect.minX+rect.maxX) / 2.0
	cy := float64(rect.minY+rect.maxY) / 2.0
	if !pointInPolygon(cx, cy, polygon) {
		return false
	}
	
	for i := 0; i < len(polygon); i++ {
		next := (i + 1) % len(polygon)
		if edgeCrossesRect(polygon[i], polygon[next], rect) {
			return false
		}
	}
	
	return true
}

func pointInPolygon(x, y float64, poly []Point) bool {
	inside := false
	j := len(poly) - 1
	
	for i := 0; i < len(poly); i++ {
		xi, yi := float64(poly[i].x), float64(poly[i].y)
		xj, yj := float64(poly[j].x), float64(poly[j].y)
		
		if ((yi > y) != (yj > y)) && (x < (xj-xi)*(y-yi)/(yj-yi)+xi) {
			inside = !inside
		}
		j = i
	}
	
	return inside
}

func edgeCrossesRect(p1, p2 Point, r Rect) bool {
	isVertical := p1.x == p2.x
	isHorizontal := p1.y == p2.y
	
	if isVertical {
		withinX := p1.x > r.minX && p1.x < r.maxX
		
		segMin, segMax := utils.Min(p1.y, p2.y), utils.Max(p1.y, p2.y)
		overlapsY := utils.Max(segMin, r.minY) < utils.Min(segMax, r.maxY)
		
		return withinX && overlapsY
	}
	
	if isHorizontal {
		withinY := p1.y > r.minY && p1.y < r.maxY
		
		segMin, segMax := utils.Min(p1.x, p2.x), utils.Max(p1.x, p2.x)
		overlapsX := utils.Max(segMin, r.minX) < utils.Min(segMax, r.maxX)
		
		return withinY && overlapsX
	}
	
	return false
}

func part1(points []Point) int {
	maxArea := 0
	
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			width := utils.Abs(points[i].x-points[j].x) + 1
			height := utils.Abs(points[i].y-points[j].y) + 1
			area := width * height
			
			if area > maxArea {
				maxArea = area
			}
		}
	}
	
	return maxArea
}

func part2(points []Point) int {
	maxArea := 0
	
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			
			width := utils.Abs(p1.x-p2.x) + 1
			height := utils.Abs(p1.y-p2.y) + 1
			area := width * height
			
			if area > maxArea && isInsidePolygon(p1, p2, points) {
				maxArea = area
			}
		}
	}
	
	return maxArea
}

func main() {
	points := readPoints(input)
	fmt.Println("Part 1:", part1(points))
	fmt.Println("Part 2:", part2(points))
}