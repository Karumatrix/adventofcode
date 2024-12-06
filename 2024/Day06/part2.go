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
	if file[guardPos.X + offsetX][guardPos.Y + offsetY] == '#' || file[guardPos.X + offsetX][guardPos.Y + offsetY] == 'O' {
		switch direction {
		case '^':
			direction = '>'
		case '>':
			direction = 'v'
		case '<':
			direction = '^'
		case 'v':
			direction = '<'
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

func copyFile(file []string) []string {
    copy := make([]string, len(file))
    for i := range file {
        copy[i] = file[i]
    }
    return copy
}

func testMap(guardPos Position, file []string) bool {
	copyFile := copyFile(file)
	directionsCheck := make(map[Position]string)
	directionsCheck[guardPos] = string(copyFile[guardPos.X][guardPos.Y])
	for {
		switch copyFile[guardPos.X][guardPos.Y] {
		case '^':
			guardPos, copyFile = updateGuardPos(copyFile, -1, 0, guardPos)
		case '>':
			guardPos, copyFile = updateGuardPos(copyFile, 0, 1, guardPos)
		case '<':
			guardPos, copyFile = updateGuardPos(copyFile, 0, -1, guardPos)
		case 'v':
			guardPos, copyFile = updateGuardPos(copyFile, 1, 0, guardPos)
		}
		if guardPos.X == -1 || guardPos.Y == -1 {
			break
		}
		for _, dir := range directionsCheck[guardPos] {
			if dir == rune(copyFile[guardPos.X][guardPos.Y]) {
				return true
			}
		}
		directionsCheck[guardPos] += string(copyFile[guardPos.X][guardPos.Y])
	}
	return false
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	originalGuardPos := getGuardPos(file)
	result := 0

	for indexX, line := range file {
		for indexY, _ := range line {
			if file[indexX][indexY] == '.' {
				fmt.Println("Testing:", indexX, indexY)
				row := []rune(file[indexX])
				row[indexY] = 'O'
				file[indexX] = string(row)
				if testMap(originalGuardPos, file) {
					result++
				}
				row = []rune(file[indexX])
				row[indexY] = '.'
				file[indexX] = string(row)
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
