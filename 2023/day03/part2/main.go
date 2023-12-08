package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"unicode"
)

type coordinates struct {
	x     int
	y     int
	value string
}

type number_entity struct {
	value   int
	start_x int
	end_x   int
	y       int
}

func solver(input string) int {
	result := 0
	var symbol_coordinates []coordinates
	number_entities := make(map[int][]number_entity)
	var number_buffer []rune

	for y, line := range strings.Split(input, "\n") {
		number_flag := false
		number_buffer = nil
		start_x := -1

		for x, char := range line {
			stringified_char := string(char)
			if stringified_char == "*" {
				symbol_coordinates = append(symbol_coordinates, coordinates{x, y, string(char)})
			} else if unicode.IsDigit(char) {
				if !number_flag {
					start_x = x
					number_flag = true
				}
				number_buffer = append(number_buffer, char)
			}

			if (!unicode.IsDigit(char) && len(number_buffer) > 0) || (x == len(line)-1 && len(number_buffer) > 0) {
				value, _ := strconv.Atoi(string(number_buffer))
				number_entities[y] = append(number_entities[y], number_entity{value, start_x, x - 1, y})
				start_x = -1
				number_flag = false
				number_buffer = nil
			}
		}
	}

	for _, symbol_coordinate := range symbol_coordinates {
		number_entities_in_range := append(number_entities[symbol_coordinate.y], append(number_entities[symbol_coordinate.y-1], number_entities[symbol_coordinate.y+1]...)...)
		var gear_ratio_buffer []int
		for _, number_entity := range number_entities_in_range {
			if (symbol_coordinate.y == number_entity.y || symbol_coordinate.y == number_entity.y-1 || symbol_coordinate.y == number_entity.y+1) && (symbol_coordinate.x <= number_entity.end_x+1 && symbol_coordinate.x >= number_entity.start_x-1) {
				gear_ratio_buffer = append(gear_ratio_buffer, number_entity.value)
			}
		}
		if len(gear_ratio_buffer) == 2 {
			result = result + (gear_ratio_buffer[0] * gear_ratio_buffer[1])
		}
		gear_ratio_buffer = nil
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
