package main

import (
	"bufio"
	"fmt"
	"os"
)

type Limits struct {
	Top, Bottom, Left, Right int
}

type Position struct {
	X, Y int
}

func readInput(fileName string) [][]byte {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var f [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f = append(f, []byte(line))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return f
}

func replace1WithDot(f [][]byte, b Limits) {
	for z := b.Top; z <= b.Bottom; z++ {
		for w := b.Left; w <= b.Right; w++ {
			if f[z][w] == '1' {
				f[z][w] = '.'
			}
		}
	}
}

func getSides(f [][]byte, limits Limits) [][][]int {
	rows := limits.Bottom - limits.Top + 1
	cols := limits.Right - limits.Left + 1

	res := make([][][]int, 4)
	for i := 0; i < 4; i++ {
		if i < 2 {
			res[i] = make([][]int, rows)
		} else {
			res[i] = make([][]int, cols)
		}
	}

	for i := limits.Top; i <= limits.Bottom; i++ {
		for j := limits.Left; j <= limits.Right; j++ {
			if f[i][j] == '1' {
				if j == 0 || f[i][j-1] != '1' {
					res[0][i-limits.Top] = append(res[0][i-limits.Top], j)
				}
				if j+1 >= len(f[0]) || f[i][j+1] != '1' {
					res[1][i-limits.Top] = append(res[1][i-limits.Top], j)
				}
				if i == 0 || f[i-1][j] != '1' {
					res[2][j-limits.Left] = append(res[2][j-limits.Left], i)
				}
				if i+1 >= len(f) || f[i+1][j] != '1' {
					res[3][j-limits.Left] = append(res[3][j-limits.Left], i)
				}
			}
		}
	}

	return res
}

func getShape(f [][]byte, i, j int) (int, [][][]int) {
	dirs := []Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	searched := f[i][j]
	limits := Limits{i, i, j, j}
	area := 0

	var flood func(x, y int)
	flood = func(x, y int) {
		if x < 0 || x >= len(f) || y < 0 || y >= len(f[0]) || f[x][y] != searched {
			return
		}

		if x < limits.Top {
			limits.Top = x
		}
		if x > limits.Bottom {
			limits.Bottom = x
		}
		if y < limits.Left {
			limits.Left = y
		}
		if y > limits.Right {
			limits.Right = y
		}

		f[x][y] = '1'
		area++

		for _, d := range dirs {
			flood(x+d.X, y+d.Y)
		}
	}

	flood(i, j)
	sides := getSides(f, limits)
	replace1WithDot(f, limits)

	return area, sides
}

func getTotalSides(sides [][][]int) int {
	tot := 0
	for _, side := range sides {
		for i := 0; i < len(side); i++ {
			for len(side[i]) > 0 {
				tot++
				elt := side[i][0]
				side[i] = side[i][1:]
				j := i + 1
				for j < len(side) {
					found := false
					for k, v := range side[j] {
						if v == elt {
							side[j] = append(side[j][:k], side[j][k+1:]...)
							found = true
							break
						}
					}
					if !found {
						break
					}
					j++
				}
			}
		}
	}
	return tot
}

func main() {
	f := readInput("input.txt")
	score2 := 0

	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[0]); j++ {
			if f[i][j] == '.' {
				continue
			}

			area, sides := getShape(f, i, j)
			circumference := 0
			for _, side := range sides {
				for _, val := range side {
					circumference += len(val)
				}
			}

			score2 += getTotalSides(sides) * area
		}
	}

	fmt.Println(score2)
}
