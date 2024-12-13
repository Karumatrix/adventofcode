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

func getSides(file []string, rows, cols int, cases map[Position]bool, target byte) (sides int) {
	interiorAngles := 0
	directions := []Position{
		{-1, 0},{1, 0},{0, -1},{0, 1},
	}
	checked := make(map[Position]int)
	for pos := range cases {
		nbNeighbour := 0
		for _, dir := range directions {
			if pos.X + dir.X < 0 || pos.X + dir.X >= len(file) {
				nbNeighbour++
			}
			if pos.Y + dir.Y < 0 || pos.Y + dir.Y >= len(file[0]) {
				nbNeighbour++
			}
			if pos.X + dir.X >= 0 && pos.X + dir.X < len(file) && pos.Y + dir.Y >= 0 && pos.Y + dir.Y < len(file[0]) && file[pos.X + dir.X][pos.Y + dir.Y] != target {
				nbNeighbour++
				checked[Position{pos.X + dir.X, pos.Y + dir.Y}]++
			}
		}

		switch nbNeighbour {
		case 2:
			fmt.Print("+90")
			interiorAngles += 90
			break
		case 3:
			fmt.Print("+180")
			interiorAngles += 180
			break
		}
	}
	for _, nb := range checked {
		if nb == 2 {
			fmt.Print("+270")
			interiorAngles += 270
		}
	}
	fmt.Println()
	sides = (interiorAngles / 180) + 2
	return
}

func getArea(file []string, x, y, rows, cols int, region byte, checked [][]bool, cases map[Position]bool) (area int) {
	if x < 0 || y < 0 || x >= rows || y >= cols || file[x][y] != region {
		return 0
	}
	if checked[x][y] {
		return 0
	}
	checked[x][y] = true
	cases[Position{x, y}] = true
	area = 1
	directions := []Position{
		{-1, 0},{1, 0},{0, -1},{0, 1},
	}
	for _, dir := range directions {
		area += getArea(file, x + dir.X, y + dir.Y, rows, cols, region, checked, cases)
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
				cases := make(map[Position]bool)
				area := getArea(file, i, j, rows, cols, file[i][j], checked, cases)
				sides := getSides(file, rows, cols, cases, file[i][j])
				fmt.Println(area, sides)
				fmt.Println()
				result += area * sides
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
