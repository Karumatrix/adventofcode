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

func part1() int {
	instructions := stringToLines(readFileContent("./input.txt"))
	cycles := 0
	X := 1
	signalStrengthSum := 0
	checkCycles := map[int]bool{20: true, 60: true, 100:true, 140: true, 180: true, 220: true}

	for _, instruction := range instructions {
		parts := strings.Fields(instruction)
		value := 0
		if parts[0] == "addx" {
			fmt.Sscanf(parts[1], "%d", &value)
			cycles++
			if checkCycles[cycles] {
				signalStrengthSum += cycles * X
			}
		}
		cycles++
		if checkCycles[cycles] {
			signalStrengthSum += cycles * X
		}
		X += value
	}
	return signalStrengthSum
}

func main() {
	fmt.Println(part1())
}