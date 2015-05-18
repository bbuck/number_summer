package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/seer-server/script-engine"
)

var (
	sumNumbers = `
		function process(nums)
			sum = 0
			for i = 1, #nums do
				sum = sum + nums[i]
			end

			print("The sum of the given values was " .. sum)
		end
	`

	multNumbers = `
		function process(nums)
			tot = 1
			for i = 1, #nums do
				tot = tot * nums[i]
			end

			print("The product of the given values was " .. tot)
		end
	`
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("USAGE: %s --process PROCESS num[, num[, num[, ...]]]\n", os.Args[0])

		return
	}

	process := flag.String("process", "sum", "Which process to peform on the numbers")
	flag.Parse()

	e := lua.NewEngine()
	nums := e.NewTable()

	numStrings := flag.Args()
	for _, ns := range numStrings {
		intVal, err := strconv.Atoi(ns)
		if err != nil {
			fmt.Printf("Invalid number %q was given!\n", ns)

			return
		}
		nums.Append(intVal)
	}

	switch *process {
	default:
		if err := e.LoadString(sumNumbers); err != nil {
			fmt.Printf("Script Error: %s\n", err)

			return
		}
	case "mult":
		if err := e.LoadString(multNumbers); err != nil {
			fmt.Printf("Script Error: %s\n", err)

			return
		}
	}

	_, err := e.Call("process", 0, nums)
	if err != nil {
		fmt.Printf("Script Error: %s\n", err)

		return
	}
}
