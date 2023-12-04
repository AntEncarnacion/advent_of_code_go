package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"unicode"
)

func reverse_string(input string) (result string) {
	var input_rune_slice []rune = []rune(input)
	var buffer []rune
	for pos := range input_rune_slice {
		buffer = append(buffer, input_rune_slice[len(input_rune_slice)-1-pos])
	}
	return string(buffer)
}

func line_calibration(line string) int {
	var first_digit, second_digit int
	var letter_buffer []rune
	spelled_numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

first_digit_loop:
	for _, char := range line {
		if unicode.IsDigit(char) && int(char-'0') != 0 {
			first_digit = int(char - '0')
			break first_digit_loop
		}

		if len(letter_buffer) > 0 {
			letter_buffer = append(letter_buffer, char)

			for number, spelled_number := range spelled_numbers {
				match, _ := regexp.MatchString("^"+spelled_number+"$", string(letter_buffer))
				if match {
					first_digit = number + 1
					break first_digit_loop
				}
			}

			for len(letter_buffer) > 0 {
				for _, spelled_number := range spelled_numbers {
					match, _ := regexp.MatchString("^"+string(letter_buffer), spelled_number)
					if match {
						continue first_digit_loop
					}
				}
				letter_buffer = letter_buffer[1:]
			}
		}
		for _, spelled_number := range spelled_numbers {
			match, _ := regexp.MatchString("^"+string(char), spelled_number)
			if match {
				letter_buffer = append(letter_buffer, char)
				continue first_digit_loop
			}
		}
	}

	var line_in_rune_slice []rune = []rune(line)
second_digit_loop:
	for pos := range line_in_rune_slice {
		char := line_in_rune_slice[len(line_in_rune_slice)-1-pos]

		if unicode.IsDigit(char) && int(char-'0') != 0 {
			second_digit = int(char - '0')
			break second_digit_loop
		}

		if len(letter_buffer) > 0 {
			letter_buffer = append(letter_buffer, char)

			for number, spelled_number := range spelled_numbers {
				reversed_spelled_number := reverse_string(spelled_number)
				match, _ := regexp.MatchString("^"+reversed_spelled_number+"$", string(letter_buffer))
				if match {
					second_digit = number + 1
					break second_digit_loop
				}
			}

			for len(letter_buffer) > 0 {
				for _, spelled_number := range spelled_numbers {
					reversed_spelled_number := reverse_string(spelled_number)
					match, _ := regexp.MatchString("^"+string(letter_buffer), reversed_spelled_number)
					if match {
						continue second_digit_loop
					}
				}
				letter_buffer = letter_buffer[1:]
			}
		}
		for _, spelled_number := range spelled_numbers {
			reversed_spelled_number := reverse_string(spelled_number)
			match, _ := regexp.MatchString("^"+string(char), reversed_spelled_number)
			if match {
				letter_buffer = append(letter_buffer, char)
				continue second_digit_loop
			}
		}
	}

	return first_digit*10 + second_digit
}

func calibration(input string) int {
	var total_sum int = 0

	for _, line := range strings.Split(input, "\n") {
		result := line_calibration(line)
		total_sum = total_sum + result
	}

	return total_sum
}

func __FILE__() string {
	_, fn, _, _ := runtime.Caller(0)
	return filepath.Dir(fn)
}

func main() {
	input, _ := os.ReadFile(__FILE__() + "/../input.txt")
	fmt.Println(calibration(string(input)))
}
