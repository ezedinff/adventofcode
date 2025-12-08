package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	input = "2025/08/input.txt"
	conLimitPart1  = 1000
)

type Point struct{ x, y, z int }

func (p Point) distSq(q Point) int {
	dx, dy, dz := p.x-q.x, p.y-q.y, p.z-q.z
	return dx*dx + dy*dy + dz*dz
}

type DSU struct {
	parent, size []int
}

func newDSU(n int) *DSU {
	d := &DSU{make([]int, n), make([]int, n)}
	for i := range d.parent {
		d.parent[i], d.size[i] = i, 1
	}
	return d
}

func (d *DSU) find(i int) int {
	if d.parent[i] != i {
		d.parent[i] = d.find(d.parent[i])
	}
	return d.parent[i]
}

func (d *DSU) union(i, j int) bool {
	ri, rj := d.find(i), d.find(j)
	if ri == rj {
		return false
	}
	if d.size[ri] < d.size[rj] {
		ri, rj = rj, ri
	}
	d.parent[rj], d.size[ri] = ri, d.size[ri]+d.size[rj]
	return true
}

func readPoints() []Point {
	data, _ := os.ReadFile(input)
	var points []Point
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
		points = append(points, Point{x, y, z})
	}
	return points
}

func buildEdges(points []Point) [][3]int {
	edges := make([][3]int, 0, len(points)*(len(points)-1)/2)
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, [3]int{i, j, points[i].distSq(points[j])})
		}
	}
	sort.Slice(edges, func(a, b int) bool { return edges[a][2] < edges[b][2] })
	return edges
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func part1() int {
	points, edges, dsu := readPoints(), buildEdges(readPoints()), newDSU(len(readPoints()))
	for i := 0; i < min(conLimitPart1, len(edges)); i++ {
		dsu.union(edges[i][0], edges[i][1])
	}
	var sizes []int
	for i := range points {
		if dsu.parent[i] == i {
			sizes = append(sizes, dsu.size[i])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}

func part2() int {
	points, edges, dsu := readPoints(), buildEdges(readPoints()), newDSU(len(readPoints()))
	components := len(points)
	for _, e := range edges {
		if dsu.union(e[0], e[1]) {
			components--
			if components == 1 {
				return points[e[0]].x * points[e[1]].x
			}
		}
	}
	return 0
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}