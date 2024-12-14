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

type Robot struct {
	px, py, vx, vy int
}

func printGrid(robots []Robot, width, height int) {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for _, r := range robots {
		grid[r.py][r.px] = '#'
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}


func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	var robots []Robot
	width, height := 101, 103

	for _, line := range file {
		var px, py, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, Robot{px, py, vx, vy})
	}
	i := 0
	for {
		robotMap := make(map[[2]int]int)
		for j := 0; j < len(robots); j++ {
			robot := &robots[j]
			robot.px = (robot.px + robot.vx + width) % width
			robot.py = (robot.py + robot.vy + height) % height
			robotMap[[2]int{robot.px, robot.py}]++
		}
		if len(robotMap) == len(robots) {
			printGrid(robots, width, height)
			fmt.Println("Second:", i + 1)
			os.Exit(0)
		}
		i++
	}
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
