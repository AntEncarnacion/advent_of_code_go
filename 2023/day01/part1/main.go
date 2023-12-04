package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
)

func line_calibration(line string) int {
	var first_digit, second_digit int
	var reversed_runes []rune = []rune(line)

	for _, char := range line {
		if unicode.IsDigit(char) {
			first_digit = int(char - '0')
			break
		}
	}

	for pos := range reversed_runes {
		char := reversed_runes[len(reversed_runes)-1-pos]
		if unicode.IsDigit(char) {
			second_digit = int(char - '0')
			break
		}
	}

	return first_digit*10 + second_digit
}

func calibration(input string) int {
	var total_sum int = 0

	for _, line := range strings.Split(input, "\n") {
		total_sum = total_sum + line_calibration(line)
	}

	return total_sum
}

// TODO improvement: why reverse the string? just get both digits in one go

func __FILE__() string {
	_, fn, _, _ := runtime.Caller(0)
	return filepath.Dir(fn)
}

func main() {
	input, _ := os.ReadFile(__FILE__() + "/../input.txt")
	fmt.Println(calibration(string(input)))
}
