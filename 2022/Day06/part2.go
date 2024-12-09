package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func readFileContent(filepath string) string {
    content, err := ioutil.ReadFile(filepath)
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

func allUnique(s string) bool {
    seen := make(map[rune]bool)
    for _, c := range s {
        if seen[c] {
            return false
        }
        seen[c] = true
    }
    return true
}

func part2() int {
    buffer := readFileContent("input.txt")

    for i := 13; i < len(buffer); i++ {
        chars := buffer[i - 13 : i + 1]
        if allUnique(chars) {
            return i + 1
        }
    }
    return 0
}

func main() {
    fmt.Println(part2())
}
