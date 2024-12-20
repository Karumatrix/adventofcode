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

func getMaze(file []string) (maze [][]int, start, end Position) {
	rows := len(file)
	cols := len(file[0])
	maze = make([][]int, rows)
	for indX, line := range file {
		maze[indX] = make([]int, cols)
		for indY, letter := range line {
			switch letter {
			case 'S':
				start = Position{indX, indY}
				maze[indX][indY] = -1
			case 'E':
				end = Position{indX, indY}
				maze[indX][indY] = -2
			case '.':
				maze[indX][indY] = -3
			case '#':
				maze[indX][indY] = -4
			}
		}
	}
	return
}

func findPath(maze [][]int, start, end Position) int {
	directions := []Position{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}
	score := 0

	for {
		if start.X == end.X && start.Y == end.Y {
			break
		}
		maze[start.X][start.Y] = score
		score++

		for _, dir := range directions {
			if maze[start.X + dir.X][start.Y + dir.Y] == -3 || maze[start.X + dir.X][start.Y + dir.Y] == -2 {
				maze[start.X + dir.X][start.Y + dir.Y] = score
				start.X += dir.X
				start.Y += dir.Y
				break
			}
		}
	}
	return score
}

func abs(nb int) int {
	if nb < 0 {
		return -nb
	}
	return nb
}

func getScore(maze [][]int, size int) int {
	directions := make(map[Position]bool)
	for i := 0; i < size; i++ {
		for j := 0; j < size - i; j++ {
			directions[Position{i, j}] = true
			directions[Position{i, -j}] = true
			directions[Position{-i, j}] = true
			directions[Position{-i, -j}] = true
		}
	}
	score := 0
	for x, line := range maze {
		for y, tmp := range line {
			if tmp == -4 {
				continue
			}
			for dir := range directions {
				tmpX, tmpY := x + dir.X, y + dir.Y
				if tmpX < 0 || tmpX >= len(maze) || tmpY < 0 || tmpY >= len(maze[0]) {
					continue
				}
				if maze[tmpX][tmpY] == -4 {
					continue
				}
				if tmp > maze[tmpX][tmpY] {
					continue
				}
				dist := maze[tmpX][tmpY] - tmp - abs(dir.X) - abs(dir.Y)
				if dist >= 100 {
					score++
				}
			}
		}
	}
	return score
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	maze, start, end := getMaze(file)
	findPath(maze, start, end)
	fmt.Println(getScore(maze, 3))
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
