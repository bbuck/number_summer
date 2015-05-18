package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func sumNumbers(nums []int) {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	fmt.Printf("The sum of the given values was %d\n", sum)
}

func multNumbers(nums []int) {
	tot := 1
	for _, n := range nums {
		tot *= n
	}

	fmt.Printf("The product of the given values was %d\n", tot)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("USAGE: %s --process PROCESS num[, num[, num[, ...]]]\n", os.Args[0])

		return
	}

	process := flag.String("process", "sum", "Which process to peform on the numbers")
	flag.Parse()

	numStrings := flag.Args()
	nums := make([]int, len(numStrings))
	for i, ns := range numStrings {
		intVal, err := strconv.Atoi(ns)
		if err != nil {
			fmt.Printf("Invalid number %q was given!\n", ns)

			return
		}
		nums[i] = intVal
	}

	switch *process {
	default:
		sumNumbers(nums)
	case "mult":
		multNumbers(nums)
	}
}
