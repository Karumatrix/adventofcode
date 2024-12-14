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
	x, y int
}

func aoc(filepath string) {
	file := filepathToString(filepath)
	housesMap := make(map[Position]int)
	x, y := 0, 0

	housesMap[Position{x, y}]++
	for _, letter := range file {
		switch letter {
		case '<':
			y--
		case '>':
			y++
		case 'v':
			x++
		case '^':
			x--
		}
		housesMap[Position{x, y}]++
	}
	fmt.Println(len(housesMap))
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
