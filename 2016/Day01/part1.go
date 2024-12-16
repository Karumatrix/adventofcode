package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) Turn(turn rune) Direction {
	if turn == 'L' {
		return (d + 3) % 4
	}
	return (d + 1) % 4
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func aoc(filepath string) {
	file := filepathToString(filepath)
	if file[len(file) - 1] == '\n' {
		file = file[:len(file) - 1]
	}
	directions := strings.Split(file, ", ")
	direction := North
	x, y := 0, 0

	for _, dir := range directions {
		turn := rune(dir[0])
		steps, _ := strconv.Atoi(dir[1:])
		direction = direction.Turn(turn)
		switch direction {
		case North:
			x -= steps
		case South:
			x += steps
		case East:
			y += steps
		case West:
			y -= steps
		}
	}
	distance := abs(x) + abs(y)
	fmt.Println(distance)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
