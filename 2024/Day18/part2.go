package main

import (
	"bufio"
	"container/heap"
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

type Node struct {
	point Position
	cost int
}

type Queue []Node

var directions = []Position {
	{0, 1},{0, -1},{1, 0},{-1, 0},
}

func (queue Queue) Len() int {
	return len(queue)
}

func (queue Queue) Less(i, j int) bool {
	return queue[i].cost < queue[j].cost
}

func (queue Queue) Swap(i, j int) {
	queue[i], queue[j] = queue[j], queue[i]
}

func (queue *Queue) Push(x interface{}) {
	*queue = append(*queue, x.(Node))
}

func (queue *Queue) Pop() interface{} {
	old := *queue
	n := len(old)
	x := old[n-1]
	*queue = old[:n-1]
	return x
}

func isValid(x, y, n int, grid[][]bool) bool {
	return x >= 0 && x < n && y >= 0 && y < n && !grid[x][y]
}

func findShortestPath(n int, grid[][]bool) int {
	start := Position{0, 0}
	end := Position{n - 1, n - 1}
	queue := &Queue{}
	heap.Init(queue)
	heap.Push(queue, Node{point: start, cost: 0})
	visited := make(map[Position]bool)
	visited[start] = true

	for queue.Len() > 0 {
		current := heap.Pop(queue).(Node)

		if current.point == end {
			return current.cost
		}

		for _, dir := range directions {
			nextX, nextY := current.point.X + dir.X, current.point.Y + dir.Y
			nextPoint := Position{nextX, nextY}

			if isValid(nextX, nextY, n, grid) && !visited[nextPoint] {
				visited[nextPoint] = true
				heap.Push(queue, Node{point: nextPoint, cost: current.cost + 1})
			}
		}
	}
	return -1
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	size := 71
	grid := make([][]bool, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]bool, size)
	}
	for _, line := range file {
		x, y := 0, 0
		fmt.Sscanf(line, "%d,%d", &x, &y)
		grid[y][x] = true
		cost := findShortestPath(size, grid)
		if cost == -1 {
			fmt.Println(line)
			break
		}
	}
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
