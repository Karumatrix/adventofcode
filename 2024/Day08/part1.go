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

func isAntenna(char rune) bool {
	if char >= 'a' && char <= 'z' {
		return true
	}
	if char >= 'A' && char <= 'Z' {
		return true
	}
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

type Position struct {
	X, Y int
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	antennasMap := make(map[rune][]Position)
	antinodes := make(map[Position]int)

	for indX, line := range file {
		for indY, letter := range line {
			if isAntenna(letter) {
				oldPos := antennasMap[letter]
				oldPos = append(oldPos, Position{indX, indY})
				antennasMap[letter] = oldPos
			}
		}
	}
	for _, positions := range antennasMap {
		for i := 0; i < len(positions) - 1; i++ {
			for j := i + 1; j < len(positions); j++ {
				pos1 := positions[i]
				pos2 := positions[j]
				diffPos := Position{pos2.X - pos1.X, pos2.Y - pos1.Y}
				if pos1.X - diffPos.X >= 0 && pos1.X - diffPos.X < len(file) && pos1.Y - diffPos.Y >= 0 && pos1.Y - diffPos.Y < len(file) {
					if file[pos1.X - diffPos.X][pos1.Y - diffPos.Y] == '.' {
						row := []rune(file[pos1.X - diffPos.X])
						row[pos1.Y - diffPos.Y] = '#'
						file[pos1.X - diffPos.X] = string(row)
					}
					antinodes[Position{pos1.X - diffPos.X, pos1.Y - diffPos.Y}]++
				}
				if pos2.X + diffPos.X >= 0 && pos2.X + diffPos.X < len(file) && pos2.Y + diffPos.Y >= 0 && pos2.Y + diffPos.Y < len(file) {
					if file[pos2.X + diffPos.X][pos2.Y + diffPos.Y] == '.' {
						row := []rune(file[pos2.X + diffPos.X])
						row[pos2.Y + diffPos.Y] = '#'
						file[pos2.X + diffPos.X] = string(row)
					}
					antinodes[Position{pos2.X + diffPos.X, pos2.Y + diffPos.Y}]++
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
