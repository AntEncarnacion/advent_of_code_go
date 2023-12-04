package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func solver(cube_quantities map[string]int, game_text string) int {
	result := 0
	for _, line := range strings.Split(game_text, "\n") {
		game_id_and_info_slice := strings.Split(line, ": ")
		game_id, _ := strconv.Atoi(strings.Split(game_id_and_info_slice[0], " ")[1])
		game_info_slice := strings.Split(game_id_and_info_slice[1], "; ")
		cube_count_map := map[string]int{"red": 0, "blue": 0, "green": 0}

		for _, round := range game_info_slice {
			for _, cube_quantity_and_color := range strings.Split(round, ", ") {
				cube_quantity_and_color_split := strings.Split(cube_quantity_and_color, " ")
				cube_color := cube_quantity_and_color_split[1]
				cube_quantity, _ := strconv.Atoi(cube_quantity_and_color_split[0])

				if cube_count_map[cube_color] < cube_quantity {
					cube_count_map[cube_color] = cube_quantity
				}
			}

		}

		if cube_count_map["red"] <= cube_quantities["red"] && cube_count_map["blue"] <= cube_quantities["blue"] && cube_count_map["green"] <= cube_quantities["green"] {
			result = result + game_id
		}
	}
	return result
}

func __FILE__() string {
	_, fn, _, _ := runtime.Caller(0)
	return filepath.Dir(fn)
}

func main() {
	cube_quantities := map[string]int{"red": 12, "green": 13, "blue": 14}
	input, _ := os.ReadFile(__FILE__() + "/../input.txt")
	fmt.Println(solver(cube_quantities, string(input)))
}
