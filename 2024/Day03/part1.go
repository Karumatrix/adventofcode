package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func readFileContent(filepath string) string {
    content, err := os.ReadFile(filepath)
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}

func part1() {
    buffer := readFileContent("./input.txt")
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	muls := re.FindAllString(buffer, -1)
	result := 0
	for _, mul := range muls {
		var int1, int2 int
		fmt.Sscanf(mul, "mul(%d,%d)", &int1, &int2)
		result += int1 * int2
	}
	fmt.Println(result)
}

func main() {
    part1()
}