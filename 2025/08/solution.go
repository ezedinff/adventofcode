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

type Edge struct {
	u, v, dist int
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

func buildEdges(points []Point) []Edge {
	edges := make([]Edge, 0, len(points)*(len(points)-1)/2)
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{i, j, points[i].distSq(points[j])})
		}
	}
	sort.Slice(edges, func(a, b int) bool { return edges[a].dist < edges[b].dist })
	return edges
}

func part1() int {
	points, edges, dsu := readPoints(), buildEdges(readPoints()), newDSU(len(readPoints()))
	for i := 0; i < conLimitPart1; i++ {
		dsu.union(edges[i].u, edges[i].v)
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
		if dsu.union(e.u, e.v) {
			components--
			if components == 1 {
				return points[e.u].x * points[e.v].x
			}
		}
	}
	return 0
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
