# itertools

I'm using Advent of Code to learn Go.

Here is how I solved AoC 2020 Day 01:


```go
package main

import (
	"fmt"
	"strconv"

	aoc "github.com/Fiend3d/aoc_utils"
	"github.com/Fiend3d/itertools"
)

func solve(n int) int {
	lines := aoc.ReadLines("input.txt")
	numbers := make([]int, len(lines))

	for i, line := range lines {
		n, _ := strconv.Atoi(line)
		numbers[i] = n
	}

	result := 1

	for comb := range itertools.Combinations(numbers, n) {
		sum := 0
		for i := range n {
			sum += comb[i]
		}
		if sum == 2020 {
			for i := range n {
				result *= comb[i]
			}
			break
		}
	}

	return result
}

func solvePart1() {
	fmt.Printf("Part1 result: %d\n", solve(2))
}

func solvePart2() {
	fmt.Printf("Part2 result: %d\n", solve(3))
}

func main() {
	solvePart1()
	solvePart2()
}
```
