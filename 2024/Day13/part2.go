package main

import (
	"bufio"
	"fmt"
	"math"
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

type Coordinate struct {
	X, Y int
}

func readCoordinates(line string, isPrize bool, isButtA bool) (result Coordinate) {
	if isPrize {
		fmt.Sscanf(line, "Prize: X=%d, Y=%d", &result.X, &result.Y)
		result.X += 10000000000000
		result.Y += 10000000000000
	} else {
		if isButtA {
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &result.X, &result.Y)
		} else {
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &result.X, &result.Y)
		}
	}
	return result
}

func findSmallestCost(buttA, buttB, prize Coordinate) (minCost int, found bool) {
	minCost = math.MaxInt64

	detA := buttA.X * buttB.Y - buttA.Y * buttB.X

	if detA == 0 {
		fmt.Println("No solution")
		return
	}

	detA1 := prize.X * buttB.Y - prize.Y * buttB.X
	detA2 := prize.Y * buttA.X - prize.X * buttA.Y

	a := float64(detA1) / float64(detA)
	b := float64(detA2) / float64(detA)

	if int(a) * buttA.X + int(b) * buttB.X == prize.X && int(a) * buttA.Y + int(b) * buttB.Y == prize.Y {
		minCost = 3 * int(a) + int(b)
		found = true
	}
	return
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	result := 0

	for index := 0; index < len(file); index += 4 {
		buttA := readCoordinates(file[index], false, true)
		buttB := readCoordinates(file[index + 1], false, false)
		prize := readCoordinates(file[index + 2], true, false)

		cost, found := findSmallestCost(buttA, buttB, prize)
		if found {
			result += cost
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
