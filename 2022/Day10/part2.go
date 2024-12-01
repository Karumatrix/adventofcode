package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	width  = 40
	height = 6
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

func renderPixel(screen *[height][width]rune, xRegister int, cycle int) {
	row := cycle / width
	col := cycle % width
	if row >= height {
		return
	}
	if col >= xRegister - 1 && col <= xRegister + 1 {
		screen[row][col] = '#'
	}
}

func part2() {
	instructions := stringToLines(readFileContent("./input.txt"))
	var screen [height][width]rune
	for i := range screen {
		for j := range screen[i] {
			screen[i][j] = '.'
		}
	}

	xRegister := 1
	cycle := 0

	for _, line := range instructions {
		parts := strings.Split(line, " ")
		value := 0
		if parts[0] == "addx" {
			fmt.Sscanf(parts[1], "%d", &value)
			renderPixel(&screen, xRegister, cycle)
			cycle++
		}
		renderPixel(&screen, xRegister, cycle)
		cycle++
		xRegister += value
	}

	for _, row := range screen {
		fmt.Println(string(row[:]))
	}
}

func main() {
	part2()
}

