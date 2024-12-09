package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	Items      []int
	Operation  func(int) int
	TestDiv    int
	IfTrue     int
	IfFalse    int
	Inspections int
}

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

func parseOperation(opStr string) func(int) int {
	parts := strings.Fields(opStr)
	if len(parts) != 3 {
		fmt.Println(opStr)
		panic("Invalid operation format")
	}

	operand := parts[2]
	operator := parts[1]

	return func(old int) int {
		operandVal := 0
		if operand == "old" {
			operandVal = old
		} else {
			fmt.Sscanf(operand, "%d", &operandVal)
		}

		switch operator {
		case "*":
			return old * operandVal
		case "+":
			return old + operandVal
		default:
			panic("Unknown operator")
		}
	}
}

func createMonkeys(buffer []string) []Monkey {
	monkeys := []Monkey{}

	for i := 0; i < len(buffer); i += 7 {
		monkey := Monkey{}
		itemsStr := strings.Split(strings.TrimPrefix(buffer[i + 1], "  Starting items: "), ", ")
		for _, itemStr := range itemsStr {
			item := 0
			fmt.Sscanf(itemStr, "%d", &item)
			monkey.Items = append(monkey.Items, item)
		}
		opStr := strings.TrimPrefix(buffer[i + 2], "  Operation: new = ")
		monkey.Operation = parseOperation(opStr)
		monkey.TestDiv, _ = strconv.Atoi(strings.TrimPrefix(buffer[i + 3], "  Test: divisible by "))
		monkey.IfTrue, _ = strconv.Atoi(strings.TrimPrefix(buffer[i + 4], "    If true: throw to monkey "))
		monkey.IfFalse, _ = strconv.Atoi(strings.TrimPrefix(buffer[i + 5], "    If false: throw to monkey "))
		monkeys = append(monkeys, monkey)
	}

	return monkeys
}

func part1() int {
	buffer := stringToLines(readFileContent("./input.txt"))
	monkeys := createMonkeys(buffer)
	for round := 0; round < 20; round++ {
		for i := 0; i < len(monkeys); i++ {
			currentMonkey := &monkeys[i]
			newItems := make([]int, 0)
			for _, item := range currentMonkey.Items {
				currentMonkey.Inspections++
				worryLevel := currentMonkey.Operation(item)
				worryLevel = int(math.Floor(float64(worryLevel) / 3))
				if worryLevel % currentMonkey.TestDiv == 0 {
					monkeys[currentMonkey.IfTrue].Items = append(monkeys[currentMonkey.IfTrue].Items, worryLevel)
				} else {
					monkeys[currentMonkey.IfFalse].Items = append(monkeys[currentMonkey.IfFalse].Items, worryLevel)
				}
			}
			currentMonkey.Items = newItems
		}
	}
	highest := 0
	secondHighest := 0
	for _, monkey := range monkeys {
		if monkey.Inspections > highest {
			secondHighest = highest
			highest = monkey.Inspections
		} else if monkey.Inspections > secondHighest {
			secondHighest = monkey.Inspections
		}
	}
	levelOfMonkeyBusiness := highest * secondHighest
	return levelOfMonkeyBusiness
}

func main() {
	fmt.Println(part1())
}
