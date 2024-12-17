package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

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

func operandValue(operand, A, B, C int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	default:
		panic(fmt.Sprintf("Invalid combo operand"))
	}
}

func run(a, b, c int, Program []int) []int {
	i := 0
	var out []int
	for i < len(Program) {
		op := Program[i]
		operand := Program[i+1]
		switch op {
		case 0:
			a = int(a / int(math.Pow(2, float64(operandValue(operand, a, b, c)))))
		case 1:
			b = b ^ operand
		case 2:
			b = operandValue(operand, a, b, c) % 8
		case 3:
			if a == 0 {
				break
			}
			i = operand
			continue
		case 4:
			b = b ^ c
		case 5:
			out = append(out, operandValue(operand, a, b, c)%8)
		case 6:
			b = int(a / int(math.Pow(2, float64(operandValue(operand, a, b, c)))))
		case 7:
			c = int(a / int(math.Pow(2, float64(operandValue(operand, a, b, c)))))
		}
		i += 2
	}
	return out
}

func f(a, n int) interface{} {
	if n > len(Program) {
		return a
	}
	for i := 0; i < 8; i++ {
		_a := (a << 3) | i
		out := run(_a, 0, 0, Program)
		if equal(out, Program[len(Program)-n:]) {
			result := f(_a, n+1)
			if result != nil {
				return result
			}
		}
	}
	return nil
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

var Program []int

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	parts := strings.Split(file[4], " ")
	program := strings.Split(parts[1], ",")
	for _, str := range program {
		num, _ := strconv.Atoi(str)
		Program = append(Program, num)
	}

	result := f(0, 1)
	fmt.Printf("part 2: %v\n", result)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}

