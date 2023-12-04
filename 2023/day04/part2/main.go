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
	split_input := strings.Split(input, "\n")
	card_count := make([]int, len(split_input))

	for card_index := range card_count {
		card_count[card_index] = 1
	}

	for card_index, line := range split_input {
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
				card_points = card_points + 1
			}
		}

		for i := 0; i < card_count[card_index]; i++ {
			for j := card_index + 1; j <= card_index+card_points; j++ {
				card_count[j] = card_count[j] + 1
			}
		}
	}

	for _, card_total := range card_count {
		result = result + card_total
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
