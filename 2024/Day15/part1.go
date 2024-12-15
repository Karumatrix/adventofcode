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

func findRobotPos(maze []string) (result Position) {
	for x, line := range maze {
		for y, char := range line {
			if char == '@' {
				return Position{x, y}
			}
		}
	}
	return Position{-1, -1}
}

func changePosition(maze []string, dir rune, robot *Position) []string {
    directions := map[rune]Position{
        '>': {0, 1}, '<': {0, -1}, '^': {-1, 0}, 'v': {1, 0},
    }

    move := directions[dir]
    next := Position{robot.X + move.X, robot.Y + move.Y}

    if maze[next.X][next.Y] == '#' {
        return maze
    }

    boxes := []Position{}
    for cursor := next; maze[cursor.X][cursor.Y] == 'O'; cursor.X, cursor.Y = cursor.X+move.X, cursor.Y+move.Y {
        boxes = append(boxes, cursor)
    }

    lastBox := next
    if len(boxes) > 0 {
        lastBox = boxes[len(boxes)-1]
        afterLastBox := Position{lastBox.X + move.X, lastBox.Y + move.Y}
        if maze[afterLastBox.X][afterLastBox.Y] != '.' {
            return maze
        }
    }

    for i := len(boxes) - 1; i >= 0; i-- {
        from := boxes[i]
        to := Position{from.X + move.X, from.Y + move.Y}
        maze = updateMaze(maze, from, to)
    }
    maze = updateMaze(maze, *robot, next)
    robot.X, robot.Y = next.X, next.Y
    return maze
}

func updateMaze(maze []string, from, to Position) []string {
    fromLine := []rune(maze[from.X])
    toLine := []rune(maze[to.X])

	if from.X == to.X {
		toLine[to.Y] = fromLine[from.Y]
		toLine[from.Y] = '.'
		maze[to.X] = string(toLine)
		} else {
		toLine[to.Y] = fromLine[from.Y]
		fromLine[from.Y] = '.'
		maze[from.X] = string(fromLine)
		maze[to.X] = string(toLine)
	}
    return maze
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	maze, inputs := getMazeAndInputs(file)
	robot := findRobotPos(maze)

	for _, line := range inputs {
		for _, dir := range line {
			maze = changePosition(maze, dir, &robot)
		}
	}
	result := 0
	for x, line := range maze {
		for y, char := range line {
			if char == 'O' {
				result += 100 * x + y
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
