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

func checkTrail(file []string, indX, indY int, currHeight byte) int {
	if file[indX][indY] == '9' {
		return 1
	}
	totalTrails := 0

	for _, direction := range directions {
		if indX + direction[0] >= 0 && indX + direction[0] < len(file) && indY + direction[1] >= 0 && indY + direction[1] < len(file[0]) {
			if file[indX + direction[0]][indY + direction[1]] == currHeight + 1 {
				totalTrails += checkTrail(file, indX + direction[0], indY + direction[1], currHeight + 1)
			}
		}
	}
	return totalTrails
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	result := 0

	for indX, line := range file {
		for indY, letter := range line {
			if letter == '0' {
				result += checkTrail(file, indX, indY, '0')
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
