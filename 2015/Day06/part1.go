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

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	width, length := 1000, 1000
	lights := make([][]bool, width)
	for i := range lights {
		lights[i] = make([]bool, length)
	}

	for _, line := range file {
		parts := strings.Split(line, " ")
		if parts[0] == "toggle" {
			coorX1, coorY1, coorX2, coorY2 := 0, 0, 0, 0
			fmt.Sscanf(parts[1], "%d,%d", &coorX1, &coorY1)
			fmt.Sscanf(parts[3], "%d,%d", &coorX2, &coorY2)
			for i := coorX1; i <= coorX2; i++ {
				for j := coorY1; j <= coorY2; j++ {
					lights[i][j] = !lights[i][j]
				}
			}
		}
		if parts[0] == "turn" {
			onOff := false
			coorX1, coorY1, coorX2, coorY2 := 0, 0, 0, 0
			fmt.Sscanf(parts[2], "%d,%d", &coorX1, &coorY1)
			fmt.Sscanf(parts[4], "%d,%d", &coorX2, &coorY2)
			if parts[1] == "on" {
				onOff = true
			}
			for i := coorX1; i <= coorX2; i++ {
				for j := coorY1; j <= coorY2; j++ {
					lights[i][j] = onOff
				}
			}
		}
	}
	result := 0
	for _, line := range lights {
		for _, light := range line {
			if light {
				result++
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
