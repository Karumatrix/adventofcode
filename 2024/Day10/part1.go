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

type Case struct {
	X, Y int
	Height byte
}

var directions = [][2]int {
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func checkTrail(file []string, indX, indY int) int {
	toCheck := []Case{{indX, indY, '0'}}
	reachable := make(map[[2]int]bool)

	for len(toCheck) > 0 {
		x, y, height := toCheck[0].X, toCheck[0].Y, toCheck[0].Height
		toCheck = toCheck[1:]

		if file[x][y] == '9' {
			reachable[[2]int{x, y}] = true
			continue
		}
		for _, direction := range directions {
			if x + direction[0] >= 0 && x + direction[0] < len(file) && y + direction[1] >= 0 && y + direction[1] < len(file[0]) {
				if file[x + direction[0]][y + direction[1]] == height + 1 {
					toCheck = append(toCheck, Case{x + direction[0], y + direction[1], height + 1})
				}
			}
		}
	}
	return len(reachable)
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	result := 0

	for indX, line := range file {
		for indY, letter := range line {
			if letter == '0' {
				result += checkTrail(file, indX, indY)
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
