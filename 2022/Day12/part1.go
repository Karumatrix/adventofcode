package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Position struct {
	X, Y int
}

func readFileContent(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func getGrid(filepath string) ([][]rune, Position, Position) {
	lines := readFileContent(filepath)
	rows := len(lines)
	cols := len(lines[0])

	grid := make([][]rune, rows)
	var start, end Position

	for i := 0; i < rows; i++ {
		grid[i] = []rune(lines[i])
		for j := 0; j < cols; j++ {
			switch grid[i][j] {
			case 'S':
				start = Position{i, j}
				grid[i][j] = 'a'
			case 'E':
				end = Position{i, j}
				grid[i][j] = 'z'
			}
		}
	}

	return grid, start, end
}

func bfsShortestPath(grid [][]rune, start, end Position) int {
	rows := len(grid)
	cols := len(grid[0])
	directions := []Position{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := []Position{start}
	distances := make([][]int, rows)
	for i := range distances {
		distances[i] = make([]int, cols)
	}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == end {
			return distances[curr.X][curr.Y]
		}
		for _, dir := range directions {
			newX, newY := curr.X+dir.X, curr.Y+dir.Y
			if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
				currElevation := grid[curr.X][curr.Y]
				nextElevation := grid[newX][newY]
				if nextElevation <= currElevation+1 && distances[newX][newY] == 0 && (newX != start.X || newY != start.Y) {
					distances[newX][newY] = distances[curr.X][curr.Y] + 1
					queue = append(queue, Position{newX, newY})
				}
			}
		}
	}
	return -1
}

func part1() {
	grid, start, end := getGrid("input.txt")

	result := bfsShortestPath(grid, start, end)

	if result != -1 {
		fmt.Println(result)
	} else {
		fmt.Println("No valid path found")
	}
}

func main() {
	part1()
}

