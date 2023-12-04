package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func linear_search(haystack []string, needle string) bool {
	for _, char := range haystack {
		if char == needle {
			return true
		}
	}
	return false
}

func solver(input string) int {
	result := 0

	for _, line := range strings.Split(input, "\n") {
		line = strings.Split(line, ": ")[1]
		line := strings.Split(line, " | ")
		winning_numbers := strings.Split(line[0], " ")
		card_numbers := strings.Split(line[1], " ")
		card_points := 0

		for _, card_number := range card_numbers {
			if card_number == "" {
				continue
			}
			if linear_search(winning_numbers, card_number) {
				if card_points == 0 {
					card_points = 1
					continue
				}

				card_points = card_points * 2
			}
		}
		result = result + card_points
	}

	return result
}

func __FILE__() string {
	_, fn, _, _ := runtime.Caller(0)
	return filepath.Dir(fn)
}

func main() {
	input, _ := os.ReadFile(__FILE__() + "/../input.txt")
	fmt.Println(solver(string(input)))
}
