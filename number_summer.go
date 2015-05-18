package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("USAGE: %s num[, num[, num[, ...]]]\n", os.Args[0])

		return
	}

	numStrings := os.Args[1:]
	nums := make([]int, len(numStrings))
	for i, ns := range numStrings {
		intVal, err := strconv.Atoi(ns)
		if err != nil {
			fmt.Printf("Invalid number %q was given!\n", ns)
		}
		nums[i] = intVal
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}

	fmt.Printf("The sum of the given values was %d\n", sum)
}
