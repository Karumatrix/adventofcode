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

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	var robots []Robot
	width, height := 101, 103

	for _, line := range file {
		var px, py, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, Robot{px, py, vx, vy})
	}
	for i := 0; i < 100; i++ {
		for j := 0; j < len(robots); j++ {
			robot := &robots[j]
			robot.px = (robot.px + robot.vx + width) % width
			robot.py = (robot.py + robot.vy + height) % height
		}
	}

	quadrants := [4]int{}
	centerX, centerY := width/2, height/2

	for _, robot := range robots {
		if robot.px == centerX || robot.py == centerY {
			continue
		}
		if robot.px < centerX && robot.py < centerY {
			quadrants[0]++
		} else if robot.px >= centerX && robot.py < centerY {
			quadrants[1]++
		} else if robot.px < centerX && robot.py >= centerY {
			quadrants[2]++
		} else {
			quadrants[3]++
		}
	}
	safetyFactor := 1
	for _, count := range quadrants {
		safetyFactor *= count
	}
	fmt.Println(safetyFactor)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
