package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func readFileContent(filepath string) string {
    content, err := os.ReadFile(filepath)
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}

func part2() {
    buffer := readFileContent("./input.txt")
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	instructions := re.FindAllString(buffer, -1)
	result := 0
	enable := true
	for _, instruction := range instructions {
		if instruction == "do()" {
			enable = true
		} else if instruction == "don't()" {
			enable = false
		} else if enable == true && strings.HasPrefix(instruction, "mul") {
			var int1, int2 int
			fmt.Sscanf(instruction, "mul(%d,%d)", &int1, &int2)
			result += int1 * int2
		}
	}
	fmt.Println(result)
}

func main() {
    part2()
}