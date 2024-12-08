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
				for k := 0; k < len(file); k++ {
					if pos1.X - (diffPos.X * k) >= 0 && pos1.X - (diffPos.X * k) < len(file) && pos1.Y - (diffPos.Y * k) >= 0 && pos1.Y - (diffPos.Y * k) < len(file[0]) {
						if file[pos1.X - (diffPos.X * k)][pos1.Y - (diffPos.Y * k)] == '.' {
							row := []rune(file[pos1.X - (diffPos.X * k)])
							row[pos1.Y - (diffPos.Y * k)] = '#'
							file[pos1.X - (diffPos.X * k)] = string(row)
						}
						antinodes[Position{pos1.X - (diffPos.X * k), pos1.Y - (diffPos.Y * k)}]++
					}
					if pos2.X + (diffPos.X * k) >= 0 && pos2.X + (diffPos.X * k) < len(file) && pos2.Y + (diffPos.Y * k) >= 0 && pos2.Y + (diffPos.Y * k) < len(file[0]) {
						if file[pos2.X + (diffPos.X * k)][pos2.Y + (diffPos.Y * k)] == '.' {
							row := []rune(file[pos2.X + (diffPos.X * k)])
							row[pos2.Y + (diffPos.Y * k)] = '#'
							file[pos2.X + (diffPos.X * k)] = string(row)
						}
						antinodes[Position{pos2.X + (diffPos.X * k), pos2.Y + (diffPos.Y * k)}]++
					}
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
