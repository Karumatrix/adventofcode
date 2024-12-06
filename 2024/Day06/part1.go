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

func getGuardPos(file []string) Position {
	for x, line := range file {
		for y, letter := range line {
			if letter == '^' {
				return Position{x, y}
			}
		}
	}
	return Position{-1, -1}
}

func updateGuardPos(file []string, offsetX, offsetY int, guardPos Position) (Position, []string) {
	direction := file[guardPos.X][guardPos.Y]
	row := []rune(file[guardPos.X])
	row[guardPos.Y] = 'X'
	file[guardPos.X] = string(row)
	if guardPos.X < 0 || guardPos.X >= len(file) || guardPos.Y < 0 || guardPos.Y >= len(file[0]) || guardPos.X + offsetX < 0 || guardPos.X +offsetX >= len(file[0]) || guardPos.Y + offsetY < 0 || guardPos.Y + offsetY >= len(file[0]){
		return Position{-1, -1}, file
	}
	if file[guardPos.X + offsetX][guardPos.Y + offsetY] == '#' {
		switch direction {
		case '^':
			direction = '>'
			guardPos.Y++
		case '>':
			direction = 'v'
			guardPos.X++
		case '<':
			direction = '^'
			guardPos.X--
		case 'v':
			direction = '<'
			guardPos.Y--
		}
	} else {
		guardPos.X += offsetX
		guardPos.Y += offsetY
	}
	row = []rune(file[guardPos.X])
	row[guardPos.Y] = rune(direction)
	file[guardPos.X] = string(row)
	return guardPos, file
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	guardPos := getGuardPos(file)

	for {
		if guardPos.X == -1 || guardPos.Y == -1 {
			break
		}
		switch file[guardPos.X][guardPos.Y] {
		case '^':
			guardPos, file = updateGuardPos(file, -1, 0, guardPos)
		case '>':
			guardPos, file = updateGuardPos(file, 0, 1, guardPos)
		case '<':
			guardPos, file = updateGuardPos(file, 0, -1, guardPos)
		case 'v':
			guardPos, file = updateGuardPos(file, 1, 0, guardPos)
		}
	}
	result := 0
	for _, line := range file {
		for _, letter := range line {
			if letter == 'X' {
				result++
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
