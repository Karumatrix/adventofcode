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

type MatchScore struct {
	Player, Opponent byte
	Score int
}

/*
A, X rock
B, Y paper
C, Z scissors
*/

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	result := 0
	scoreMap := []MatchScore{
		{Opponent: 'A', Player: 'X', Score: 4},
		{Opponent: 'A', Player: 'Y', Score: 8},
		{Opponent: 'A', Player: 'Z', Score: 3},
		{Opponent: 'B', Player: 'X', Score: 1},
		{Opponent: 'B', Player: 'Y', Score: 5},
		{Opponent: 'B', Player: 'Z', Score: 9},
		{Opponent: 'C', Player: 'X', Score: 7},
		{Opponent: 'C', Player: 'Y', Score: 2},
		{Opponent: 'C', Player: 'Z', Score: 6},
	}

	for _, line := range file {
		for _, score := range scoreMap {
			if score.Opponent == line[0] && score.Player == line[2] {
				result += score.Score
				break
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
