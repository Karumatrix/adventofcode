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

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	buttons := []string{
		"..1..",
		".234.",
		"56789",
		".ABC.",
		"..D..",
	}
	x, y := 3, 0

	for _, line := range file {
		for _, dir := range line {
			i, j := x, y
			switch dir {
			case 'R':
				j++
			case 'L':
				j--
			case 'U':
				i--
			case 'D':
				i++
			}
			if i >= 0 && i < 5 && j >= 0 && j < 5 && buttons[i][j] != '.' {
				x, y = i, j
			}
		}
		fmt.Print(string(buttons[x][y]))
	}
	fmt.Println()
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
