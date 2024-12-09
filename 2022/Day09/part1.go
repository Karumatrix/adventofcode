package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFileContent(filepath string) string {
    content, err := os.ReadFile(filepath)
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}

func stringToLines(content string) []string {
    var lines []string
    scanner := bufio.NewScanner(strings.NewReader(content))

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return lines
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func moveTail(headPos [2]int, tailPos [2]int) [2]int {
	deltaX := headPos[0] - tailPos[0]
	deltaY := headPos[1] - tailPos[1]

	if abs(deltaX) > 1 || abs(deltaY) > 1 {
		if deltaX > 0 {
			tailPos[0]++
		} else if deltaX < 0 {
			tailPos[0]--
		}
		if deltaY > 0 {
			tailPos[1]++
		} else if deltaY < 0 {
			tailPos[1]--
		}
	}
	return tailPos
}

func part1() int {
	moves := stringToLines(readFileContent("./input.txt"))
	tailPos := [2]int{0,0}
	headPos := [2]int{0,0}
	seen := map[[2]int]bool{}
	seen[tailPos] = true
	for _, value := range moves {
		parts := strings.Split(value, " ")
		direction := parts[0]
		steps := 0
		fmt.Sscanf(parts[1], "%d", &steps)

		for i := 0; i < steps; i++ {
			switch direction {
			case "R":
				headPos[0]++
			case "L":
				headPos[0]--
			case "D":
				headPos[1]++
			case "U":
				headPos[1]--
			}
			tailPos = moveTail(headPos, tailPos)
			seen[tailPos] = true
		}
	}
	return len(seen)
}

func main() {
	fmt.Println(part1())
}