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
	santaX, santaY, robX, robY := 0, 0, 0, 0
	isSanta := true

	housesMap[Position{santaX, santaY}]++
	housesMap[Position{robX, robY}]++
	for _, letter := range file {
		if isSanta {
			switch letter {
			case '<':
				santaY--
			case '>':
				santaY++
			case 'v':
				santaX++
			case '^':
				santaX--
			}
			housesMap[Position{santaX, santaY}]++
			isSanta = false
		} else {
			switch letter {
			case '<':
				robY--
			case '>':
				robY++
			case 'v':
				robX++
			case '^':
				robX--
			}
			housesMap[Position{robX, robY}]++
			isSanta = true
		}
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
