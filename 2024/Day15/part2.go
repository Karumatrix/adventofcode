package main

import (
	"bufio"
	"fmt"
	"os"
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

type Point struct {
	r, c int
	chr rune
}

type Position struct {
	X, Y int
}

var mapGrid [50][50 * 2]rune
var toMove [50*50*2]Point
var numToMove int

func getMazeAndInputs(file []string) (maze, inputs []string) {
	i := 0
	for {
		if file[i] == "" {
			break
		}
		maze = append(maze, file[i])
		i++
	}
	i++
	for {
		if i >= len(file) {
			break
		}
		inputs = append(inputs, file[i])
		i++
	}
	return
}

func findRobotPos(maze []string) Position {
	for r, line := range maze {
		for c, char := range line {
			if char == '@' {
				return Position{r, c}
			}
		}
	}
	return Position{-1, -1}
}

func resizeMaze(maze []string) []string {
	resized := []string{}
	for _, line := range maze {
		newLine := strings.Builder{}
		for _, char := range line {
			switch char {
			case '#':
				newLine.WriteString("##")
			case 'O':
				newLine.WriteString("[]")
			case '.':
				newLine.WriteString("..")
			case '@':
				newLine.WriteString("@.")
			}
		}
		resized = append(resized, newLine.String())
	}
	return resized
}

func getToMove(r, c int, direction rune) bool {
	if mapGrid[r][c] == '.' {
		return false
	}
	if mapGrid[r][c] == '#' {
		return true
	}
	if mapGrid[r][c] == ']' && (direction == '^' || direction == 'v') {
		c--
	}

	result := true
	switch direction {
	case 'v':
		if r+1 < 50-1 {
			if mapGrid[r][c] == '[' {
				result = getToMove(r+1, c, direction) || getToMove(r+1, c+1, direction)
			} else {
				result = getToMove(r+1, c, direction)
			}
		}
	case '^':
		if r-1 > 0 {
			if mapGrid[r][c] == '[' {
				result = getToMove(r-1, c, direction) || getToMove(r-1, c+1, direction)
			} else {
				result = getToMove(r-1, c, direction)
			}
		}
	case '>':
		if c+1 < 50*2-1 {
			result = getToMove(r, c+1, direction)
		}
	case '<':
		if c-1 > 0 {
			result = getToMove(r, c-1, direction)
		}
	}

	if !includesPosition(r, c) {
		toMove[numToMove] = Point{r, c, mapGrid[r][c]}
		numToMove++
	}

	if mapGrid[r][c] == '[' && (direction == '^' || direction == 'v') {
		if !includesPosition(r, c+1) {
			toMove[numToMove] = Point{r, c + 1, ']'}
			numToMove++
		}
	}

	return result
}

func includesPosition(r, c int) bool {
	for i := 0; i < numToMove; i++ {
		if r == toMove[i].r && c == toMove[i].c {
			return true
		}
	}
	return false
}

func printMap() {
	for r := 0; r < 50; r++ {
		for c := 0; c < 50*2; c++ {
			fmt.Print(string(mapGrid[r][c]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	maze, inputs := getMazeAndInputs(file)
	rr, rc := 0, 0

	for r := 0; r < len(maze); r++ {
		for c := 0; c < len(maze[r]); c++ {
			if maze[r][c] == '@' {
				rr = r;
				rc = c*2;
				mapGrid[r][c*2] = '@'
				mapGrid[r][c*2+1] = '.'
			} else if maze[r][c] == 'O' {
				mapGrid[r][c*2] = '['
				mapGrid[r][c*2+1] = ']'
			} else {
				mapGrid[r][c*2] = rune(maze[r][c])
				mapGrid[r][c*2+1] = rune(maze[r][c])
			}
		}
	}
	for _, line := range inputs {
		for _, dir := range line {
			numToMove = 0
			blocked := getToMove(rr, rc, dir)
			if !blocked {
				switch dir {
				case 'v':
					rr++
				case '^':
					rr--
				case '>':
					rc++
				case '<':
					rc--
				}

				for i := 0; i < numToMove; i++ {
					mapGrid[toMove[i].r][toMove[i].c] = '.'
					switch dir {
					case 'v':
						mapGrid[toMove[i].r+1][toMove[i].c] = toMove[i].chr
					case '^':
						mapGrid[toMove[i].r-1][toMove[i].c] = toMove[i].chr
					case '>':
						mapGrid[toMove[i].r][toMove[i].c+1] = toMove[i].chr
					case '<':
						mapGrid[toMove[i].r][toMove[i].c-1] = toMove[i].chr
					}
				}
			}
		}
	}

	result := 0
	for r := 0; r < 50; r++ {
		for c := 0; c < 100; c++ {
			if mapGrid[r][c] == '[' {
				result += 100 * r + c
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

