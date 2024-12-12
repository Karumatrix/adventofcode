package main

import (
	"bufio"
	"fmt"
	"os"
)

func filepathToString(filepath string) string {
    content, err := os.ReadFile(filepath)
    if err != nil {
		panic(err)
    }
    return string(content)
}

func filepathToStringArray(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

type Position struct {
	X, Y int
}

func getAreaAndPerimeter(file []string, x, y, rows, cols int, region byte, checked [][]bool) (area, perimeter int) {
	if x < 0 || y < 0 || x >= rows || y >= cols || file[x][y] != region {
		return 0, 1
	}
	if checked[x][y] {
		return 0, 0
	}
	checked[x][y] = true
	area = 1
	perimeter = 0
	directions := []Position{
		{-1, 0},{1, 0},{0, -1},{0, 1},
	}
	for _, dir := range directions {
		da, dp := getAreaAndPerimeter(file, x + dir.X, y + dir.Y, rows, cols, region, checked)
		area += da
		perimeter += dp
	}
	return
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)

	rows := len(file)
	cols := len(file[0])
	checked := make([][]bool, rows)
	for i := range checked {
		checked[i] = make([]bool, cols)
	}
	result := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !checked[i][j] {
				area, perimeter := getAreaAndPerimeter(file, i, j, rows, cols, file[i][j], checked)
				result += area * perimeter
			}
		}
	}
	fmt.Println(result)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
