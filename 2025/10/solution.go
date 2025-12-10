package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	input = "2025/10/input.txt"
	epsilon = 1e-9 // 0.000000001
	boundPart2 = 300
)

type Machine struct {
	Buttons [][]int
	Target1 []int
	Target2 []int
}

func parse(path string) []Machine {
	data, _ := os.ReadFile(path)
	var machines []Machine
	re := regexp.MustCompile(`\(([\d,]+)\)`)

	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		if line == "" {
			continue
		}

		//target1 [###...]
		i1, i2 := strings.Index(line, "["), strings.Index(line, "]")
		t1 := make([]int, i2-i1-1)
		for i, c := range line[i1+1 : i2] {
			if c == '#' {
				t1[i] = 1
			}
		}

		// buttons
		btns := make([][]int, 0)
		for _, m := range re.FindAllStringSubmatch(line, -1) {
			btn := make([]int, len(t1))
			for _, s := range strings.Split(m[1], ",") {
				if idx, _ := strconv.Atoi(strings.TrimSpace(s)); idx < len(t1) {
					btn[idx] = 1
				}
			}
			btns = append(btns, btn)
		}

		// target2 {...}
		j1, j2 := strings.Index(line, "{"), strings.Index(line, "}")
		t2 := make([]int, 0)
		for _, s := range strings.Split(line[j1+1:j2], ",") {
			n, _ := strconv.Atoi(strings.TrimSpace(s))
			t2 = append(t2, n)
		}

		machines = append(machines, Machine{btns, t1, t2})
	}
	return machines
}

func solvePart1(btns [][]int, target []int) int {
	rows, cols := len(target), len(btns)
	mat := make([][]int, rows)
	for r := 0; r < rows; r++ {
		mat[r] = make([]int, cols+1)
		for c := 0; c < cols; c++ {
			mat[r][c] = btns[c][r]
		}
		mat[r][cols] = target[r]
	}

	pivots := make([]int, cols)
	for i := range pivots {
		pivots[i] = -1
	}
	pr := 0

	for c := 0; c < cols && pr < rows; c++ {
		sel := -1
		for r := pr; r < rows; r++ {
			if mat[r][c] == 1 {
				sel = r
				break
			}
		}
		if sel == -1 {
			continue
		}

		mat[pr], mat[sel] = mat[sel], mat[pr]

		for r := 0; r < rows; r++ {
			if r != pr && mat[r][c] == 1 {
				for k := c; k <= cols; k++ {
					mat[r][k] ^= mat[pr][k]
				}
			}
		}
		pivots[c] = pr
		pr++
	}

	for r := pr; r < rows; r++ {
		if mat[r][cols] == 1 {
			return 0
		}
	}

	var free []int
	for c, p := range pivots {
		if p == -1 {
			free = append(free, c)
		}
	}

	minW := math.MaxInt32
	for i := 0; i < (1 << len(free)); i++ {
		x := make([]int, cols)
		w := 0
		for j, f := range free {
			if (i>>j)&1 == 1 {
				x[f] = 1
				w++
			}
		}
		for c := cols - 1; c >= 0; c-- {
			if pivots[c] != -1 {
				val := mat[pivots[c]][cols]
				for k := c + 1; k < cols; k++ {
					if mat[pivots[c]][k] == 1 {
						val ^= x[k]
					}
				}
				x[c] = val
				w += val
			}
		}
		if w < minW {
			minW = w
		}
	}

	if minW == math.MaxInt32 {
		return 0
	}
	return minW
}

func findMinWeight(idx int, x []float64, free []int, cols int, pivots []int, mat [][]float64, btns [][]int, target []int, minW int) int {
	if idx == len(free) {
		temp := make([]float64, cols)
		copy(temp, x)
		w := 0
		for c := cols - 1; c >= 0; c-- {
			if pivots[c] != -1 {
				val := mat[pivots[c]][cols]
				for k := c + 1; k < cols; k++ {
					val -= mat[pivots[c]][k] * temp[k]
				}
				if val < -epsilon || math.Abs(val-math.Round(val)) > epsilon {
					return minW
				}
				temp[c] = math.Round(val)
			}
			if temp[c] < -epsilon {
				return minW
			}
			w += int(temp[c])
		}
		if w < minW {
			return w
		}
		return minW
	}

	bound := boundPart2
	for r := 0; r < len(target); r++ {
		if btns[free[idx]][r] == 1 && target[r] < bound {
			bound = target[r]
		}
	}
	
	currentMin := minW
	for v := 0; v <= bound; v++ {
		x[free[idx]] = float64(v)
		res := findMinWeight(idx+1, x, free, cols, pivots, mat, btns, target, currentMin)
		if res < currentMin {
			currentMin = res
		}
	}
	return currentMin
}

func solvePart2(btns [][]int, target []int) int {
	rows, cols := len(target), len(btns)
	mat := make([][]float64, rows)
	for r := 0; r < rows; r++ {
		mat[r] = make([]float64, cols+1)
		for c := 0; c < cols; c++ {
			mat[r][c] = float64(btns[c][r])
		}
		mat[r][cols] = float64(target[r])
	}

	pivots := make([]int, cols)
	for i := range pivots {
		pivots[i] = -1
	}
	pr := 0
	eps := epsilon

	for c := 0; c < cols && pr < rows; c++ {
		sel := -1
		for r := pr; r < rows; r++ {
			if math.Abs(mat[r][c]) > eps {
				sel = r
				break
			}
		}
		if sel == -1 {
			continue
		}

		mat[pr], mat[sel] = mat[sel], mat[pr]

		div := mat[pr][c]
		for k := c; k <= cols; k++ {
			mat[pr][k] /= div
		}

		for r := 0; r < rows; r++ {
			if r != pr && math.Abs(mat[r][c]) > eps {
				f := mat[r][c]
				for k := c; k <= cols; k++ {
					mat[r][k] -= f * mat[pr][k]
				}
			}
		}
		pivots[c] = pr
		pr++
	}

	for r := pr; r < rows; r++ {
		if math.Abs(mat[r][cols]) > eps {
			return 0
		}
	}

	var free []int
	for c, p := range pivots {
		if p == -1 {
			free = append(free, c)
		}
	}

	minW := findMinWeight(0, make([]float64, cols), free, cols, pivots, mat, btns, target, math.MaxInt32)

	if minW == math.MaxInt32 {
		return 0
	}
	return minW
}

func part1(machines []Machine) int {
	total := 0
	for _, m := range machines {
		total += solvePart1(m.Buttons, m.Target1)
	}
	return total
}

func part2(machines []Machine) int {
	total := 0
	for _, m := range machines {
		total += solvePart2(m.Buttons, m.Target2)
	}
	return total
}

func main() {
	machines := parse(input)
	fmt.Println("Part 1:", part1(machines))
	fmt.Println("Part 2:", part2(machines))
}
