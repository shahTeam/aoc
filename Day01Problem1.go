package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/dergeberl/aoc/utils"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part1: %v\n", SolvedDay1Part1(string(input)))
	fmt.Printf("Part2: %v\n", SolvedDay1Part2(string(input)))
}

func SolvedDay1Part1(input string) int {
	calories := getCaloriesFromInput(input)
	sort.Ints(calories)
	return calories[len(calories)-1]
}

func SolvedDay1Part2(input string) int {
	calories := getCaloriesFromInput(input)
	sort.Ints(calories)
	return utils.SumSlice(calories[len(calories)-3:])
}

func getCaloriesFromInput(input string) []int {
	line, _ := utils.InputToSlice(input)

	elf := 0
	var calories []int

	for i := range line {
		if line[i] == "" {
			elf++
			continue
		}
		if len(calories)-1 < elf {
			calories = append(calories, 0)
		}
		linInt, err := strconv.Atoi(line[i])
		if err != nil {
			continue
		}
		calories[elf] += linInt
	}
	return calories
}
